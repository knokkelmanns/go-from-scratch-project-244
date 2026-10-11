package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/diff"
)

var tree = []diff.Node{
	{Key: "added", Status: diff.Added, NewValue: map[string]any{
		"b": map[string]any{"c": nil},
		"a": float64(12345),
	}},
	{Key: "group", Status: diff.Nested, Children: []diff.Node{
		{Key: "empty", Status: diff.Changed, OldValue: "", NewValue: "so much"},
		{Key: "list", Status: diff.Changed, OldValue: []any{float64(1)}, NewValue: nil},
		{Key: "same", Status: diff.Unchanged, OldValue: true, NewValue: true},
	}},
	{Key: "removed", Status: diff.Removed, OldValue: float64(1.5)},
}

func TestFormat(t *testing.T) {
	tests := []struct {
		format   string
		expected string
	}{
		{"", expectedStylish},
		{"stylish", expectedStylish},
		{"plain", expectedPlain},
		{"json", expectedJSON},
	}

	for _, tt := range tests {
		t.Run("format "+tt.format, func(t *testing.T) {
			got, err := Format(tree, tt.format)
			require.NoError(t, err)
			assert.Equal(t, tt.expected, got)
		})
	}
}

func TestFormatUnknown(t *testing.T) {
	_, err := Format(tree, "xml")
	require.Error(t, err)
}
