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
	errInvalidKeyword                 = errors.New("invalid keyword")
	errInvalidQuoting                 = errors.New("bad quoting") // TODO: add more info?
	errWarnSingleBackslashTransformed = errors.New("1 or more single backslashes changed to 2 backslashes since OpenSSH ssh_config does this (this always happens: st\\ring → st\\\\ring)")
)

type rawObject Keyword

// returned Keyword might be a raw TLD object
//
// possible errors: nil, errInvalidQuoting
func parseLine(data string) (rawObject, error) {
	trimmedLine := strings.TrimSpace(data)
	if strings.HasPrefix(trimmedLine, "#") {
		// dedicated comment line
		return rawObject{Comment: strings.TrimPrefix(trimmedLine, "#")}, nil
	}
	kvSeperatorPos := strings.IndexAny(trimmedLine, " =")
	key := trimmedLine[:kvSeperatorPos]
	valuesblob := trimmedLine[kvSeperatorPos+1:]

	values, comment, err := decodeValue(valuesblob) //TODO
	if err == errInvalidQuoting {
		err = fmt.Errorf("%q: %w", data, err)
	}

	// TODO: typecheck:

	return rawObject{key, values, comment, string(trimmedLine[kvSeperatorPos]) == "="}, err
}

// decodes an ssh_config valuepart
// does not know of types
//
// possible errors: nil, errInvalidQuoting
func decodeValue(s string) (values []Value, comment string, err error) {
	// func inspired by https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz misc.c#1889 and strings.FieldsFunc()
	strings, currentString, quoted := []Value{}, "", 0
runereader:
	for pos, rune := range s {
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
			pos += 1 // skip next rune
			continue
		}

		switch quoted { // while quoted, ignore the other quoting style
		default:
			panic(fmt.Sprintf("unescapeValue runtime enum: quoted should always be 0..2, is %d", quoted))
		case 1: // single quoted
			if rune == '\'' {
				quoted = 0 // macro a
				strings = append(strings, Value{currentString, 1, ""})
				currentString = ""
				// spaces between are handled below
				continue
			}
		case 2: // double quoted
			if rune == '"' {
				quoted = 0 // duplicate codeblock a
				strings = append(strings, Value{currentString, 2, ""})
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
				strings = append(strings, Value{currentString, 0, ""})
				currentString = ""
				continue
			}
			if rune == '#' {
				comment = s[pos+1:]
				break runereader
			}
		}
		currentString += string(rune)
	}
	if quoted != 0 {
		err = fmt.Errorf("unescapeValues: %q: %w", currentString, errInvalidQuoting)
	}
	return append(strings, Value{currentString, 0, ""}), comment, err
}

// encodes an ssh_config valuepart
// if string begins or ends with space, it is automatically quoted without warning
//
// does not perform type checking
//
// possible errors: nil, errWarnSingleBackslashTransformed
func encodeValue(values []Value, comment string) (encoded string, err error) {
	for i, v := range values {
		if i != 0 {
			encoded += " " // not required, but spaces between arguments is nice
		}

		if v.Quoted == 0 {
			if strings.HasPrefix(v.Value, " ") || strings.HasSuffix(v.Value, " ") {
				v.Quoted = 1
			}
		}

		switch v.Quoted { // macro: quote
		case 1:
			encoded += "'"
		case 2:
			encoded += "\""
		}

		for pos, rune := range v.Value {
			if rune == '\\' { // single backslash
				switch string(v.Value[pos+1]) {
				case "'":
					if v.Quoted != 2 { // 2: already escaped by quotes
						encoded += "\\'" //: \'
					}
				case "\"":
					if v.Quoted != 1 {
						encoded += "\\\"" //: \"
					}
				case "\\":
					encoded += "\\\\"
				default:
					encoded += "\\\\" // 1 backslash gets turned to 2 when read by OpenSSH
					err = errWarnSingleBackslashTransformed
					continue
				}
				pos += 1
				continue
			}
		}

		switch v.Quoted { // macro: quote
		case 1:
			encoded += "'"
		case 2:
			encoded += "\""
		}
	}
	if comment != "" {
		comment = " #" + comment
	}
	return encoded + comment, err
}

// encoding: indentchar

//func typeCheck()
