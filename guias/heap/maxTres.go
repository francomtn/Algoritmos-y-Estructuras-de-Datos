package heap

type fcmpHeap[T any] func(T, T) int

/*type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      fcmpHeap[T]
}

func (h heap[T]) Max3() []T {

	n := 3
	if h.cantidad < 3 {
		n = h.cantidad
	}
	res := make([]T, n)
	res = append(res, h.datos[0])

	h._max3(res, 0, n)

	return res
}
func (h *heap[T]) _max3(res []T, ind int, n int) {

	if len(res) == n {
		return
	}

	h_izq := ind*2 + 1
	h_der := ind*2 + 2

	if h.cmp(h.datos[h_izq], h.datos[h_der]) > 0 {
		res = append(res, h.datos[h_izq])
		h._max3(res, h_izq, n)
	} else if h.cmp(h.datos[h_izq], h.datos[h_der]) < 0 {
		res = append(res, h.datos[h_der])
		h._max3(res, h_der, n)
	}

}
*/
