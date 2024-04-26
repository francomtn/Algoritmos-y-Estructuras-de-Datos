package guia0

func Maximo[T any](arr []T, comparacion func(T, T) int) T {

	max := arr[0]
	for _, val := range arr[1:] {
		if comparacion(max, val) < 0 {
			max = val
		}
	}
	return max
}
