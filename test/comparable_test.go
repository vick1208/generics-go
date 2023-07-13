package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[C comparable](val1, val2 C) bool {
	return val1 == val2
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Viking", "Viking"))
	assert.True(t, IsSame[int](220, 220))
	assert.True(t, IsSame[int](-98, -98))
}
