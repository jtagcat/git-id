package ssh_config

// 1b: the raw, no-typing way
import (
	"bufio"
	"fmt"
	"io"
	"reflect"
	"strings"
)

// wrapper func for non-xkey usage
func DecodeToRaw(data io.Reader) ([]RawTopLevel, error) {
	rootXKeyMap := make(map[string]bool)
	subXKeyMap := make(map[string]bool)
	return DecodeToRawXKeys(data, rootXKeyMap, subXKeyMap)
}

// xkeys: Custom keys nested inside comments.
// rootXKeys: list of root-level xkeys; bool: may have children (recommend default: true)
// subXKeys: list of sub-level xkeys; bool means nothing
func DecodeToRawXKeys(data io.Reader, rootXKeyMap map[string]bool, subXKeyMap map[string]bool) ([]RawTopLevel, error) {
	var deep bool         // under a host or match
	var cfg []RawTopLevel // tree is flusehd to cfg
	var tree RawTopLevel  // current level

	keywordType := reflect.TypeOf(Keywords{})
	keywordKMap := make(map[string]bool)
	for i := 0; i < keywordType.NumField(); i++ {
		keywordKMap[strings.ToLower(keywordType.Field(i).Name)] = false
	}

	var prevLineComment []string // buffer

	scanner := bufio.NewScanner(data)
	for i := 1; scanner.Scan(); i++ {
		line, err := decodeLine(strings.ToValidUTF8(scanner.Text(), "")) // [macro B]
		if err == ErrInvalidQuoting {                                    // crash and burn
			err = fmt.Errorf("while parsing line %d: %w", i, err)
		}
		if err != nil {
			return cfg, err
		}

		// xkey
		var rootXKey, rootXKeyMayHaveChildren, subXKey bool
		if line.Key == "" && line.Comment != "" {
			locaseCmt := strings.ToLower(line.Comment)

			for rk, mayHaveChildren := range rootXKeyMap {
				if strings.HasPrefix(locaseCmt, rk) {
					rootXKey = true
					rootXKeyMayHaveChildren = mayHaveChildren
				}
			}
			for sk := range subXKeyMap {
				if strings.HasPrefix(locaseCmt, sk) {
					subXKey = true
				}
			}

			line, err = decodeLine(strings.ToValidUTF8(line.Comment, "")) // [macro B]
			if err == ErrInvalidQuoting {
				err = fmt.Errorf("while parsing xkey on line %d: %w", i, err)
			}
			if err != nil {
				return cfg, err
			}
		}

		if line.Key == "" {
			if !deep { // already at TLD level
				if i != 1 { // flush previous tree
					cfg = append(cfg, tree)
				}
				tree = RawTopLevel{Comment: line.Comment} // may be an empty line as well; empty comments are trimmed
			} else { // in a tree, buffering to check afterwards whether we hit a subkey (including comment in the tree) or a new tree / EOF
				prevLineComment = append(prevLineComment, line.Comment)
			}
			continue // would-be switch statement (emulating it), if not for the xkey check
		}

		locaseKey := strings.ToLower(line.Key)
		if rootXKey || locaseKey == "host" || locaseKey == "match" || locaseKey == "include" {
			if locaseKey == "include" || rootXKey && !rootXKeyMayHaveChildren {
				deep = false
			} else {
				deep = true
			}
			if i != 1 { // handle attaching comments to trees [macro A]
				for _, c := range prevLineComment { // i != 1 anyway
					if c == "" {
						break // empty line (unless the next key is subkey) moves the line, and subsequent comments to root
					}
					// flush comment to previous tree
					tree.Children = append(tree.Children, RawKeyword{Comment: c})
					prevLineComment = prevLineComment[1:] // remove parsed from buffer
				}
				cfg = append(cfg, tree)             // flush previous tree (ALSO happens when no comments in buffer)
				for _, c := range prevLineComment { // continue as !deep, at root level
					tree = RawTopLevel{Comment: c}
					cfg = append(cfg, tree)
				}
				prevLineComment = []string{} // clear buffer
			}
			// create a new tree (previous already flushed)
			tree = RawTopLevel{line.Key, line.Values, line.Comment, line.EncodingKVSeperatorIsEquals, []RawKeyword{}}
			continue // emulating switch statement
		}

		// any other key will be subkey (switch: default)
		if !deep {
			return cfg, fmt.Errorf("while parsing line %d: %w: TLD %s", i, ErrInvalidKeyword, line.Key)
		}

		// flush comments
		for _, c := range prevLineComment { // i != 1 anyway
			tree.Children = append(tree.Children, RawKeyword{Comment: c})
		}
		prevLineComment = []string{}

		if !subXKey {
			// basic 'does (non-x)key exist'
			if _, ok := keywordKMap[locaseKey]; !ok {
				return cfg, fmt.Errorf("while parsing line %d: %w", i, ErrInvalidKeyword)
			}
		}

		tree.Children = append(tree.Children, line)
	}

	for _, c := range prevLineComment { // flush comments before newline to last tree [macro A]
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

// indent = "  " is reccommended
func EncodeFromRaw(rawobj []RawTopLevel, data io.Writer, indent string) (err error) {
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
