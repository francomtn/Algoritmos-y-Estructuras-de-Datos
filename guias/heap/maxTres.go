package heap

import "strings"

type fcmpHeap[T any] func(T, T) int

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func (h heap[T]) Max3() []T {

	heap := CrearHeapArr[T](h.cmp(a, b T) int { return a - b})
	res := make([]T, 3)
	
	
	return res
}
