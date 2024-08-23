package seq

import (
	"fmt"
	"slices"
	"testing"
)

func TestRange(t *testing.T) {
	for _, test := range []struct {
		args     func() (int, int)
		expected []int
	}{
		{args: func() (int, int) { return 0, 0 }, expected: []int{}},
		{args: func() (int, int) { return 0, 2 }, expected: []int{0, 1}},
		{args: func() (int, int) { return 1, 6 }, expected: []int{1, 2, 3, 4, 5}},
		{args: func() (int, int) { return 1, -1 }, expected: []int{1, 0}},
		{args: func() (int, int) { return 5, 0 }, expected: []int{5, 4, 3, 2, 1}},
	} {
		s := Range(test.args())
		got := slices.Collect(s)
		fmt.Println("\ngot:", got)
		if !slices.Equal(got, test.expected) {
			t.Errorf("got %v, exp %v", got, test.expected)
		}
		got2 := slices.Collect(s)
		fmt.Println("\ngot2:", got2)
		if !slices.Equal(got2, test.expected) {
			t.Errorf("got %v, exp %v", got, test.expected)
		}
	}
}

// func TestStep(t *testing.T) {
// 	for _, test := range []struct {
// 		args     []int
// 		expected []int
// 	}{
// 		{args: []int{0, 0}, expected: []int{}},
// 		{args: []int{0, 2}, expected: []int{0, 1}},
// 		{args: []int{1, 6}, expected: []int{1, 2, 3, 4, 5}},
// 		{args: []int{1, -1}, expected: []int{1, 0}},
// 		{args: []int{5, 0}, expected: []int{5, 4, 3, 2, 1}},
// 	} {
// 		got := slices.Collect(Step(get()))
// 		if !slices.Equal(got, test.expected) {
// 			t.Errorf("got %v, exp %v", got, test.expected)
// 		}
// 	}
// }
