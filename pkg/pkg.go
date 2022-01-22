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

// os.WriteFile, but fs.ErrExist
func WriteFileExisting(name string, data []byte) error {
	if _, err := os.Stat(name); err == nil {
		return os.WriteFile(name, data, 0000)
	}
	return fs.ErrExist
}
