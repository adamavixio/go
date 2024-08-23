package tuple

import (
	"iter"
	"slices"
)

type One[V1 any] struct {
	V1 V1
}

func NewOne[V1 any](v1 V1) One[V1] {
	return One[V1]{V1: v1}
}

func (t *One[V1]) Unpack() V1 {
	return t.V1
}

type Ones[V1 any] []One[V1]
type OnesChain[V1 any] func(V1) OnesChain[V1]

func (t *Ones[V1]) Push(v1 V1) OnesChain[V1] {
	*t = append(*t, NewOne(v1))

	var chain OnesChain[V1]
	chain = func(v1 V1) OnesChain[V1] {
		*t = append(*t, NewOne(v1))
		return chain
	}

	return chain
}

func (t Ones[V1]) Values() iter.Seq[One[V1]] {
	return slices.Values(t)
}

func (t Ones[V1]) All() iter.Seq2[int, One[V1]] {
	return slices.All(t)
}

type Two[V1, V2 any] struct {
	V1 V1
	V2 V2
}

func NewTwo[V1, V2 any](v1 V1, v2 V2) Two[V1, V2] {
	return Two[V1, V2]{V1: v1, V2: v2}
}

func (t *Two[V1, V2]) Unpack() (V1, V2) {
	return t.V1, t.V2
}

type Twos[V1, V2 any] []Two[V1, V2]
type TwosChain[V1, V2 any] func(V1, V2) TwosChain[V1, V2]

func (t *Twos[V1, V2]) Push(v1 V1, v2 V2) TwosChain[V1, V2] {
	*t = append(*t, NewTwo(v1, v2))

	var chain TwosChain[V1, V2]
	chain = func(v1 V1, v2 V2) TwosChain[V1, V2] {
		*t = append(*t, NewTwo(v1, v2))
		return chain
	}

	return chain
}

func (t Twos[V1, V2]) Values() iter.Seq[Two[V1, V2]] {
	return slices.Values(t)
}

func (t Twos[V1, V2]) All() iter.Seq2[int, Two[V1, V2]] {
	return slices.All(t)
}
