package ssh_config

import (
	"errors"
	"fmt"
	"strings"
)

// func Decode(data io.Reader) error {
// type scannerstate struct {
// 	level string // enum: "", "host", "match"
// }
// 	var ss scannerstate
// 	scanner := bufio.NewScanner(data)
// 	for scanner.Scan() {
// 		parseLine(strings.ToValidUTF8(scanner.Text(), ""))
// 	}
// 	return scanner.Err() // no need to wrap
// }

var (
	errInvalidKeyword = errors.New("invalid keyword")
	errInvalidQuoting = errors.New("bad quoting") // TODO: add more info?
)

type rawObject Keyword

// returned Keyword might be a raw TLD object
func parseLine(data string) (rawObject, error) {
	trimmedLine := strings.TrimSpace(data)
	if strings.HasPrefix(trimmedLine, "#") {
		// dedicated comment line
		return rawObject{Comment: strings.TrimPrefix(trimmedLine, "#")}, nil
	}
	kvSeperatorPos := strings.IndexAny(trimmedLine, " =")
	key := trimmedLine[:kvSeperatorPos]
	valuesblob := trimmedLine[kvSeperatorPos+1:]
	_, _, err := unescapeValue(valuesblob) //TODO
	if err != nil {
		return rawObject{}, fmt.Errorf("%q: %w", data, err)
	}

	return rawObject{Key: key}, nil //TODO::::::

	// valuetype, keyInIndex := keywordIndex[key]
	// if !keyInIndex {
	// 	return rawObject{Key: key, Values: }, errInvalidKeyword
	// }

	// strings.FieldsFunc(data)
	// strings.SplitN(data, " ")
}

// unescapes an ssh_config valuepart
//
// inspired by https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz misc.c#1889 and strings.FieldsFunc()
func unescapeValue(s string) (values []string, comment string, err error) {
	strings, currentString, skipnext, quoted := []string{}, "", 0, 0
	for pos, rune := range s {
		if skipnext > 0 {
			skipnext -= 1
			continue
		}
		if rune == '\\' { // single backslash
			switch string(s[pos+1]) {
			case "'":
				currentString += "'"
			case "\"":
				currentString += "\""
			case "\\":
				currentString += "\\\\" // 2 backslashes
			default: // turn single backslash without a pairing backslash, or accepted quotes to double, as OpenSSH does
				currentString += "\\\\" // 2 backslashes
				continue
			}
			skipnext += 1 // escapes are parsed in sets, next char already got parsed and added to currentString
			continue
		}

		switch quoted { // while quoted, ignore the other quoting style
		default:
			panic(fmt.Sprintf("unescapeValue runtime enum: quoted should always be 0..2, is %d", quoted))
		case 1: // single quoted
			if rune == '\'' {
				quoted = 0 // macro a
				strings = append(strings, currentString)
				currentString = ""
				// spaces between are handled below
				continue
			}
		case 2: // double quoted
			if rune == '"' {
				quoted = 0 // duplicate codeblock a
				strings = append(strings, currentString)
				currentString = ""
				// spaces between are handled below
				continue
			}
		case 0: // not quoted
			if rune == '\'' {
				quoted = 1
				continue
			}
			if rune == '"' {
				quoted = 2
				continue
			}
			if rune == ' ' {
				if len(currentString) == 0 {
					continue
				}
				strings = append(strings, currentString)
				currentString = ""
				continue
			}
			if rune == '#' {
				comment = s[pos+1:]
				break
			}
		}
		currentString += string(rune)
	}
	if quoted != 0 {
		return nil, "", fmt.Errorf("unescapeValues: %q: %w", currentString, errInvalidQuoting)
	}
	return append(strings, currentString), comment, nil
}

// encoding: indentchar
