package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func FindTheMin[A interface{ int | int64 | float64 }](first, second A) A {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestFindTheMin(t *testing.T) {
	assert.Equal(t, int64(200), FindTheMin[int64](200, 600))
	assert.Equal(t, 200, FindTheMin[int](200, 600))
	assert.Equal(t, 200.0, FindTheMin[float64](200.0, 600.0))
}

func GetFirst[W []E, E any](slice W) E {
	first := slice[0]
	return first
}

func TestGetFirst(t *testing.T) {
	names := []string{
		"Rian",
		"Santoso",
		"Wibowo",
	}
	first := GetFirst[[]string, string](names)

	assert.Equal(t, "Rian", first)
}
