package pkg

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestOpenFileExisting_notexist(t *testing.T) {
	tf := path.Join(t.TempDir(), "file")

	// expected: err
	f, err := OpenFileExisting(tf, os.O_RDONLY)
	if err == nil {
		t.Fatal("tf does not exist, should return err‽")
	}
	if !errors.Is(err, fs.ErrNotExist) {
		t.Fatal(err)
	}
	defer f.Close()
}

func TestOpenFileExisting_exist(t *testing.T) {
	tf := path.Join(t.TempDir(), "file")
	if err := os.WriteFile(tf, []byte(""), 0600); err != nil {
		t.Fatal(err)
	}

	f, err := OpenFileExisting(tf, os.O_RDONLY)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
}

func TestFileAppend(t *testing.T) {
	tf := path.Join(t.TempDir(), "file")
	content1 := []byte("hello")
	content2 := []byte("world")

	if err := os.WriteFile(tf, content1, 0600); err != nil {
		t.Fatal(err)
	}

	if err := FileAppend(tf, content2); err != nil {
		t.Fatal(err)
	}

	if read, err := ioutil.ReadFile(tf); err != nil {
		t.Fatal(err)
	} else {
		want := string(content1) + string(content2)
		if want != string(read) {
			t.Fatalf("expected %q, got %q", want, read)
		}
	}
}
