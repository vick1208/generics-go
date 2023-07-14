package test

import (
	"fmt"
	"testing"
)

type Bag[C any] []C

func PBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	words := Bag[string]{"Anything", "Anywhere", "Eureka"}
	PBag(words)
}
func TestBagInt(t *testing.T) {
	nums := Bag[int]{90, 23, 33}
	PBag[int](nums)
}
