package maybe_test

import (
	"testing"

	"github.com/crufter/functional/data/maybe"
	"github.com/stretchr/testify/assert"
)

func TestMaybe(t *testing.T) {
	var n interface{}
	n = maybe.NewNothing[string]()
	v, ok := n.(maybe.Maybe[string])
	assert.Equal(t, true, ok)
	assert.Equal(t, false, v.IsJust())
}
