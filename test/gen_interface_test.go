package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[R any] interface {
	GetValue() R
	SetValue(value R)
}

func ChangeValue[S any](parameter GetterSetter[S], value S) S {
	parameter.SetValue(value)
	return parameter.GetValue()
}

type ThisData[T any] struct {
	Value T
}

func (d *ThisData[T]) GetValue() T {
	return d.Value
}
func (d *ThisData[T]) SetValue(value T) {
	d.Value = value
}

func TestGenInterface(t *testing.T) {
	thisData := ThisData[string]{}

	result := ChangeValue[string](&thisData, "Edhi")

	assert.Equal(t, "Edhi", result)
	assert.Equal(t, "Edhi", thisData.Value)
}
