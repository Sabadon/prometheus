package atomicutil

import (
	"sync/atomic"
)

type packedValue[T any] struct {
	val T
}

type GenericValue[T any] struct {
	v atomic.Value
}

func NewGenericValue[T any](initialValue T) *GenericValue[T] {
	g := &GenericValue[T]{}
	g.Store(initialValue)
	return g
}

func (g *GenericValue[T]) Store(value T) {
	g.v.Store(&packedValue[T]{val: value})
}

func (g *GenericValue[T]) Load() T {
	packed, _ := g.v.Load().(*packedValue[T])

	if packed == nil {
		var zero T
		return zero
	}

	return packed.val
}
