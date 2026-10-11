package formatters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"code/diff"
)

// JSON renders tree as an indented JSON array. Every node has "key" and
// "type"; changed nodes add "oldValue" and "newValue", nested nodes add
// "children", and all other nodes add "value".
func JSON(tree []diff.Node) (string, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "  ")

	if err := encoder.Encode(jsonNodes(tree)); err != nil {
		return "", fmt.Errorf("encode json: %w", err)
	}

	return strings.TrimSuffix(buf.String(), "\n"), nil
}

// jsonNode fixes the field order of the output. Values are pointers so that
// absent fields are omitted while a null value is still encoded as null.
type jsonNode struct {
	Key      string      `json:"key"`
	Type     diff.Status `json:"type"`
	Value    *any        `json:"value,omitempty"`
	OldValue *any        `json:"oldValue,omitempty"`
	NewValue *any        `json:"newValue,omitempty"`
	Children []jsonNode  `json:"children,omitempty"`
}

func jsonNodes(nodes []diff.Node) []jsonNode {
	result := make([]jsonNode, 0, len(nodes))
	for _, node := range nodes {
		item := jsonNode{Key: node.Key, Type: node.Status}

		switch node.Status {
		case diff.Nested:
			item.Children = jsonNodes(node.Children)
		case diff.Changed:
			item.OldValue = &node.OldValue
			item.NewValue = &node.NewValue
		case diff.Added:
			item.Value = &node.NewValue
		case diff.Removed, diff.Unchanged:
			item.Value = &node.OldValue
		}

		result = append(result, item)
	}

	return result
}
