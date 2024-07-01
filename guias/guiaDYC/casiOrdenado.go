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

	if ini+1 == fin {
		if arr[ini] > arr[fin] {
			return arr[ini]
		}
		return -1
	}

	medio := (ini + fin) / 2

	// Verificar si el elemento desordenado está en el punto medio
	if arr[medio] < arr[medio-1] {
		return arr[medio]
	} else if arr[medio] > arr[medio+1] {
		return arr[medio+1]
	}

	// Buscar en la mitad izquierda
	izq := obtenerElemento(arr, ini, medio)
	if izq != -1 {
		return izq
	}

	// Buscar en la mitad derecha
	return obtenerElemento(arr, medio+1, fin)
}

/*func ElementoDesordenado(arr []int) int {
	return desordenado(arr, 0, len(arr)-1)
}

func desordenado(arr []int, inicio, fin int) int {
	if inicio == fin {
		return -1
	}
	if fin == inicio+1 {
		if arr[inicio] > arr[fin] {
			return arr[inicio]
		} else {
			return -1
		}
	}
	mid := (inicio + fin) / 2
	if arr[mid] < arr[mid-1] {
		return arr[mid]
	} else if arr[mid] > arr[mid+1] {
		return arr[mid+1]
	}
	if arr[mid] >= arr[inicio] {
		return desordenado(arr, mid+1, fin)
	} else {
		return desordenado(arr, inicio, mid-1)
	}
}*/
