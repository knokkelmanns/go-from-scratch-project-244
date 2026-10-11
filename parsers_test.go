package code

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFileJSON(t *testing.T) {
	got, err := ParseFile(fixturePath("file1.json"))
	require.NoError(t, err)

	expected := map[string]any{
		"host":    "hexlet.io",
		"timeout": float64(50),
		"proxy":   "123.234.53.22",
		"follow":  false,
	}
	assert.Equal(t, expected, got)
}

func TestParseJSON(t *testing.T) {
	got, err := Parse([]byte(`{"a": 1}`), ".json")
	require.NoError(t, err)
	assert.Equal(t, map[string]any{"a": float64(1)}, got)
}

func TestParseErrors(t *testing.T) {
	_, err := ParseFile(fixturePath("missing.json"))
	require.Error(t, err)

	_, err = Parse([]byte(`{"a": `), ".json")
	require.Error(t, err)

	_, err = Parse([]byte("a: 1"), ".txt")
	require.Error(t, err)
}
