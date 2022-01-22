package pkg

import (
	"errors"
	"io/fs"
	"os"
)

// os.Rename(), but newpath:fs.ErrExist
func RenameNoOverwrite(oldpath, newpath string) error {
	_, err := os.Stat(newpath)
	if err == nil {
		return fs.ErrExist
	}
	if errors.Is(err, os.ErrNotExist) {
		return os.Rename(oldpath, newpath)
	}
	return err
}

// os.OpenFile, but fs.ErrNotExist
func OpenFileExisting(name string, flag int) (*os.File, error) {
	if _, err := os.Stat(name); err == nil {
		return os.OpenFile(name, flag, 0000)
	}
	return nil, fs.ErrNotExist
}
