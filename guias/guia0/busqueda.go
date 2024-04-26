package guia0

func BuscarElem(arr []int, elem int) int {
	for ind, valor := range arr {
		if valor == elem {
			return ind
		}
	}
	return -1
}
