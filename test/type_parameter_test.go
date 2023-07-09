package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[P any](param P) P {
	fmt.Println(param)
	return param
}

func TestIndex(t *testing.T) {
	var resultText string = Length[string]("Viking")
	assert.Equal(t, "Viking", resultText)

	resultNumber := Length[int](300)
	assert.Equal(t, 300, resultNumber)
}
