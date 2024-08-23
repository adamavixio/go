package seq

import (
	"iter"
)

func Range[N Number](s, e N) iter.Seq[N] {
	if s < e {
		return inc(s, e, 1)
	}
	return dec(s, e, 1)
}

func RangeS[N Number](s, e, i N) iter.Seq[N] {
	if i < 1 {
		panic("step size cannot be less than 1")
	}
	if s < e {
		return inc(s, e, i)
	}
	return dec(s, e, i)
}

func inc[N Number](s, e, i N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for v := s; v < e; v += i {
			if !yield(v) {
				return
			}
		}
	}
}

func dec[N Number](s, e, i N) iter.Seq[N] {
	return func(yield func(N) bool) {
		for v := s; v < e; v += i {
			if !yield(v) {
				return
			}
		}
	}
}
