package guiaDYC

func Alternar(arr []int) []int {

	_alternar(arr, 1, len(arr)-1)
	return arr
}

func _alternar(arr []int, ini, fin int) {

	if len(arr) < 4 {
		return
	}
	medio := (ini + fin) / 2
	_alternar(arr, ini, medio)
	_alternar(arr, medio, fin)
	swap(&arr[ini], &arr[medio])
	arr[ini], arr[medio] = arr[medio], arr[ini]

}

func swap(a *int, b *int) {
	*a, *b = *b, *a

}
