package formatters

import (
	"fmt"
	"strings"

	"code/diff"
)

// Plain renders tree as a list of human-readable sentences, one per changed
// property. Unchanged properties are skipped.
func Plain(tree []diff.Node) string {
	return strings.Join(plainLines(tree, ""), "\n")
}

func plainLines(nodes []diff.Node, parent string) []string {
	var lines []string
	for _, node := range nodes {
		path := node.Key
		if parent != "" {
			path = parent + "." + node.Key
		}

		switch node.Status {
		case diff.Nested:
			lines = append(lines, plainLines(node.Children, path)...)
		case diff.Added:
			lines = append(lines, fmt.Sprintf("Property '%s' was added with value: %s",
				path, plainValue(node.NewValue)))
		case diff.Removed:
			lines = append(lines, fmt.Sprintf("Property '%s' was removed", path))
		case diff.Changed:
			lines = append(lines, fmt.Sprintf("Property '%s' was updated. From %s to %s",
				path, plainValue(node.OldValue), plainValue(node.NewValue)))
		case diff.Unchanged:
		}
	}

	return lines
}

func plainValue(value any) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case string:
		return "'" + v + "'"
	case map[string]any, []any:
		return "[complex value]"
	default:
		return fmt.Sprint(v)
	}
}
