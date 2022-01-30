package ssh_config

import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// 1b: the raw, no-typing way

func DecodeToRaw(data io.Reader) ([]RawTopLevel, error) {
	var deep bool // under a host or match
	var tl []RawTopLevel
	var cl RawTopLevel // current level
	var i int          // can't integrate it in to for?
	keywordType := reflect.TypeOf(Keywords{})
	keywordsTotal := keywordType.NumField()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		i++
		rkw, err := parseLine(strings.ToValidUTF8(scanner.Text(), ""))
		if err == errInvalidQuoting { // crash and burn
			err = fmt.Errorf("while parsing line %d: %w", i, err)
		}
		if err != nil {
			return tl, err
		}

		switch rkw.Key {
		case "Host", "Match", "Include":
			if rkw.Key == "Include" {
				deep = false
			} else {
				deep = true
			}
			if cl.Key != "" { // flush
				tl = append(tl, cl)
			}
			cl.Key = rkw.Key
			cl.Values = rkw.Values
			cl.Comment = rkw.Comment
			cl.EncodingKVSeperatorIsEquals = rkw.EncodingKVSeperatorIsEquals
			cl.Children = []RawKeyword{}
		default:
			if !deep {
				return tl, fmt.Errorf("while parsing line %d: %w", i, errInvalidKeyLocation)
			}

			// basic 'does key exist'
			var exists bool
			if rkw.Key == "" {
				exists = true // comment line
			}
			for i := 0; i < keywordsTotal; i++ {
				if keywordType.Field(i).Name == rkw.Key {
					exists = true
					break
				}
			}
			if !exists {
				return tl, fmt.Errorf("while parsing line %d: %w", i, errInvalidKeyword)
			}

			cl.Children = append(cl.Children, rkw)
		}
	}
	// flush last
	return append(tl, cl), scanner.Err() // no need to wrap
}

var indent = "  "

func EncodeFromRaw(rawobj []RawTopLevel, data io.Writer) (err error) {
	w := bufio.NewWriter(data)
	defer w.Flush()

	for _, rt := range rawobj {
		switch rt.Key {
		default:
			return fmt.Errorf("while encoding %q: %w", rt.Key, errInvalidKeyword)
		case "Host", "Match", "Include":
			var enline string
			enline, err = encodeLine("", RawKeyword{rt.Key, rt.Values, rt.Comment, rt.EncodingKVSeperatorIsEquals})
			w.WriteString(enline + "\n")

			for _, c := range rt.Children {
				enline, err = encodeLine(indent, c)
				w.WriteString(enline + "\n")
			}
		}
	}
	return err
}

func encodeKVSeperator(sepIsEquals bool) string {
	if sepIsEquals {
		return "="
	} else {
		return " "
	}
}
