package parsers

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func fixturePath(name string) string {
	return filepath.Join("..", "testdata", "fixture", name)
}

func TestParseFile(t *testing.T) {
	expected := map[string]any{
		"host":    "hexlet.io",
		"timeout": float64(50),
		"proxy":   "123.234.53.22",
		"follow":  false,
	}

	for _, name := range []string{"file1.json", "file1.yml"} {
		t.Run(name, func(t *testing.T) {
			got, err := ParseFile(fixturePath(name))
			require.NoError(t, err)
			assert.Equal(t, expected, got)
		})
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		content string
		ext     string
	}{
		{"json", `{"a": 1, "b": null}`, ".json"},
		{"yml", "a: 1\nb: null\n", ".yml"},
		{"yaml", "a: 1\nb: null\n", ".yaml"},
		{"upper case extension", "a: 1\nb: null\n", ".YML"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse([]byte(tt.content), tt.ext)
			require.NoError(t, err)
			assert.Equal(t, map[string]any{"a": float64(1), "b": nil}, got)
		})
	}
}

func TestParseYAMLNestedNumbers(t *testing.T) {
	content := "nested:\n  list: [1, 2]\n  big: 18446744073709551615\n"
	got, err := Parse([]byte(content), ".yml")
	require.NoError(t, err)

	expected := map[string]any{
		"nested": map[string]any{
			"list": []any{float64(1), float64(2)},
			"big":  float64(18446744073709551615),
		},
	}
	assert.Equal(t, expected, got)
}

func TestParseEmptyYAML(t *testing.T) {
	got, err := Parse([]byte(""), ".yml")
	require.NoError(t, err)
	assert.Empty(t, got)
}

func TestParseFileErrors(t *testing.T) {
	for _, name := range []string{"missing.json", "invalid.json", "invalid.yml", "list.yml", "file.txt"} {
		t.Run(name, func(t *testing.T) {
			_, err := ParseFile(fixturePath(name))
			require.Error(t, err)
		})
	}
}
