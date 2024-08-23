package minheap

import (
	"container/heap"
	"math/rand"
	"testing"

	"github.com/adamavixio/go/std/container/seq"
)

func TestMinHeap(t *testing.T) {
	h := &[]int{5, 4, 3, 2, 1}
	Init(h)
	for exp := range seq.Range(1, 5) {
		got := Pop(h)
		if got != exp {
			t.Errorf("got %v, expected %v", got, exp)
		}
	}
	for v := 1; v <= 5; v++ {
		Push(h, v)
	}
	for exp := range seq.Range(1, 5) {
		got := Pop(h)
		if got != exp {
			t.Errorf("got %v, expected %v", got, exp)
		}
	}
}

func BenchmarkMinHeap(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	h := &[]int{}
	for i := 0; i < b.N; i++ {
		if len(*h) == 0 || r.Intn(2)%2 == 0 {
			Push(h, i)
		} else {
			Pop(h)
		}
	}
}

type StdLibMinHeap []int

func (h StdLibMinHeap) Len() int {
	return len(h)
}

func (h StdLibMinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h StdLibMinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *StdLibMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *StdLibMinHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func BenchmarkHeapStd(b *testing.B) {
	r := rand.New(rand.NewSource(0))
	h := &StdLibMinHeap{}
	for i := 0; i < b.N; i++ {
		if len(*h) == 0 || r.Intn(2)%2 == 0 {
			heap.Push(h, i)
		} else {
			heap.Pop(h)
		}
	}
}
