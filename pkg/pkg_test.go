package pkg

import (
	"errors"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestRenameNoOverwrite(t *testing.T) {
	td := t.TempDir()
	src := path.Join(td, "from")
	src_content := []byte("hello")
	dstexists := path.Join(td, "to")
	dstexists_content := []byte("world")
	dstempty := path.Join(td, "notexist")

	if err := ioutil.WriteFile(src, src_content, 0600); err != nil {
		t.Fatal(err)
	}
	if err := ioutil.WriteFile(dstexists, dstexists_content, 0600); err != nil {
		t.Fatal(err)
	}

	// 1st run, shall fail
	if err := RenameNoOverwrite(src, dstexists); err == nil {
		t.Fatal(err) // dst exists; should fail‽
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

	if read, err := ioutil.ReadFile(dstexists); err != nil {
		t.Fatal(err)
	} else {
		if string(dstexists_content) != string(read) {
			t.Fatalf("existing dst got written to; expected %q, got %q", dstexists_content, read)
		}
	}

	// 2nd run, shall succeed
	if err := RenameNoOverwrite(src, dstempty); err != nil {
		t.Fatal(err)
	}

	if _, err := os.Stat(src); err == nil {
		t.Fatal("src exists‽ it should be renamed to dstempty")
	} else {
		if !errors.Is(err, os.ErrNotExist) {
			t.Fatal(err)
		}
	}

	if read, err := ioutil.ReadFile(dstempty); err != nil {
		t.Fatal(err)
	} else {
		if string(src_content) != string(read) {
			t.Fatalf("dstempty: expected %q, got %q", src_content, read)
		}
	}
}
