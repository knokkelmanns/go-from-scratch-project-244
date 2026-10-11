package formatters

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const expectedPlain = `Property 'added' was added with value: [complex value]
Property 'group.empty' was updated. From '' to 'so much'
Property 'group.list' was updated. From [complex value] to null
Property 'removed' was removed`

func TestPlain(t *testing.T) {
	assert.Equal(t, expectedPlain, Plain(tree))
}

func TestPlainEmpty(t *testing.T) {
	assert.Empty(t, Plain(nil))
}
