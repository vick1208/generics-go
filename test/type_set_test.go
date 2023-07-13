package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type Number interface {
	int | int8 | int16 | int32 |
		int64 | float32 | float64
}

func Min[T Number](first, second T) T {
	if first < second {
		return first
	} else {
		return second
	}
}

func TestMin(t *testing.T) {
	assert.Equal(t, 100, Min[int](100, 300))
	assert.Equal(t, int64(100), Min[int64](int64(100), int64(300)))
	assert.Equal(t, float32(100.76), Min[float32](100.76, 222.6))
	assert.Equal(t, float64(100.76), Min[float64](100.76, 222.6))
}
