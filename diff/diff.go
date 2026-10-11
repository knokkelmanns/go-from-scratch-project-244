// Package diff builds the internal representation of the difference
// between two decoded configuration files.
package diff

import (
	"maps"
	"reflect"
	"slices"
)

// Status describes how a key changed between the two structures.
type Status string

// Possible statuses of a diff node.
const (
	Added     Status = "added"
	Removed   Status = "removed"
	Unchanged Status = "unchanged"
	Changed   Status = "changed"
	Nested    Status = "nested"
)

// Node is a single key of the diff tree. OldValue is set for removed,
// changed and unchanged keys, NewValue for added, changed and unchanged
// keys. Children is set only for nested keys.
type Node struct {
	Key      string
	Status   Status
	OldValue any
	NewValue any
	Children []Node
}

// Build compares data1 with data2 and returns the diff nodes sorted by key.
func Build(data1, data2 map[string]any) []Node {
	keys := slices.Collect(maps.Keys(data1))
	for key := range data2 {
		if _, ok := data1[key]; !ok {
			keys = append(keys, key)
		}
	}
	slices.Sort(keys)

	nodes := make([]Node, 0, len(keys))
	for _, key := range keys {
		nodes = append(nodes, buildNode(key, data1, data2))
	}

	return nodes
}

func buildNode(key string, data1, data2 map[string]any) Node {
	value1, in1 := data1[key]
	value2, in2 := data2[key]

	switch {
	case !in2:
		return Node{Key: key, Status: Removed, OldValue: value1}
	case !in1:
		return Node{Key: key, Status: Added, NewValue: value2}
	}

	map1, isMap1 := value1.(map[string]any)
	map2, isMap2 := value2.(map[string]any)

	switch {
	case isMap1 && isMap2:
		return Node{Key: key, Status: Nested, Children: Build(map1, map2)}
	case reflect.DeepEqual(value1, value2):
		return Node{Key: key, Status: Unchanged, OldValue: value1, NewValue: value2}
	default:
		return Node{Key: key, Status: Changed, OldValue: value1, NewValue: value2}
	}
}
