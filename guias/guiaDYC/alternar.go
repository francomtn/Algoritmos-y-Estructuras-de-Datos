package guiaDYC

func Alternar(arr []int) []int {
	if len(arr) == 0 || len(arr)%2 != 0 {
		return arr
	}
	_alternar(arr, 0, len(arr)-1)
	return arr
}

// Función recursiva para alternar los elementos del arreglo en un rango
func _alternar(arr []int, ini, fin int) {
	if fin-ini <= 1 {
		return
	}
	medio := (ini + fin) / 2
	cuarto := (medio - ini + 1) / 2

	// Intercambiar los elementos correspondientes
	for i := 0; i < cuarto; i++ {
		arr[medio-cuarto+1+i], arr[medio+1+i] = arr[medio+1+i], arr[medio-cuarto+1+i]
	}

	// Llamada recursiva para las dos mitades
	_alternar(arr, ini, medio)
	_alternar(arr, medio+1, fin)
}
