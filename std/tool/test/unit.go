package test

import (
	"iter"
	"kvs/pkg/tuple"
)

type Unit11[A1, E1 any] struct {
	Fn   func(A1) E1
	Args tuple.Ones[A1]
	Exps tuple.Ones[E1]
}

func NewUnit11[A1, E1 any](fn func(A1) E1) *Unit11[A1, E1] {
	return &Unit11[A1, E1]{
		Fn:   fn,
		Args: tuple.Ones[A1]{},
		Exps: tuple.Ones[E1]{},
	}
}

func (u *Unit11[A1, E1]) Case(a1 A1) func(E1) *Unit11[A1, E1] {
	u.Args.Push(a1)
	return func(e1 E1) *Unit11[A1, E1] {
		u.Exps.Push(e1)
		return u
	}
}

func (u *Unit11[A1, E1]) Run(fn func(tuple.One[E1], tuple.One[E1])) {
	for got, exp := range u.Results() {
		fn(got, exp)
	}
}

func (u *Unit11[A1, E1]) Results() iter.Seq2[tuple.One[E1], tuple.One[E1]] {
	return func(yield func(tuple.One[E1], tuple.One[E1]) bool) {
		argsNext, argsStop := iter.Pull(u.Args.Values())
		expsNext, expsStop := iter.Pull(u.Exps.Values())
		defer argsStop()
		defer expsStop()
		for {
			args, argsOk := argsNext()
			exps, expsOk := expsNext()
			if !argsOk || !expsOk {
				return
			}
			gots := tuple.NewOne(u.Fn(args.Unpack()))
			if !yield(gots, exps) {
				return
			}
		}
	}
}

// func NewA1E1[A1 any, E1 any](fn func(A1) E1) (func(a1 A1) func(e1 E1), func(t *testing.T)) {
// 	a, e := []func() A1{}, []func() E1{}
// 	add := func(a1 A1) func(e1 E1) {
// 		a = append(a, func() A1 { return a1 })
// 		return func(e1 E1) {
// 			e = append(e, func() E1 { return e1 })
// 		}
// 	}
// 	seq := func(yield func(E1, E1) bool) {
// 		for i := range min(len(a), len(e)) {
// 			if !yield(fn(a[i]()), e[i]()) {
// 				return
// 			}
// 		}
// 	}
// 	run := func(t *testing.T) {
// 		for got, expected := range seq {
// 			if got != expected {
// 				t.Errorf("got %v, expected %v", got, expected)
// 			}
// 		}
// 	}
// 	return add, run
// }

// func NewA2E1[A1, A2 any, E1 any](fn func(A1, A2) E1) (func(A1, A2) func(E1), func(t *testing.T)) {
// 	a, e := []func() (A1, A2){}, []func() E1{}
// 	add := func(a1 A1, a2 A2) func(e1 E1) {
// 		a = append(a, func() (A1, A2) { return a1, a2 })
// 		return func(e1 E1) {
// 			e = append(e, func() E1 { return e1 })
// 		}
// 	}
// 	seq := func(yield func(func() (E1, E1)) bool) {
// 		for i := range min(len(a), len(e)) {
// 			if !yield(func() (E1, E1) { return fn(a[i]()), e[i]() }) {
// 				return
// 			}
// 		}
// 	}
// 	run := func(t *testing.T) {
// 		for test := range seq {
// 			testE1(test())
// 		}
// 	}
// 	return add, run
// }
