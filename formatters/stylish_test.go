package formatters

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Lines are joined explicitly to keep the trailing space after "empty:".
var expectedStylish = strings.Join([]string{
	"{",
	"  + added: {",
	"        a: 12345",
	"        b: {",
	"            c: null",
	"        }",
	"    }",
	"    group: {",
	"      - empty: ",
	"      + empty: so much",
	"      - list: [1]",
	"      + list: null",
	"        same: true",
	"    }",
	"  - removed: 1.5",
	"}",
}, "\n")

func TestStylish(t *testing.T) {
	assert.Equal(t, expectedStylish, Stylish(tree))
}

func TestStylishEmpty(t *testing.T) {
	assert.Equal(t, "{\n}", Stylish(nil))
}
