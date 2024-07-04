package finales

import (
	"tdas/cola_prioridad"
	"tdas/diccionario"
	"tdas/lista"
)

/*
1. Implementar un algoritmo que reciba un arreglo de n enteros (con n ≥ 3) en el que todos sus elementos son iguales
salvo 1, y determine (utilizando división y conquista) cual es dicho elemento no repetido. Indicar y justificar la
complejidad del algoritmo implementado.
*/

func EncontarDistinto(arr []int) int {

	var repetido int
	if arr[0] == arr[1] {
		repetido = arr[0]
	} else if arr[1] == arr[2] {
		repetido = arr[1]
	} else {
		repetido = arr[2]
	}
	return distinto(arr, 0, len(arr)-1, repetido)
}

func distinto(arr []int, ini, fin, repetido int) int {

	if ini == fin {
		if arr[ini] == repetido {
			return -1
		}
	}
	medio := (ini + fin) / 2

	if arr[medio] != repetido {
		return arr[medio]
	}
	izq := distinto(arr, ini, medio, repetido)
	der := distinto(arr, medio+1, fin, repetido)

	if izq == -1 {
		return der
	}
	return izq
}

/*
2. Implementar una función que reciba un arreglo A de n enteros y un número k y devuelva un nuevo arreglo en el que
para cada posición i de dicho arreglo, contenga el resultado de la multiplicación de los primeros k máximos del arreglo A
entre las posición [0;i] (incluyendo a i). Las primeras k − 1 posiciones del arreglo a devolver deben tener como valor -1.
Por ejemplo, para el arreglo [1, 5, 3, 4, 2, 8] y k = 3, el resultado debe ser [-1, -1, 15, 60, 60, 160]. Indicar
y justificar la complejidad del algoritmo implementado.
*/

func MultipK(arr []int, k int) []int {

	res := make([]int, len(arr))

	for i := range k - 1 {
		res[i] = -1
	}

	for i := k - 1; i < len(arr); i++ {
		heap := cola_prioridad.CrearHeapArr(arr[:i+1], func(a, b int) int { return a - b })
		multip := 1
		for j := 0; j < k; j++ {
			if !heap.EstaVacia() {
				multip *= heap.Desencolar()
			}

		}
		res[i] = multip
	}
	return res
}

/* HECHO EN GUIAS/GRAFOS
3. Tenemos un plan de estudios que nos indica las correlatividades de las materias que debemos realizar. Suponer que no
hay electivas, ni correlativas por cantidad de créditos. Tenemos un alumno, al que llamaremos agus9900, que quiere
recibirse lo antes posible (es decir, en la mínima cantidad de cuatrimestres). Modelar este problema con grafos, e
implementar una función que reciba dicho grafo y devuelva una lista de listas, donde en la lista i diga qué materias hay
que cursar en el i-ésimo cuatrimestre, de tal manera de tomar la menor cantidad de cuatrimestres en recibirse. Por
supuesto, siempre debe suceder que para toda materia de la lista i, todas sus correlativas deben haberse cursado en
cuatrimestres anteriores. Pueden asumir que agus9900 es tan genio que aprobó todas las cursadas y todos los finales
(en el mismo cuatrimestre de haberlas cursado). Indicar y justificar la complejidad del algoritmo implementado en
función de la cantidad de materias del plan de estudios, y la cantidad de correlatividades.
*/

/*
5. Implementar una primitiva de árbol binario de búsqueda que devuelva un diccionario en el cual las claves sean los
niveles (int) y los datos sean listas de todos las claves del ABB que se encuentran en dicho nivel. Indicar y justificar la
complejidad del algoritmo implementado.
*/

func (ab *abb[K, V]) EncontarNiveles() diccionario.Diccionario[int, lista.Lista[K]] {

	dic := diccionario.CrearHash[int, lista.Lista[K]]()
	ab.recorrer(ab.raiz, dic, 1)
	return dic
}

func (ab *abb[K, V]) recorrer(nodo *nodoAbb[K, V], dic diccionario.Diccionario[int, lista.Lista[K]], nivel int) {

	if nodo == nil {
		return
	}
	if !dic.Pertenece(nivel) {
		lis := lista.CrearListaEnlazada[K]()
		lis.InsertarUltimo(nodo.clave)
		dic.Guardar(nivel, lis)
	} else {
		lis := dic.Obtener(nivel)
		lis.InsertarUltimo(nodo.clave)
		dic.Guardar(nivel, lis)
	}

	ab.recorrer(nodo.izquierdo, dic, nivel+1)
	ab.recorrer(nodo.derecho, dic, nivel+1)

}
