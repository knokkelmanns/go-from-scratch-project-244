// Package formatters renders a diff tree in one of the output formats.
package formatters

import (
	"fmt"

	"code/diff"
)

// DefaultFormat is used when no format is given.
const DefaultFormat = "stylish"

// Format renders tree in the given format. An empty format means DefaultFormat.
func Format(tree []diff.Node, format string) (string, error) {
	if format == "" {
		format = DefaultFormat
	}

	switch format {
	case "stylish":
		return Stylish(tree), nil
	case "plain":
		return Plain(tree), nil
	case "json":
		return JSON(tree)
	default:
		return "", fmt.Errorf("unknown format: %q", format)
	}
}
