package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"code/diff"
)

const expectedJSON = `{
  "added": {
    "type": "added",
    "value": {
      "a": 12345,
      "b": {
        "c": null
      }
    }
  },
  "group": {
    "type": "nested",
    "children": {
      "empty": {
        "type": "changed",
        "oldValue": "",
        "newValue": "so much"
      },
      "list": {
        "type": "changed",
        "oldValue": [
          1
        ],
        "newValue": null
      },
      "same": {
        "type": "unchanged",
        "value": true
      }
    }
  },
  "removed": {
    "type": "removed",
    "value": 1.5
  }
}`

func TestJSON(t *testing.T) {
	got, err := JSON(tree)
	require.NoError(t, err)
	assert.JSONEq(t, expectedJSON, got)
}

func TestJSONEmpty(t *testing.T) {
	got, err := JSON(nil)
	require.NoError(t, err)
	assert.Equal(t, "{}", got)
}

func TestJSONNoHTMLEscape(t *testing.T) {
	got, err := JSON([]diff.Node{{Key: "a", Status: diff.Added, NewValue: "<b>&"}})
	require.NoError(t, err)
	assert.Contains(t, got, `"<b>&"`)
}

func TestJSONUnsupportedValue(t *testing.T) {
	_, err := JSON([]diff.Node{{Key: "a", Status: diff.Added, NewValue: make(chan int)}})
	require.Error(t, err)
}
