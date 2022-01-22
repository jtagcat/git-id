package pkg

import (
	"errors"
	"io/fs"
	"os"
)

// os.Rename() if dst not exist
//
// use: errors.Is(err, fs.ErrExist)
func RenameNoOverwrite(src, dst string) error {
	_, err := os.Stat(dst)
	if err == nil {
		return fs.ErrExist
	}
	if errors.Is(err, os.ErrNotExist) {
		return os.Rename(src, dst)
	}
	return err
}
