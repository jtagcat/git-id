package ssh_config

// Handles conversion between raw structured lines and raw structures.

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"reflect"
	"strings"
)

type RawTopLevel struct {
	Key string // enum(4): "" (comment / empty line), Host, Match,
	//                     Include: not recursed, nothing is done (no Children)
	Values []RawValue
	// "# foobar" → " foobar", note the leading space
	Comment                     string
	EncodingKVSeperatorIsEquals bool // "Key=Value" instead of "Key Value"

	Children []RawKeyword
}

// Decode from ssh_config to Config.
//
// Host, Match and Include are hardcoded since there are only 3, and Include means different things in different contexts.
func Decode(o Opts, data io.Reader) ([]RawTopLevel, error) {
	var deep bool           // under a host or match or rootxkey
	var deepXRoot bool      // under a rootxkey
	var includeIsChild bool // if any other root header has been encountered before include, it is a subkey
	var cfg []RawTopLevel   // tree is flusehd to cfg
	var tree RawTopLevel    // current level

	keywordType := reflect.TypeOf(Keywords{})
	keywordKMap := make(map[string]bool)
	for i := 0; i < keywordType.NumField(); i++ {
		keywordKMap[strings.ToLower(keywordType.Field(i).Name)] = false
	}
	subXKeyMap := make(map[string]bool)
	for _, k := range o.SubXKeys {
		subXKeyMap[strings.ToLower(k)] = false
	}

	var prevLineComment []string // buffer

	scanner := bufio.NewScanner(data)
	for i := 1; scanner.Scan(); i++ {
		line, err := decodeLine(strings.ToValidUTF8(scanner.Text(), "")) // [macro B]
		if errors.Is(err, ErrInvalidQuoting) {                           // crash and burn
			err = fmt.Errorf("while parsing line %d: %w", i, err)
		}
		if err != nil {
			return cfg, err
		}

		// xkey
		var rootXKey, rootXKeyMayHaveChildren, subXKey bool
		if line.Key == "" && line.Comment != "" {
			locaseCmt := strings.ToLower(line.Comment)

			for rk, mayHaveChildren := range o.RootXKeys {
				if strings.HasPrefix(locaseCmt, rk) {
					rootXKey = true
					rootXKeyMayHaveChildren = mayHaveChildren
				}
			}
			if !rootXKey {
				for sk := range subXKeyMap {
					if strings.HasPrefix(locaseCmt, sk) {
						subXKey = true
					}
				}
			}

			if rootXKey || subXKey { // parse xkey comment to key
				line, err = decodeLine(line.Comment) // [macro B]
				if errors.Is(err, ErrInvalidQuoting) {
					err = fmt.Errorf("while parsing xkey on line %d: %w", i, err)
				}
				if err != nil {
					return cfg, err
				}
			}
		}

		// comments find-your-parent
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

		// parse a root tree, validate key
		locaseKey := strings.ToLower(line.Key)
		if rootXKey || locaseKey == "host" || locaseKey == "match" || (locaseKey == "include" && !includeIsChild) {
			if !rootXKey && locaseKey != "include" {
				includeIsChild = true // include is a subkey after host/match is encountered
			}

			if rootXKeyMayHaveChildren {
				deepXRoot = true
			}
			if locaseKey == "include" || (rootXKey && !rootXKeyMayHaveChildren) {
				deep, deepXRoot = false, false
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
			tree = RawTopLevel{Key: line.Key, Values: line.Values, Comment: line.Comment, EncodingKVSeperatorIsEquals: line.EncodingKVSeperatorIsEquals}
			continue
		}

		// not rootkey, might be subkey

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
			if deepXRoot {
				return cfg, fmt.Errorf("while encoding %q: %w", i, ErrValidSubkeyAfterXRoot)
			}
			// basic 'does (non-x)key exist' [macro D]
			if _, ok := keywordKMap[locaseKey]; !ok {
				return cfg, fmt.Errorf("while parsing line %d: %w: %s", i, ErrInvalidKeyword, line.Key)
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

// Encode from ssh_config to Config
func Encode(o Opts, cfg []RawTopLevel, data io.Writer) (err error) {
	keywordType := reflect.TypeOf(Keywords{})
	keywordKMap := make(map[string]bool)
	for i := 0; i < keywordType.NumField(); i++ {
		keywordKMap[strings.ToLower(keywordType.Field(i).Name)] = false
	}

	// xKeys are encoded in this func because indenting
	subXKeyMap := make(map[string]bool)
	for _, k := range o.SubXKeys {
		subXKeyMap[strings.ToLower(k)] = false
	}

	w := bufio.NewWriter(data)
	defer w.Flush()

	var includeIsChild bool
	for _, r := range cfg {
		locaseRK := strings.ToLower(r.Key)
		var enline string

		_, isRootXKey := o.RootXKeys[locaseRK]
		if isRootXKey || r.Key == "" || locaseRK == "host" || locaseRK == "match" || (locaseRK == "include" && !includeIsChild) {
			if !isRootXKey && locaseRK != "" && locaseRK != "include" {
				includeIsChild = true
			}

			rc := r // copy for rootobj reset + children looping
			if isRootXKey {
				r = RawTopLevel{}
				r.Comment, err = encodeLine("", RawKeyword{rc.Key, rc.Values, rc.Comment, rc.EncodingKVSeperatorIsEquals})
				if err != nil {
					return fmt.Errorf("while encoding rootXKey: %w", err)
				}
			}

			enline, err = encodeLine("", RawKeyword{r.Key, r.Values, r.Comment, r.EncodingKVSeperatorIsEquals})
			if _, err := w.WriteString(enline + "\n"); err != nil {
				return fmt.Errorf("while Encode'ing to io.Writer: %w", err)
			}

			for _, c := range rc.Children {
				locaseCK := strings.ToLower(c.Key)
				if _, isSubXKey := subXKeyMap[locaseCK]; isSubXKey {
					x := c
					c = RawKeyword{}
					c.Comment, _ = encodeLine("", x)

				} else if c.Key != "" { // not comment
					if isRootXKey {
						return fmt.Errorf("while encoding %q: %w", c.Key, ErrValidSubkeyAfterXRoot)
					}
					// basic 'does (non-x)key exist' [macro D]
					if _, ok := keywordKMap[locaseCK]; !ok {
						return fmt.Errorf("while encoding %q: %w", c.Key, ErrInvalidKeyword)
					}
				}

				enline, err = encodeLine(o.Indent, c)
				if _, err := w.WriteString(enline + "\n"); err != nil {
					return fmt.Errorf("while Encode'ing to io.Writer: %w", err)
				}
			}
			continue
		}
		return fmt.Errorf("while encoding TLD %q: %w", r.Key, ErrInvalidKeyword)
	}
	return err
}
