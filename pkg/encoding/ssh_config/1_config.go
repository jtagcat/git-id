package ssh_config

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/google/renameio/v2"
	"github.com/mitchellh/go-homedir"
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

// bool: new file created
func OpenConfig(o Opts, name string) (*Config, bool, error) {
	path, err := homedir.Expand(name)
	if err != nil {
		return nil, false, err
	}

	var init bool
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		init = true
	} else if err != nil {
		return nil, false, fmt.Errorf("couldn't stat config at %s: %w", path, err)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, init, fmt.Errorf("couldn't open config file at %s: %w", path, err)
	}
	defer f.Close()

	cfg, err := Decode(o, bufio.NewReader(f))

	return &Config{
		cfg:  cfg,
		o:    o,
		path: path,
	}, init, err
}

func (c *Config) Write() error {
	// atomically replace
	f, err := renameio.NewPendingFile(c.path)
	if err != nil {
		return err
	}

	w := bufio.NewWriter(f)

	if err := Encode(c.o, c.cfg, w); err != nil {
		return err
	}

	if err := w.Flush(); err != nil {
		return err
	}
	return f.CloseAtomicallyReplace()
}
