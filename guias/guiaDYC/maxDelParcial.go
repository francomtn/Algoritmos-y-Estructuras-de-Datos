package guiaDYC

/*
3. Se tiene un arreglo de tamaño conocido n, inicialmente ordenado de menor a mayor y sin elementos repetidos. Al mismo
se le aplica un corrimiento de una cierta cantidad k de elementos. Esto es, todos los elementos que están a partir del
índice k se los correrá hacia la izquierda k posiciones. Los elementos entre las posiciones 0 y k − 1, estarán ahora al
final del arreglo (cada sub-segmento mantendrá el orden original de los elementos). Por ejemplo: v = [0, 1, 3, 5,
7, 8, 9] luego de correrlo con k = 3 resulta en v = [5, 7, 8, 9, 0, 1, 3].
Implementar en GO una función que devuelva el valor del máximo elemento de un arreglo ya corrido en una cantidad
k desconocida tal que 1 ≤ k ≤ n − 1, utilizando un algoritmo de orden O(log(n)).
La firma de la función es: func buscar_maximo(v int[], ini int, fin int) int, y será llamada inicialmente con:
buscar_maximo(v, 0, n-1). Justificar la complejidad del algoritmo propuesto.
*/

func BuscarMax(arr []int) int {

	return buscar_maximo(arr, 0, len(arr)-1)

}

func buscar_maximo(arr []int, ini, fin int) int {

	if ini == fin {
		return arr[ini]
	}

	medio := (ini + fin) / 2

	if arr[medio] > arr[medio-1] && arr[medio] > arr[medio+1] {
		return arr[medio]
	} else if arr[medio] > arr[medio+1] {
		return buscar_maximo(arr, medio, fin-1)
	} else {
		return buscar_maximo(arr, ini, medio)
	}
}

//[6, 7, 8, 9, 0, 1, 3, 4, 5].
