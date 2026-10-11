package formatters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"

	"code/diff"
)

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

type jsonNode struct {
	Type     diff.Status         `json:"type"`
	Value    *any                `json:"value,omitempty"`
	OldValue *any                `json:"oldValue,omitempty"`
	NewValue *any                `json:"newValue,omitempty"`
	Children map[string]jsonNode `json:"children,omitempty"`
}

func jsonNodes(nodes []diff.Node) map[string]jsonNode {
	result := make(map[string]jsonNode, len(nodes))
	for _, node := range nodes {
		item := jsonNode{Type: node.Status}

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

		result[node.Key] = item
	}

	return result
}
