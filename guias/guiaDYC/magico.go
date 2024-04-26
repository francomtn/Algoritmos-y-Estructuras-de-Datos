package guiaDYC

func ArregloEsMagico(arr []int) bool {

	return EsMagico(arr, 0, len(arr)-1)
}

func EsMagico(arr []int, inicio, fin int) bool {
	if inicio > fin {
		return false
	}

	medio := (inicio + fin) / 2
	if arr[medio] == medio {
		return true
	} else if arr[medio] < medio {
		return EsMagico(arr, medio+1, fin)
	}
	return EsMagico(arr, inicio, medio-1)
}
