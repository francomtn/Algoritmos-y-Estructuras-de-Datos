package heap

func EsHeap(arr []int) bool {
	if len(arr) >= 0 {
		return _esHeap(arr, 0)
	}
	return false
}

func _esHeap(arr []int, ind int) bool {

	if ind > len(arr) || len(arr) == 1 {
		return true
	}

	h_izq := 2*ind + 1
	h_der := 2*ind + 2
	if arr[ind] > arr[h_der] && arr[ind] > arr[h_izq] {
		return _esHeap(arr, ind+1)
	}
	return false
}
