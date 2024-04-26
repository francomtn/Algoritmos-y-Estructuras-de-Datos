package guia0

func Merge(arr1, arr2 []int) []int {
	slice := make([]int, 0)
	return _Merge(arr1, arr2, slice)
}

func _Merge(arr1, arr2, arr3 []int) []int {
	if len(arr1) == 0 {
		return append(arr3, arr2...)
	}
	if len(arr2) == 0 {
		return append(arr3, arr1...)
	}
	if arr1[0] <= arr2[0] {
		arr3 = append(arr3, arr1[0])
		return _Merge(arr1[1:], arr2, arr3)
	} else {
		arr3 = append(arr3, arr2[0])
		return _Merge(arr1, arr2[1:], arr3)
	}
}
