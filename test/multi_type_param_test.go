package test

import (
	"fmt"
	"testing"
)

func MultipleParam[T1 any, T2 any](alpha T1, beta T2) {
	fmt.Println(alpha)
	fmt.Println(beta)
}

func TestMultipleParameter(t *testing.T) {
	MultipleParam[string, bool]("Kon Satoshi", false)
	MultipleParam[bool, int](12 != 0, 201)

}
