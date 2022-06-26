package ssh_config

// Handles conversion between raw text and raw structurized lines

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type RawKeyword struct {
	Key    string
	Values []RawValue // when key set, len(Values) >= 1
	// "# foobar" → " foobar", note the leading space
	Comment string // at the end of same line as Key

	EncodingKVSeperatorIsEquals bool // "Key=Value" instead of "Key Value"
}

// Parse an ssh_config line without type checking.
// rawKeyword might be a TLD object;
//
// possible errors: nil, errInvalidQuoting
func decodeLine(data string) (RawKeyword, error) {
	trimmedLine := strings.TrimSpace(data)
	if trimmedLine == "" {
		return RawKeyword{}, nil
	}
	if strings.HasPrefix(trimmedLine, "#") {
		// dedicated comment line
		return RawKeyword{Comment: strings.TrimPrefix(trimmedLine, "#")}, nil
	}

	key, valuesblob, kvSeperator := CutAny(trimmedLine, " =")

	values, comment, err := DecodeValue(valuesblob)
	if errors.Is(err, ErrInvalidQuoting) {
		err = fmt.Errorf("%q: %w", data, err)
	}

	return RawKeyword{key, values, comment, kvSeperator == "="}, err
}

// Encodes an ssh_config line without type checking.
// rawKeyword might be a TLD object
//
// possible errors: nil, errWarnSingleBackslashTransformed
func encodeLine(indent string, rkw RawKeyword) (string, error) {
	if rkw.Key == "" && rkw.Comment == "" {
		return "", nil
	}
	keyPart := indent

	if rkw.Key != "" && rkw.Values != nil {
		keyPart += rkw.Key
		if rkw.EncodingKVSeperatorIsEquals {
			keyPart += "="
		} else {
			keyPart += " "
		}
	}

	valuePart, err := EncodeValue(rkw.Values, rkw.Comment)
	return keyPart + valuePart, err
}

type RawValue struct {
	Value  string
	Quoted int // enum: 0: not, 1: single, 2: double
}

// decodes an ssh_config valuepart
// does not know of types
//
// possible errors: nil, errInvalidQuoting
func DecodeValue(s string) (strings []RawValue, comment string, err error) {
	// func inspired by https://ftp.openbsd.org/pub/OpenBSD/OpenSSH/openssh-8.8.tar.gz misc.c#1889 and strings.FieldsFunc()
	currentString, quoted := "", 0
	maxPos := utf8.RuneCountInString(s) - 1 // prevents index (of next rune) out of range

runereader:
	for pos, rune := range s {
		if rune == '\\' { // single backslash
			if pos == maxPos { // last rune
				currentString += "\\\\" // 2 backslashes
				continue
			}

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
			pos++ // skip next rune
			continue
		}

		switch quoted { // while quoted, ignore the other quoting style
		default:
			panic(fmt.Sprintf("unescapeValue runtime enum: quoted should always be 0..2, is %d", quoted))
		case 1: // single quoted
			if rune == '\'' {
				quoted = 0 // macro a
				strings = append(strings, RawValue{currentString, 1})
				currentString = ""
				// spaces between are handled below
				continue
			}
		case 2: // double quoted
			if rune == '"' {
				quoted = 0 // duplicate codeblock a
				strings = append(strings, RawValue{currentString, 2})
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
				strings = append(strings, RawValue{currentString, 0})
				currentString = ""
				continue
			}
			if rune == '#' {
				if pos == maxPos { // last rune
					comment = ""
				} else {
					comment = s[pos+1:]
				}

				break runereader
			}
		}
		currentString += string(rune)
	}
	if quoted != 0 {
		return nil, "", fmt.Errorf("unescapeValues: %q: %w", currentString, ErrInvalidQuoting)
	}
	if currentString != "" {
		strings = append(strings, RawValue{currentString, 0})
	}
	return strings, comment, nil
}

// encodes an ssh_config valuepart
// if string begins or ends with space, it is automatically quoted without warning
//
// does not perform type checking
//
// possible errors: nil, errWarnSingleBackslashTransformed
func EncodeValue(values []RawValue, comment string) (encoded string, err error) {
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
					err = ErrWarnSingleBackslashTransformed
					continue
				}
				pos++
				continue
			}
			encoded += string(rune)
		}

		switch v.Quoted { // macro: quote
		case 1:
			encoded += "'"
		case 2:
			encoded += "\""
		}
	}

	if comment != "" {
		comment = "#" + comment
		if len(values) != 0 { // prettify
			comment = " " + comment
		}
	}
	return encoded + comment, err
}
