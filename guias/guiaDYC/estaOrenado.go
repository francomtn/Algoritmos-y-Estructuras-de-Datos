package guiaDYC

func EstaOrdenado(arr []int) bool {

	if len(arr) == 0 {
		return true
	}
	return _EstaOrdenado(arr, 0, len(arr)-1)
}

func _EstaOrdenado(arr []int, inicio, fin int) bool {
	if inicio == fin {
		return true
	}
	medio := (inicio + fin) / 2

	if arr[medio] > arr[medio+1] {
		return false
	}

	aDerecha := _EstaOrdenado(arr, medio+1, fin)
	aIzquierda := _EstaOrdenado(arr, inicio, medio)

	return aDerecha && aIzquierda
}

// la ec de recurrencia es T(n)= 2T(n/2) +O(1) -> A=2,B=2,C=0
//log2(2)=1 > 0 -> O(n^(log2(2))) -> O(n¹)
// por lo tanto el orden es O(n)
