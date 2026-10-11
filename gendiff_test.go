package code

import (
	"path/filepath"
	"testing"
)

func TestGenDiffFlatJSON(t *testing.T) {
	got, err := GenDiff(
		filepath.Join("testdata", "fixture", "file1.json"),
		filepath.Join("testdata", "fixture", "file2.json"),
		"stylish",
	)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := `{
  - follow: false
    host: hexlet.io
  - proxy: 123.234.53.22
  - timeout: 50
  + timeout: 20
  + verbose: true
}`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s", got, want)
	}
}
