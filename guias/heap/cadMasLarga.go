package heap

import (
	"tdas/cola_prioridad"
)

func CadenasLargas(arr []string, k int) []string {

	if k > len(arr) {
		k = len(arr)
	}

	heap := cola_prioridad.CrearHeapArr[string](arr[:k], func(a, b string) int { return len(b) - len(a) })
	res := make([]string, k)
	for _, elem := range arr[k:] {

		if len(elem) <= len(heap.VerMax()) {
			continue
		}
		heap.Encolar(elem)
		heap.Desencolar()
	}
	for i := k - 1; i >= 0; i-- {
		res[i] = heap.Desencolar()
	}

	return res
}
