package test

import (
	"kvs/pkg/tuple"
	"testing"
)

func TestUnit11(t *testing.T) {
	unit := NewUnit11(func(x int) int {
		return x * x
	})
	unit.Case(1)(1).Case(2)(4).Case(3)(9)
	unit.Run(func(got, exp tuple.One[int]) {
		if got != exp {
			t.Errorf("got %v, exp %v", got, exp)
		}
	})
}
