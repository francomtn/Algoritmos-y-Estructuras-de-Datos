package heap

import (
	"tdas/cola_prioridad"
)

type compuesto struct {
	vector   int
	posicion int
	valor    int
}

func KMerge(arr [][]int) []int {

	heap := cola_prioridad.CrearHeap[compuesto](func(a, b compuesto) int { return b.valor - a.valor })
	res := make([]int, 0)

	for i, vector := range arr {
		if len(arr[i]) > 0 {
			heap.Encolar(compuesto{vector: i, posicion: 0, valor: vector[0]})
		}
	}
	for !heap.EstaVacia() {
		celda := heap.Desencolar()
		res = append(res, celda.valor)
		indice := celda.posicion + 1
		vector := celda.vector
		if indice < len(arr[vector]) {
			heap.Encolar(compuesto{vector: celda.vector, posicion: indice, valor: arr[celda.vector][indice]})

		}
	}
	return res
}
