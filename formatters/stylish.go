package formatters

import (
	"fmt"
	"maps"
	"slices"
	"strings"

	"code/diff"
)

const (
	indentSize = 4
	signWidth  = 2
)

// Stylish renders tree as a JSON-like structure where changed lines are
// marked with "+" and "-".
func Stylish(tree []diff.Node) string {
	var b strings.Builder
	b.WriteString("{\n")
	writeNodes(&b, tree, 1)
	b.WriteString("}")

	return b.String()
}

func writeNodes(b *strings.Builder, nodes []diff.Node, depth int) {
	for _, node := range nodes {
		switch node.Status {
		case diff.Nested:
			writeLine(b, depth, ' ', node.Key, "{")
			writeNodes(b, node.Children, depth+1)
			writeClosingBrace(b, depth)
			b.WriteByte('\n')
		case diff.Added:
			writeLine(b, depth, '+', node.Key, stringify(node.NewValue, depth))
		case diff.Removed:
			writeLine(b, depth, '-', node.Key, stringify(node.OldValue, depth))
		case diff.Changed:
			writeLine(b, depth, '-', node.Key, stringify(node.OldValue, depth))
			writeLine(b, depth, '+', node.Key, stringify(node.NewValue, depth))
		case diff.Unchanged:
			writeLine(b, depth, ' ', node.Key, stringify(node.OldValue, depth))
		}
	}
}

// writeLine writes "<indent><sign> <key>: <value>", where the sign is shifted
// left of the key's indentation.
func writeLine(b *strings.Builder, depth int, sign rune, key, value string) {
	indent := strings.Repeat(" ", depth*indentSize-signWidth)
	fmt.Fprintf(b, "%s%c %s: %s\n", indent, sign, key, value)
}

func writeClosingBrace(b *strings.Builder, depth int) {
	b.WriteString(strings.Repeat(" ", depth*indentSize))
	b.WriteByte('}')
}

// stringify renders a plain value. Maps are rendered as nested blocks whose
// keys sit one level deeper than depth.
func stringify(value any, depth int) string {
	switch v := value.(type) {
	case nil:
		return "null"
	case map[string]any:
		var b strings.Builder
		b.WriteString("{\n")
		for _, key := range slices.Sorted(maps.Keys(v)) {
			writeLine(&b, depth+1, ' ', key, stringify(v[key], depth+1))
		}
		writeClosingBrace(&b, depth)

		return b.String()
	default:
		return fmt.Sprint(v)
	}
}
