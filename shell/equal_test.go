package shell

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEqual(t *testing.T) {
	cases := []struct {
		file1 string
		file2 string
	}{
		{
			file1: "messy.sh",
			file2: "messy.sh",
		},
		{
			file1: "messy.sh",
			file2: "clean.sh",
		},
		{
			file1: "clean.sh",
			file2: "messy.sh",
		},
		{
			file1: "clean.sh",
			file2: "messy.sh",
		},
		{
			file1: "clean.sh",
			file2: "clean.sh",
		},
		{
			file1: "empty.sh",
			file2: "empty.sh",
		},
	}

	for i, test := range cases {
		f1, err1 := os.Open(filepath.Join("fixtures", test.file1))
		f2, err2 := os.Open(filepath.Join("fixtures", test.file2))
		if err1 != nil {
			t.Fatalf("testcase %d: unable to open %q: %s", i, test.file1, err1)
		}
		if err2 != nil {
			t.Fatalf("testcase %d: unable to open %q: %s", i, test.file2, err2)
		}
		if !Equal(f1, f2) {
			t.Errorf("testcase %d: %q and %q are not equal", i, test.file1, test.file2)
		}
		f1.Close()
		f2.Close()
	}
}

func TestNotEqual(t *testing.T) {
	cases := []struct {
		file1 string
		file2 string
	}{
		{
			file1: "messy.sh",
			file2: "empty.sh",
		},
		{
			file1: "clean.sh",
			file2: "clean2.sh",
		},
	}

	for i, test := range cases {
		f1, err1 := os.Open(filepath.Join("fixtures", test.file1))
		f2, err2 := os.Open(filepath.Join("fixtures", test.file2))
		if err1 != nil {
			t.Fatalf("testcase %d: unable to open %q: %s", i, test.file1, err1)
		}
		if err2 != nil {
			t.Fatalf("testcase %d: unable to open %q: %s", i, test.file2, err2)
		}
		if Equal(f1, f2) {
			t.Errorf("testcase %d: %q and %q are equal", i, test.file1, test.file2)
		}
		f1.Close()
		f2.Close()
	}
}
