package code

import (
	"path/filepath"
	"reflect"
	"testing"
)

func TestParseFileJSON(t *testing.T) {
	got, err := ParseFile(filepath.Join("testdata", "fixture", "file1.json"))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := map[string]any{
		"host":    "hexlet.io",
		"timeout": float64(50),
		"proxy":   "123.234.53.22",
		"follow":  false,
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestParseFileErrors(t *testing.T) {
	if _, err := ParseFile(filepath.Join("testdata", "fixture", "missing.json")); err == nil {
		t.Error("expected error for missing file")
	}
	if _, err := Parse([]byte("a: 1"), ".txt"); err == nil {
		t.Error("expected error for unsupported format")
	}
}
