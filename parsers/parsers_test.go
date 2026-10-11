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
		"group3": map[string]any{
			"deep": map[string]any{
				"id": map[string]any{"number": float64(45)},
			},
			"fee": float64(100500),
		},
		"group1": map[string]any{
			"foo":  "bar",
			"baz":  "bars",
			"nest": "str",
		},
		"common": map[string]any{
			"follow":   false,
			"setting1": "Value 1",
			"setting3": nil,
			"setting4": "blah blah",
			"setting5": map[string]any{"key5": "value5"},
			"setting6": map[string]any{
				"key":  "value",
				"ops":  "vops",
				"doge": map[string]any{"wow": "so much"},
			},
		},
	}

	for _, name := range []string{"nested2.json", "nested2.yml"} {
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

func TestParseFileMissing(t *testing.T) {
	_, err := ParseFile(filepath.Join(t.TempDir(), "missing.json"))
	require.Error(t, err)
}

func TestParseErrors(t *testing.T) {
	tests := []struct {
		name    string
		content string
		ext     string
	}{
		{"invalid json", `{"host": `, ".json"},
		{"invalid yaml", "host: [hexlet.io\n", ".yml"},
		{"yaml list instead of map", "- host\n- timeout\n", ".yml"},
		{"unsupported format", "host=hexlet.io", ".txt"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := Parse([]byte(tt.content), tt.ext)
			require.Error(t, err)
		})
	}
}
