package model

import "encoding/json"

type Maybe[T any] struct {
	Nil   bool
	Value T
}

func NilValue[T any]() Maybe[T] {
	return Maybe[T]{
		Nil: true,
	}
}

func (i *Maybe[T]) UnmarshalJSON(bs []byte) error {
	if e := json.Unmarshal(bs, &i.Value); e != nil {
		i.Nil = true
		return e
	}
	i.Nil = false
	return nil
}

func Unwrap[T any](value Maybe[T]) T {
	return value.Value
}
