package pkg

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestRenameNoOverwrite_notexist(t *testing.T) {
	td := t.TempDir()
	src := path.Join(td, "from")
	content := []byte("hello")
	dst := path.Join(td, "to")

	if err := os.WriteFile(src, content, 0600); err != nil {
		t.Fatal(err)
	}

	if err := RenameNoOverwrite(src, dst); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(src); err == nil {
		t.Fatal("src exists‽ it should be renamed to dstempty")
	} else {
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatal(err)
		}
	}

	if read, err := ioutil.ReadFile(dst); err != nil {
		t.Fatal(err)
	} else {
		if string(content) != string(read) {
			t.Fatalf("dstempty: expected %q, got %q", content, read)
		}
	}
}

func TestRenameNoOverwrite_exist(t *testing.T) {
	td := t.TempDir()
	src := path.Join(td, "from")
	src_content := []byte("hello")
	dst := path.Join(td, "to")
	dst_content := []byte("world")

	if err := os.WriteFile(src, src_content, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(dst, dst_content, 0600); err != nil {
		t.Fatal(err)
	}

	// expected: err
	if err := RenameNoOverwrite(src, dst); err == nil {
		t.Fatal("dst exists; should return err‽")
	} else {
		if !errors.Is(err, fs.ErrExist) {
			t.Fatal(err) // other error
		}
	}

	if _, err := os.Stat(src); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			t.Fatal("src not exist‽ it should not have been touched")
		}
		t.Fatal(err)
	}

	if read, err := ioutil.ReadFile(dst); err != nil {
		t.Fatal(err)
	} else {
		if string(dst_content) != string(read) {
			t.Fatalf("existing dst got written to‽ expected %q, got %q", dst_content, read)
		}
	}
}

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
