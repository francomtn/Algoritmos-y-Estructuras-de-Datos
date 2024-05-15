package heap

func Combinar(arr []int, k int) []int {

	heap := CrearHeapArr[int](arr, func(a, b int) int { return b - a })

	for heap.Cantidad() >= 2 && heap.VerMax() <= k {
		a := heap.Desencolar()
		b := heap.Desencolar()

		if b > k {
			return nil
		}
		heap.Encolar(a + 2*b)
	}
	if heap.VerMax() < k {
		return nil
	}

	res := make([]int, heap.Cantidad())

	for i := 0; i < len(res); i++ {
		res[i] = heap.Desencolar()
	}
	return res
}
