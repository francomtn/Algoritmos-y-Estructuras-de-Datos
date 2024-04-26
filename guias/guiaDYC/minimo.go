package guiaDYC

// BuscarMinimo devuelve el valor del minimo del arreglo, no su posicion
// Precondicion: el arreglo tiene al menos un elemento
func BuscarMinimo(arr []int) int {
	return minimo(arr, 0, len(arr)-1)
}

func minimo(arr []int, inicio, fin int) int {
	if inicio == fin {
		return arr[inicio]
	}
	medio := (inicio + fin) / 2
	aDerecha := minimo(arr, medio+1, fin)
	aIzquierda := minimo(arr, inicio, medio)

	if aDerecha < aIzquierda {
		return aDerecha
	} else {
		return aIzquierda
	}
}
