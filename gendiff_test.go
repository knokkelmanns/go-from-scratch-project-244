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

func writeTempFile(t *testing.T, name, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), name)
	require.NoError(t, os.WriteFile(path, []byte(content), 0o600))

	return path
}

func TestGenDiffEmptyFiles(t *testing.T) {
	path := writeTempFile(t, "empty.json", "{}")

	got, err := GenDiff(path, path, "stylish")
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
	valid := fixturePath("nested1.json")
	missing := filepath.Join(t.TempDir(), "missing.json")

	tests := []struct {
		name  string
		file1 string
		file2 string
	}{
		{"first file missing", missing, valid},
		{"second file missing", valid, missing},
		{"invalid json", writeTempFile(t, "invalid.json", `{"host": `), valid},
		{"invalid yaml", valid, writeTempFile(t, "invalid.yml", "host: [hexlet.io\n")},
		{"unsupported format", writeTempFile(t, "file.txt", "host=hexlet.io"), valid},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GenDiff(tt.file1, tt.file2, "stylish")
			require.Error(t, err)
		})
	}
}

func TestGenDiffUnknownFormat(t *testing.T) {
	_, err := GenDiff(fixturePath("nested1.json"), fixturePath("nested2.json"), "xml")
	require.Error(t, err)
}
