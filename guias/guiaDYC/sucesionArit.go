package guiaDYC

/*
Dado un arreglo de n enteros positivos que casi representa una progresión aritmética creciente (es una progresión
aritmética a la que le falta un elemento), implementar un algoritmo que devuelva el elemento faltante de manera eficiente
(complejidad logarítmica). Se puede suponer que el arreglo tiene al menos 4 elementos. Justificar la complejidad del
algoritmo implementado. Por ejemplo, si la sucesión es [5, 8, 14, 17, 20, 23] tiene que devolver 11.
*/

func SucesionAritmetica(arr []int) int {

	var dif int
	if arr[1]-arr[0] == arr[2]-arr[1] {
		dif = arr[1] - arr[0]
	} else if arr[1]-arr[0] == arr[3]-arr[2] {
		dif = arr[1] - arr[0]
	} else {
		dif = arr[2] - arr[1]
	}
	return encontrarFaltante(arr, 0, len(arr), dif)
}

func encontrarFaltante(arr []int, ini, fin, dif int) int {

	if ini >= fin-1 {
		return arr[ini] + dif
	}
	medio := (ini + fin) / 2

	// Verificar si el faltante está a la izquierda o a la derecha
	if arr[medio] != arr[0]+medio*dif {

		return encontrarFaltante(arr, ini, medio, dif) // El elemento faltante está en la primera mitad
	} else {
		return encontrarFaltante(arr, medio, fin, dif) // El elemento faltante está en la segunda mitad
	}
}
