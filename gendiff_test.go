package code

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixturePath(name string) string {
	return filepath.Join("testdata", "fixture", name)
}

func readFixture(t *testing.T, name string) string {
	t.Helper()

	content, err := os.ReadFile(fixturePath(name))
	require.NoError(t, err)

	return string(content)
}

func TestGenDiffStylish(t *testing.T) {
	tests := []struct {
		name   string
		file1  string
		file2  string
		format string
	}{
		{"json files", "nested1.json", "nested2.json", "stylish"},
		{"yaml files", "nested1.yml", "nested2.yml", "stylish"},
		{"json and yaml", "nested1.json", "nested2.yml", "stylish"},
		{"default format", "nested1.json", "nested2.json", ""},
	}

	expected := readFixture(t, "expected_stylish.txt")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenDiff(fixturePath(tt.file1), fixturePath(tt.file2), tt.format)
			require.NoError(t, err)
			assert.Equal(t, expected, got)
		})
	}
}

func TestGenDiffPlain(t *testing.T) {
	tests := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"json files", "nested1.json", "nested2.json"},
		{"yaml files", "nested1.yml", "nested2.yml"},
	}

	expected := readFixture(t, "expected_plain.txt")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenDiff(fixturePath(tt.file1), fixturePath(tt.file2), "plain")
			require.NoError(t, err)
			assert.Equal(t, expected, got)
		})
	}
}

func TestGenDiffJSON(t *testing.T) {
	tests := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"json files", "nested1.json", "nested2.json"},
		{"yaml files", "nested1.yml", "nested2.yml"},
	}

	expected := readFixture(t, "expected_json.json")
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenDiff(fixturePath(tt.file1), fixturePath(tt.file2), "json")
			require.NoError(t, err)
			assert.JSONEq(t, expected, got)
		})
	}
}

func TestGenDiffEmptyFiles(t *testing.T) {
	got, err := GenDiff(fixturePath("empty.json"), fixturePath("empty.json"), "stylish")
	require.NoError(t, err)
	assert.Equal(t, "{\n}", got)
}

func TestGenDiffAbsolutePaths(t *testing.T) {
	path1, err := filepath.Abs(fixturePath("nested1.json"))
	require.NoError(t, err)
	path2, err := filepath.Abs(fixturePath("nested2.yml"))
	require.NoError(t, err)

	got, err := GenDiff(path1, path2, "stylish")
	require.NoError(t, err)
	assert.Equal(t, readFixture(t, "expected_stylish.txt"), got)
}

func TestGenDiffErrors(t *testing.T) {
	tests := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"first file missing", "missing.json", "file2.json"},
		{"second file missing", "file1.json", "missing.json"},
		{"invalid json", "invalid.json", "file2.json"},
		{"invalid yaml", "file1.yml", "invalid.yml"},
		{"unsupported format", "file.txt", "file2.json"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenDiff(fixturePath(tt.file1), fixturePath(tt.file2), "stylish")
			require.Error(t, err)
		})
	}
}

func TestGenDiffUnknownFormat(t *testing.T) {
	_, err := GenDiff(fixturePath("file1.json"), fixturePath("file2.json"), "xml")
	require.Error(t, err)
}
