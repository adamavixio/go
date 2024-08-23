package seq

import (
	"cmp"
	"iter"
)

func Range2[T cmp.Ordered](v, n, i T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		j := 0
		for v < n {
			if !yield(j, v) {
				return
			}
			j++
			v += i
		}
	}
}
