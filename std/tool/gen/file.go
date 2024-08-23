package gen

import (
	"fmt"
	"iter"
	"kvs/pkg/seq"
)

type File struct {
	Package   string
	Functions []Function
}

// func generic(label string, types []Type) iter.Seq[iter.Seq2[string, Type]] {
// 	slices.Sort(types)
// 	t := slices.Compact(types)
// 	l := tag(label, len(t))
// 	return func(yield func(iter.Seq2[string, Type]) bool) {
// 		yield(func(yield func(string, Type) bool) {
// 			for ic := range indexCombinations(len(t)) {
// 				yield()
// 			}
// 		})
// 	}
// }

func tags(s string, n int) iter.Seq[string] {
	return func(yield func(string) bool) {
		for i := 1; i <= n; i++ {
			yield(fmt.Sprintf("%v%v", s, i))
		}
	}
}

func combinations(n int) iter.Seq[[]int] {
	return func(yield func([]int) bool) {
		c := make([]int, n)
		var recursive func(i int)
		recursive = func(i int) {
			if i == n && yield(c) {
				return
			}
			for j := range n {
				c[i] = j
				recursive(i + 1)
			}
		}
		recursive(0)
	}
}

func values[S ~[]E, E any](s S, o []int) iter.Seq[E] {
	return func(yield func(E) bool) {
		n := len(s)
		for _, i := range o {
			if i < 0 || i >= n {
				return
			}
			yield(s[i])
		}
	}
}

func zip[E1, E2 any](i1 iter.Seq[E1], i2 iter.Seq[E2]) iter.Seq2[E1, E2] {
	return func(yield func(E1, E2) bool) {
		n1, s1 := iter.Pull(i1)
		n2, s2 := iter.Pull(i2)
		defer s1()
		defer s2()
		for {
			v1, ok1 := n1()
			v2, ok2 := n2()
			if !ok1 || !ok2 {
				return
			}
			if !yield(v1, v2) {
				return
			}
		}
	}
}

func repeat[E any](s iter.Seq[E]) {

}

func main() {
	// s1 := slices.Values([]int{1, 2, 3})
	// s2 := slices.Values([]string{"a", "b", "c", "d"})
	// for e1, e2 := range zip(s1, s2) {
	// 	fmt.Println(e1, e2)
	// }

	s := seq.Range(0, 5)
	fmt.Println("first")
	s(func(v int) bool {
		fmt.Println(v)
		return true
	})
	fmt.Println("second")
	s(func(v int) bool {
		fmt.Println(v)
		return true
	})
}

// 1 2 3
// 11 12 13 21 22 23 31 32 33
// 111 121 131 121 122 123
