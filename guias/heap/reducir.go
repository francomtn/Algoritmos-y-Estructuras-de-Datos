package heap

import "tdas/cola_prioridad"

/*
2. Definimos a la reducción de una secuencia de números como la operación de sacar los dos elementos más pequeños de la
misma, y volver a guardar el resultado de 2 * mínimo + segundo_mínimo. Esto sólo puede hacerse si hay al menos
dos elementos. Ejemplo: [1, 7, 2, 3] -> [7, 3, 4]. De esta forma, la secuencia queda con un elemento menos.
Implementar un algoritmo que reciba un arreglo y devuelva el único valor que quedaría en el arreglo si aplicaramos
dicha reducción hasta que quede un único elemento en el arreglo. Indicar y justificar la complejidad del algoritmo.
*/

func Reducir(arr []int) int {

	heap := cola_prioridad.CrearHeapArr(arr, func(a, b int) int { return b - a })

	for heap.Cantidad() >= 2 {
		a := heap.Desencolar()
		b := heap.Desencolar()
		val := 2*a + b
		heap.Encolar(val)
	}
	return heap.VerMax()
}
