package database

import "time"

type Transactions[T Read | Write] struct {
	play   []Transaction[T]
	rewind []Transaction[T]
}

func (txs *Transactions[T]) Peek() {
	return txs.commit
}

type Transaction[T Read | Write] struct {
	Type      T
	Timestamp time.Time
}

func NewTransaction[T Read | Write](tx T) *Transaction[T] {
	return &Transaction[T]{
		Type:      tx,
		Timestamp: time.Now().UTC(),
	}
}
