package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	Alpha T
	Beta  T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeAlpha(first T) T {
	d.Alpha = first
	return d.Alpha
}

func TestData(t *testing.T) {
	data := Data[string]{
		Alpha: "Edhi",
		Beta:  "Kurniawan",
	}
	fmt.Print(data)
}

func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		Alpha: "Edhi",
		Beta:  "Kurniawan",
	}
	assert.Equal(t, "Bobi", data.ChangeAlpha("Bobi"))
	assert.Equal(t, "Hello Edhi", data.SayHello("Edhi"))
}
