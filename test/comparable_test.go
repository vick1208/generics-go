package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[C comparable](val1, val2 C) bool {
	if val1 == val2 {
		return true
	} else {
		return false
	}
}
func TestIsSame(t *testing.T) {
	assert.True(t, IsSame[string]("Viking", "Viking"))
	assert.True(t, IsSame[int](220, 220))
	assert.True(t, IsSame[int](-98, -98))
}
