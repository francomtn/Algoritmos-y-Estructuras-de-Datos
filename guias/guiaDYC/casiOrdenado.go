package guiaDYC

//Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado
//(todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar. Indicar y justificar el orden.

func ElementoDesordenado(arr []int) int {

	return obtenerElemento(arr, 0, len(arr)-1)
}

func obtenerElemento(arr []int, inicio, fin int) int {
	if inicio >= fin {
		return -1 // No se encontró ningún elemento fuera de lugar.
	}
	medio := (inicio + fin) / 2

	// Verificar si el elemento medio o sus vecinos están desordenados.
	if medio > inicio && medio < fin {
		if arr[medio-1] > arr[medio] && arr[medio] < arr[medio+1] {
			return arr[medio] // El elemento medio es el desordenado
		}
		if arr[medio-1] < arr[medio] && arr[medio] > arr[medio+1] {
			return arr[medio] // El elemento medio es el desordenado
		}
		if arr[medio-1] > arr[medio] && arr[medio] > arr[medio+1] {
			return arr[medio-1] // El elemento medio-1 es el desordenado
		}
	}

	// Explorar ambas mitades.
	leftResult := obtenerElemento(arr, inicio, medio-1)
	if leftResult != -1 {
		return leftResult
	}
	return obtenerElemento(arr, medio+1, fin)
}
