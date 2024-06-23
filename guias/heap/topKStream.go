package heap

import h "tdas/cola_prioridad"

func TopKStream(arr []int, k int) []int {

	heap := h.CrearHeapArr(arr[:k], func(a, b int) int { return b - a })

	for _, elem := range arr[k:] {

		if elem <= heap.VerMax() {
			continue
		}
		heap.Encolar(elem)
		heap.Desencolar()
	}
	res := make([]int, k)
	for i := range k {
		res[i] = heap.Desencolar()
	}
	return res
}

//O(k + n log k) como k << n -->  O( n log k)
