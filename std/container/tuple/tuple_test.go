package tuple

import (
	"slices"
	"testing"
)

func TestOne(t *testing.T) {
	tuples := Ones[int]{}
	tuples.Push(1)(2)(3)(4)(5)

	got1 := []int{1, 2, 3, 4, 5}
	exp1 := []int{}

	for tuple := range tuples.Values() {
		v1 := tuple.Unpack()
		exp1 = append(exp1, v1)
	}

	if !slices.Equal(got1, exp1) {
		t.Errorf("got %v, expected %v", got1, exp1)
	}
}

func TestTwo(t *testing.T) {
	tuples := Twos[int, string]{}
	tuples.Push(1, "1")(2, "2")(3, "3")(4, "4")(5, "5")

	got1 := []int{1, 2, 3, 4, 5}
	got2 := []string{"1", "2", "3", "4", "5"}
	exp1 := []int{}
	exp2 := []string{}

	for tuple := range tuples.Values() {
		v1, v2 := tuple.Unpack()
		exp1, exp2 = append(exp1, v1), append(exp2, v2)
	}

	if !slices.Equal(got1, exp1) {
		t.Errorf("got %v, expected %v", got1, exp1)
	}

	if !slices.Equal(got2, exp2) {
		t.Errorf("got %v, expected %v", got2, exp2)
	}
}
