package ssh_config

import (
	"bufio"
	"os"

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
