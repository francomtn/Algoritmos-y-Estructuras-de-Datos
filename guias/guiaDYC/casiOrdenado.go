package guiaDYC

//Implementar, por división y conquista, una función que dado un arreglo sin elementos repetidos y casi ordenado
//(todos los elementos se encuentran ordenados, salvo uno), obtenga el elemento fuera de lugar. Indicar y justificar el orden.

func ElementoDesordenado(arr []int) int {
	if arr[0] > arr[1] {
		return arr[0]
	}
	return obtenerElemento(arr, 0, len(arr)-1)
}

func obtenerElemento(arr []int, ini, fin int) int {
	if ini == fin {
		return -1
	}

	medio := (ini + fin) / 2

	if (medio+1) < len(arr) && (medio-1) >= 0 {
		if arr[medio] > arr[medio+1] && arr[medio+1] > arr[medio-1] {
			return arr[medio]
		}
		if arr[medio] < arr[medio-1] && arr[medio-1] < arr[medio+1] {
			return arr[medio]
		}
	}

	izq := obtenerElemento(arr, ini, medio-1)
	der := obtenerElemento(arr, medio+1, fin)

	if izq == -1 {
		return der
	}

	return izq
}
