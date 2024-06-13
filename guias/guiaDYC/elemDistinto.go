package guiaDYC

/*
Implementar un algoritmo que reciba un arreglo de n enteros (con n ≥ 3) en el que todos sus elementos son iguales
salvo 1, y determine (utilizando división y conquista) cual es dicho elemento no repetido. Indicar y justificar la
complejidad del algoritmo implementado.
*/

func EncontarDistinto(arr []int) int {

	return _encontrarElem(arr, 0, len(arr)-1)
}

func _encontrarElem(arr []int, ini, fin int) int {

	if ini == fin {

		return arr[ini]
	}

	if len(arr) <= 3 {
		for _, elem := range arr {
			if elem != arr[0] {
				return elem
			}
		}
	}
	medio := (ini + fin) / 2

	if arr[medio] != arr[medio+1] && arr[medio] != arr[medio-1] {
		return arr[medio]
	} else {
		izq := _encontrarElem(arr, ini, medio)
		der := _encontrarElem(arr, medio+1, fin)
		if izq != der && izq == arr[medio] {
			return der
		}
		return izq
	}
}

/*
func EncontarDistinto(arr []int) int {
	var repetido int
	if arr[0] == arr[1] {
		repetido = arr[0]
	} else if arr[1] == arr[2] {
		repetido = arr[1]
	} else if arr[0] == arr[2] {
		repetido = arr[0]
	}
	return _encontar(arr, 0, len(arr)-1, repetido)
}

func _encontar(arr []int, inicio, fin, repetido int) int {
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2

	izq := _encontar(arr, inicio, medio-1, repetido)
	der := _encontar(arr, medio+1, fin, repetido)

	if arr[medio] != repetido {
		return arr[medio]
	}
	if izq == -1 {
		return der
	}
	return izq
}
*/
