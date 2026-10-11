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

func TestGenDiffFlatJSON(t *testing.T) {
	tests := []struct {
		name     string
		file1    string
		file2    string
		expected string
	}{
		{"different files", "file1.json", "file2.json", "expected_flat.txt"},
		{"same file", "file1.json", "file1.json", "expected_same.txt"},
		{"all keys removed", "file1.json", "empty.json", "expected_to_empty.txt"},
		{"null value", "file1.json", "null.json", "expected_null.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GenDiff(fixturePath(tt.file1), fixturePath(tt.file2), "stylish")
			require.NoError(t, err)
			assert.Equal(t, readFixture(t, tt.expected), got)
		})
	}
}

func TestGenDiffEmptyFiles(t *testing.T) {
	got, err := GenDiff(fixturePath("empty.json"), fixturePath("empty.json"), "stylish")
	require.NoError(t, err)
	assert.Equal(t, "{\n}", got)
}

func TestGenDiffAbsolutePaths(t *testing.T) {
	path1, err := filepath.Abs(fixturePath("file1.json"))
	require.NoError(t, err)
	path2, err := filepath.Abs(fixturePath("file2.json"))
	require.NoError(t, err)

	got, err := GenDiff(path1, path2, "stylish")
	require.NoError(t, err)
	assert.Equal(t, readFixture(t, "expected_flat.txt"), got)
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
