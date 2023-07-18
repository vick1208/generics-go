package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func ExperMin[Y constraints.Ordered](first, second Y) Y {
	if first < second {
		return first
	} else {
		return second
	}
}
func TestExperMin(t *testing.T) {
	assert.Equal(t, 100, ExperMin[int](100, 200))
	assert.Equal(t, 100.0, ExperMin[float64](100.0, 100.8))
}

func TestExperMaps(t *testing.T) {
	first := map[string]string{
		"Name": "Eko",
	}
	second := map[string]string{
		"Name": "Eko",
	}

	assert.True(t, maps.Equal(first, second))
	// assert.True(t, maps.Equal[map[string]string](first, second)) => ekuivalen
}
func TestExperSlice(t *testing.T) {
	alpha := []string{"Bandar"}
	beta := []string{"Bandar"}
	assert.True(t, slices.Equal[string](alpha, beta))
	// assert.True(t, slices.Equal[string](alpha, beta)) => ekuivalen
}
