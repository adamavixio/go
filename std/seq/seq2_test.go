package seq

import (
	"slices"
	"testing"
)

func TestRange2(t *testing.T) {
	got := [][2]int{}
	exp := [][2]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}, {4, 4}}
	for i, v := range Range2(0, 5, 1) {
		got = append(got, [2]int{i, v})
	}
	if !slices.Equal(got, exp) {
		t.Errorf("got %v, exp %v", got, exp)
	}
}
