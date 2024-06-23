package heap

import h "tdas/cola_prioridad"

func TopK(arr []int, k int) []int {

	heap := h.CrearHeapArr(arr, func(a, b int) int { return a - b })
	var res []int
	for i := range k {
		res[i] = heap.Desencolar()
	}
	return res
}

//O( n + k log n)
