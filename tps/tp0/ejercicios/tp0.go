package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}

	posicion_mayor := 0
	for ind, valor := range vector {
		if valor > vector[posicion_mayor] {
			posicion_mayor = ind
		}
	}
	return posicion_mayor
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {

	var longitud int
	if len(vector1) < len(vector2) {
		longitud = len(vector1)
	} else {
		longitud = len(vector2)
	}
	for i := 0; i < longitud; i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	if len(vector1) < len(vector2) {
		return -1
	} else if len(vector1) > len(vector2) {
		return 1
	}
	return 0
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {

	ultimo := len(vector) - 1
	for ultimo > 0 {
		maximo_elem := Maximo(vector[:ultimo+1])
		Swap(&vector[maximo_elem], &vector[ultimo])
		ultimo = ultimo - 1
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[0] + Suma(vector[1:])

}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	if len(cadena) <= 1 {
		return true
	}
	primerCaracter := cadena[0]
	ultimoCaracter := cadena[len(cadena)-1]
	if primerCaracter == ultimoCaracter {
		return EsCadenaCapicua(cadena[1 : len(cadena)-1])
	}
	return false

}
