package diff

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	data1 := map[string]any{
		"same":    "value",
		"changed": float64(1),
		"removed": true,
		"nested": map[string]any{
			"key":  "old",
			"keep": nil,
		},
		"object": map[string]any{"a": "b"},
	}
	data2 := map[string]any{
		"same":    "value",
		"changed": float64(2),
		"added":   map[string]any{"x": "y"},
		"nested": map[string]any{
			"key":  "new",
			"keep": nil,
		},
		"object": "string",
	}

	expected := []Node{
		{Key: "added", Status: Added, NewValue: map[string]any{"x": "y"}},
		{Key: "changed", Status: Changed, OldValue: float64(1), NewValue: float64(2)},
		{Key: "nested", Status: Nested, Children: []Node{
			{Key: "keep", Status: Unchanged},
			{Key: "key", Status: Changed, OldValue: "old", NewValue: "new"},
		}},
		{Key: "object", Status: Changed, OldValue: map[string]any{"a": "b"}, NewValue: "string"},
		{Key: "removed", Status: Removed, OldValue: true},
		{Key: "same", Status: Unchanged, OldValue: "value", NewValue: "value"},
	}

	assert.Equal(t, expected, Build(data1, data2))
}

func TestBuildEmpty(t *testing.T) {
	assert.Empty(t, Build(map[string]any{}, map[string]any{}))
}
