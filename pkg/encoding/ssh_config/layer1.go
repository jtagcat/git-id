package ssh_config

// 1b: the raw, no-typing way
import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"

	"github.com/google/renameio/v2"
)

type Config struct {
	cfg  []RawTopLevel
	o    Opts
	path string
}

type Opts struct {
	// (optional) xkeys are custom keys nested inside comments,
	// extending the configuration specification.

	// valid root-level xkeys, at root level they have no parent
	RootXKeys map[string]bool // bool: may have children xkeys

	// valid children xkeys, housed under a root/parent (x)key
	SubXKeys []string

	// Indentation to use when encoding
	Indent string // standard: "  "
}

func OpenConfig(o Opts, path string) (Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return Config{}, err
	}
	defer f.Close()

	cfg, err := Decode(o, bufio.NewReader(f))

	return Config{
		cfg:  cfg,
		o:    o,
		path: path,
	}, err
}

func (c *Config) Write() error {
	// atomically replace
	f, err := renameio.NewPendingFile(c.path)
	if err != nil {
		return err
	}

	err = Encode(c.o, c.cfg, bufio.NewWriter(f))
	if err != nil {
		return err
	}

	return f.CloseAtomicallyReplace()
}

// Decode from ssh_config to Config
func Decode(o Opts, data io.Reader) ([]RawTopLevel, error) {
	var deep bool           // under a host or match
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

			for rk, mayHaveChildren := range o.RootXKeys {
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

			if rootXKey || subXKey { // parse xkey comment to key
				line, err = decodeLine(strings.ToValidUTF8(line.Comment, "")) // [macro B]
				if err == ErrInvalidQuoting {
					err = fmt.Errorf("while parsing xkey on line %d: %w", i, err)
				}
				if err != nil {
					return cfg, err
				}
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
		if rootXKey || locaseKey == "host" || locaseKey == "match" || locaseKey == "include" && !includeIsChild {
			if !rootXKey && locaseKey != "include" {
				includeIsChild = true // include is a subkey after host/match is encountered
			}

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
			tree = RawTopLevel{Key: line.Key, Values: line.Values, Comment: line.Comment, EncodingKVSeperatorIsEquals: line.EncodingKVSeperatorIsEquals}
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
			// basic 'does (non-x)key exist' [macro D]
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

// Encode from ssh_config to Config
func Encode(o Opts, cfg []RawTopLevel, data io.Writer) error {
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
	var warn error
	for _, r := range cfg {
		locaseRK := strings.ToLower(r.Key)
		var enline string

		_, isRootXKey := o.RootXKeys[locaseRK]
		if isRootXKey || r.Key == "" || locaseRK == "host" || locaseRK == "match" || locaseRK == "include" && !includeIsChild {
			if !isRootXKey && locaseRK != "" && locaseRK != "include" {
				includeIsChild = true
			}
			if isRootXKey {
				x := r
				r = RawTopLevel{}
				r.Comment, _ = encodeLine("", RawKeyword{x.Key, x.Values, x.Comment, x.EncodingKVSeperatorIsEquals})
			}

			enline, warn = encodeLine("", RawKeyword{r.Key, r.Values, r.Comment, r.EncodingKVSeperatorIsEquals})
			w.WriteString(enline + "\n")

			for _, c := range r.Children {
				locaseCK := strings.ToLower(c.Key)
				if _, isSubXKey := subXKeyMap[locaseCK]; isSubXKey {
					x := c
					c = RawKeyword{}
					c.Comment, _ = encodeLine("", x)

				} else if c.Key != "" { // basic 'does (non-x)key exist' [macro D]
					if _, ok := keywordKMap[locaseCK]; !ok {
						return fmt.Errorf("while encoding %q: %w", c.Key, ErrInvalidKeyword)
					}
				}

				enline, warn = encodeLine(o.Indent, c)
				w.WriteString(enline + "\n")
			}
			continue
		}
		return fmt.Errorf("while encoding %q: %w", r.Key, ErrInvalidKeyword)
	}
	return warn
}
