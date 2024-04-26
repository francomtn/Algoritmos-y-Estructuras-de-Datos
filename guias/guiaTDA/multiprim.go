package guiaTDA

func Multiprimeros[T any](cola Cola[T], k int) []T {
	if cola.EstaVacia() || k == 0 {
		return nil
	}

	temp := make([]T, 0)
	for !cola.EstaVacia() {
		temp = append(temp, cola.Desencolar())
	}
	if len(temp) < k {
		k = len(temp)
	}
	elem := make([]T, 0, k)

	for i := 0; i < k; i++ {
		elem = append(elem, temp[i])
	}

	for _, val := range temp {
		cola.Encolar(val)
	}

	return elem
}
