package guiaTDA

func MergePilas(pila1, pila2 Pila[int]) []int {

	arr := make([]int, 0)

	for !pila1.EstaVacia() && !pila2.EstaVacia() {
		tope1 := pila1.VerTope()
		tope2 := pila2.VerTope()
		if tope1 > tope2 {
			arr = append(arr, pila1.Desapilar())
		} else {
			arr = append(arr, pila2.Desapilar())
		}
	}
	for !pila1.EstaVacia() {
		arr = append(arr, pila1.Desapilar())
	}
	for !pila2.EstaVacia() {
		arr = append(arr, pila2.Desapilar())
	}
	aux := make([]int, 0, len(arr))
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] != arr[i-1] {
			aux = append(aux, arr[i])
		} else {
			continue // Restar 1 adicional para saltar la repetición
		}
	}
	return aux
}
