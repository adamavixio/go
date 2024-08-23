package minheap

import (
	"cmp"

	"github.com/adamavixio/go/std/heap/slice"
)

func Init[H ~[]E, E cmp.Ordered](h *H) {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		slice.Down(*h, Less, n, i)
	}
}

func Push[H ~[]E, E cmp.Ordered](h *H, v E) {
	*h = append(*h, v)
	slice.Up(*h, Less, len(*h)-1)
}

func Pop[H ~[]E, E cmp.Ordered](h *H) E {
	n := len(*h) - 1
	v := (*h)[0]
	(*h)[0] = (*h)[n]
	*h = (*h)[:n]
	slice.Down(*h, Less, n, 0)
	return v
}

func Less[H ~[]E, E cmp.Ordered](h H, i, j int) bool {
	return h[i] < h[j]
}
