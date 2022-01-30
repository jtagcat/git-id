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
	var deep bool         // under a host or match
	var cfg []RawTopLevel // tree is flusehd to cfg
	var tree RawTopLevel  // current level

	keywordType := reflect.TypeOf(Keywords{})
	keywordsTotal := keywordType.NumField()

	var prevLineComment []string // buffer f

	scanner := bufio.NewScanner(data)
	for i := 1; scanner.Scan(); i++ {
		line, err := decodeLine(strings.ToValidUTF8(scanner.Text(), ""))
		if err == ErrInvalidQuoting { // crash and burn
			err = fmt.Errorf("while parsing line %d: %w", i, err)
		}
		if err != nil {
			return cfg, err
		}

		switch line.Key {
		case "":
			if !deep { // TLD
				if i != 1 { // flush previous [macro C]
					cfg = append(cfg, tree)
				}
				tree = RawTopLevel{Comment: line.Comment} // may be an empty line as well; empty comments are trimmed
			} else { // in a tree, buffering to see whether we hit a new TLD or subkey (in the same tree) first.
				prevLineComment = append(prevLineComment, line.Comment)
			}
		case "Host", "Match", "Include":
			if line.Key == "Include" {
				deep = false
			} else {
				deep = true
			}
			if i != 1 {
				for _, c := range prevLineComment { // i != 1 anyway
					if c == "" {
						break
					}
					tree.Children = append(tree.Children, RawKeyword{Comment: c})
					prevLineComment = prevLineComment[1:]
				}
				cfg = append(cfg, tree)
				for _, c := range prevLineComment { // continue as !deep
					tree = RawTopLevel{Comment: c}
					cfg = append(cfg, tree)
				}
				prevLineComment = []string{}
			}
			tree = RawTopLevel{line.Key, line.Values, line.Comment, line.EncodingKVSeperatorIsEquals, []RawKeyword{}}
		default:
			if !deep {
				return cfg, fmt.Errorf("while parsing line %d: %w: TLD %s", i, ErrInvalidKeyword, line.Key)
			}

			// flush comments
			for _, c := range prevLineComment { // i != 1 anyway
				tree.Children = append(tree.Children, RawKeyword{Comment: c})
			}
			prevLineComment = []string{}

			// basic 'does key exist'
			var exists bool
			for i := 0; i < keywordsTotal; i++ {
				if strings.EqualFold(keywordType.Field(i).Name, line.Key) {
					exists = true
					break
				}
			}
			if !exists {
				return cfg, fmt.Errorf("while parsing line %d: %w", i, ErrInvalidKeyword)
			}

			tree.Children = append(tree.Children, line)
		}
	}
	for _, c := range prevLineComment { // add comments before newline to last tree
		if c == "" {
			break
		}
		tree.Children = append(tree.Children, RawKeyword{Comment: c})
		prevLineComment = prevLineComment[1:]
	}
	cfg = append(cfg, tree)
	for _, c := range prevLineComment { // continue as !deep
		tree = RawTopLevel{Comment: c}
		cfg = append(cfg, tree)
	}
	return cfg, scanner.Err()
}

var indent = "  "

func EncodeFromRaw(rawobj []RawTopLevel, data io.Writer) (err error) {
	w := bufio.NewWriter(data)
	defer w.Flush()

	for _, rt := range rawobj {
		var enline string
		switch rt.Key {
		default:
			return fmt.Errorf("while encoding %q: %w", rt.Key, ErrInvalidKeyword)
		case "Host", "Match", "Include":
			enline, err = encodeLine("", RawKeyword{rt.Key, rt.Values, rt.Comment, rt.EncodingKVSeperatorIsEquals})
			w.WriteString(enline + "\n")

			for _, c := range rt.Children {
				enline, err = encodeLine(indent, c)
				w.WriteString(enline + "\n")
			}
		case "":
			if rt.Comment == "" {
				enline, err = encodeLine("", RawKeyword{Comment: rt.Comment})
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
