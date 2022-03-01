package pkg

import (
	"errors"
	"io/fs"
	"os"
)

// os.OpenFile, but fs.ErrNotExist
func OpenFileExisting(name string, flag int) (*os.File, error) {
	if _, err := os.Stat(name); err == nil {
		return os.OpenFile(name, flag, 0000)
	}
	return nil, fs.ErrNotExist
}

// append to file with fs.ErrNotExist
//
// WARN: perhaps behaviour with async
func FileAppend(name string, b []byte) error {
	f, err := OpenFileExisting(name, os.O_APPEND|os.O_WRONLY)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}

// returns path, or $PWD if path empty
func PWDIfEmpty(path string) (string, error) {
	if path != "" {
		return path, nil
	}
	wd, err := os.Getwd()
	if err == nil {
		return wd, nil
	}
	return "", errors.New("path unset, failed getting working directory: " + err.Error())
}

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
