// Package code builds the difference between two configuration files.
package code

import (
	"code/diff"
	"code/formatters"
	"code/parsers"
)

// GenDiff returns the difference between two files in the given format.
// An empty format means the default "stylish" format.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := parsers.ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := parsers.ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	return formatters.Format(diff.Build(data1, data2), format)
}
