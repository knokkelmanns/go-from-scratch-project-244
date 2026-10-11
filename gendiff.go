package code

import (
	"fmt"
	"reflect"
	"slices"
	"strings"
)

// GenDiff returns the difference between two files in the given format.
func GenDiff(filepath1, filepath2, format string) (string, error) {
	data1, err := ParseFile(filepath1)
	if err != nil {
		return "", err
	}

	data2, err := ParseFile(filepath2)
	if err != nil {
		return "", err
	}

	keys := make([]string, 0, len(data1)+len(data2))
	for key := range data1 {
		keys = append(keys, key)
	}
	for key := range data2 {
		if _, ok := data1[key]; !ok {
			keys = append(keys, key)
		}
	}
	slices.Sort(keys)

	var b strings.Builder
	b.WriteString("{\n")
	for _, key := range keys {
		value1, in1 := data1[key]
		value2, in2 := data2[key]

		switch {
		case in1 && in2 && reflect.DeepEqual(value1, value2):
			fmt.Fprintf(&b, "    %s: %s\n", key, stringify(value1))
		default:
			if in1 {
				fmt.Fprintf(&b, "  - %s: %s\n", key, stringify(value1))
			}
			if in2 {
				fmt.Fprintf(&b, "  + %s: %s\n", key, stringify(value2))
			}
		}
	}
	b.WriteString("}")

	return b.String(), nil
}

func stringify(value any) string {
	if value == nil {
		return "null"
	}
	return fmt.Sprint(value)
}
