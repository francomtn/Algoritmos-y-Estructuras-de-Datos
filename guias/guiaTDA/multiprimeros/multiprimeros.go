package multiprimeros

func (cola *colaEnlazada[T]) Multiprimeros(k int) []T {

	if cola.primero == nil || k == 0 {
		return nil
	}
	elem := make([]T, 0, k)
	actual := cola.primero
	for i := 0; i < k && actual != nil; i++ {
		elem = append(elem, actual.dato)
		actual = actual.siguiente
	}
	return elem
}
