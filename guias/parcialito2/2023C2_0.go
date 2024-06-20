package parcialito2

import (
	"fmt"
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/lista"
)

/*
1. Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el que
para cada posición i de dicho arreglo, contenga el resultado de la máxima suma obtenible entre k elementos entre los
índices [0;i] del arreglo A (incluyendo a i). Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor
-1. Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 9, 12, 12, 17]. La
complejidad del algoritmo debe ser mejor que O(n · k). Indicar y justificar la complejidad del algoritmo implementado.*/

func SumarK(arr []int, k int) []int {

	nuevo := make([]int, len(arr))

	for i := 0; i < k-1; i++ {
		nuevo[i] = -1
	}
	for i := k - 1; i < len(arr); i++ {
		heap := cola_prioridad.CrearHeapArr(arr[0:i], func(a, b int) int { return a - b })
		suma := 0
		//fmt.Println(arr[0:i])
		fmt.Println(heap)
		for j := 0; j < k; j++ {
			if !heap.EstaVacia() {
				suma += heap.Desencolar()

			}
		}
		nuevo[i] = suma

	}

	return nuevo
}

/*
2. Implementar una función que reciba un hash de claves genéricas y como dato una lista cuyos datos sean genéricos, y
devuelva un nuevo hash cuyas claves sean las mismas del recibido por parámetro, pero el valor sea el dato del medio de
la lista correspondiente al dato original. Se sabe que cada lista tiene a lo sumo M elementos (un valor no constante). El
diccionario original (ni sus listas) deben verse modificados por esta función. Por ejemplo, si en el diccionario original
hay una clave boquita y como dato una lista [1977, 1978, 2000, 2001, 2003, 2007], en el nuevo diccionario debe
estar boquita como clave y como dato 2000. Indicar y justificar la complejidad del algoritmo implementado. La firma
de la función debe ser:
func DictMedio[K comparable, T any](dict Diccionario[K, Lista[T]]) Diccionario[K, T]*/

func DictMedio[K comparable, T any](dict diccionario.Diccionario[K, lista.Lista[T]]) diccionario.Diccionario[K, T] {

	dic := diccionario.CrearHash[K, T]()

	for iter := dict.Iterador(); iter.HaySiguiente(); iter.Siguiente() {

		clave, valor := iter.VerActual()
		if !dic.Pertenece(clave) {
			indice := (valor.Largo() / 2) - 1
			iterLista := valor.Iterador()
			for i := 0; i < indice; i++ {
				iterLista.Siguiente()
			}
			dato := iterLista.VerActual()
			dic.Guardar(clave, dato)
		}
	}
	return dic
}

/*
Crear un iterador para la lista y avanzar hasta el índice medio tiene una complejidad O(m/2), donde 
m es el número de elementos en la lista. Esto se puede simplificar a O(m).
Guardar el dato en el nuevo diccionario es O(1).
Dado que iteramos sobre cada clave del diccionario original una vez (O(n)) y para cada clave procesamos una lista de longitud 
m (O(m)), la complejidad total por clave es O(m).

Por lo tanto, la complejidad temporal total del algoritmo O(n⋅m), donde:
n es el número de claves en el diccionario original.
m es el tamaño máximo de las listas en el diccionario.
*/

/*
3. Implementar una primitiva del ABB que dado un valor entero M, una clave inicial inicio y una clave final fin, se
devuelva una lista con todos los datos cuyas claves estén entre inicio y fin, que estén dentro de los primeros M
niveles del árbol (considerando a la raíz en nivel 1). Indicar y justificar la complejidad temporal.
Por ejemplo, si tenemos el siguiente ABB con M = 3, inicio = 5 y fin = 15:*/


func (ab *diccionario.DiccionarioOrdenado[K, V])obtenerDatos(inicio, fin, m int) lista.Lista[K] {
	lista := lista.CrearListaEnlazada[K]()
	ab.obtenerDatosRec(ab.raiz, inicio, fin, m, 1, lista)
	return lista
}

func (ab *diccionario.DiccionarioOrdenado[K, V])obtenerDatosRec(nodo *nodoAbb[K,V],inicio, fin, m, nivel int, lista lista.Lista[K]) {

	if nodo == nil {return}
	if nivel > m { return}
	if ab.cmp(nodo.clave, inicio) >= 0 && ab.cmp(nodo.clave, fin) <= 0 {
		lista.InsertarUltimo(nodo.dato)		
	}
	if ab.cmp(nodo.clave, inicio) > 0 {
		ab.obtenerDatosRec(nodo.izq, inicio, fin, m, nivel+1, lista)
	}
	if ab.cmp(nodo.clave, inicio) < 0 {
		ab.obtenerDatosRec(nodo.der, inicio, fin, m, nivel+1, lista)
	}
}

