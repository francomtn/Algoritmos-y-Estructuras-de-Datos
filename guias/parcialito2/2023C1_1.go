package parcialito2

import "tdas/diccionario"

/*
1. Implementar en Go una primitiva de ABB (DiccionarioOrdenado) que reciba un número entero N
y devuelva la cantidad de nodos del árbol que tienen exactamente N descendientes. Indicar y justificar
la complejidad de la primitiva implementada.
*/

func (abb *abb[K, V]) TienenNHijos(n int) int {

	contador := 0
	abb._tienenNhijos(abb.raiz, n, &contador)
	return contador
}

func (abb *abb[K, V]) _tienenNhijos(nodo *nodoAbb[K, V], n int, contador *int) int {

	if nodo == nil {
		return 0
	}
	izq := abb._tienenNhijos(nodo.izq, n, contador)
	der := abb._tienenNhijos(nodo.der, n, contador)
	suma := izq + der
	if suma == n {
		*contador++
	}
	return suma + 1
}

/*
2. Sabiendo que se tiene un diccionario en que las claves y valores son cadenas, implementar una función
DiccionarioInvertirUnico que devuelva un nuevo Diccionario cuyas claves sean los valores del
original y sus valores asociados sean las claves que originalmente tenían asociados dichos valores. En
el caso algún valor esté repetido en el Diccionario original, se debe lanzar un panic con el mensaje
hay valores repetidos. Indicar y justificar la complejidad de la función.
*/

func DiccionarioInvertirUnico(dic1 diccionario.Diccionario[string, string]) diccionario.Diccionario[string, string] {

	nuevo := diccionario.CrearHash[string, string]()

	for iter := dic1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()

		if nuevo.Pertenece(valor) {
			panic("hay valores repetidos")
		} else {
			nuevo.Guardar(valor, clave)
		}
	}
	return nuevo
}

/*
3. Realizar el seguimiento de aplicar HeapSort para ordernar de menor a mayor el siguiente arreglo:
[ 3, 9, 1, 5, 8, 11, 14, 9 ].
¿Es HeapSort un algoritmo de ordenamiento estable? Justificar. Indicar y justificar la complejidad
de HeapSort.
*/

/*
Dado un arbol binario que tiene como datos numeros enteros, implementar una primitiva que reemplace
el valor de los nodos internos del arbol de manera tal que cada nodo interno tenga como valor la suma de los
valores de sus nodos hijos. No se debe modificar el valor de las hojas.
*/

func (abb *abb[K, int]) ReemplazarValores() {

	abb.modificar(abb.raiz)
}

func (abb *abb[K, int]) modificar(nodo *nodoAbb[K, int]) int {

	if nodo == nil {
		return 0
	}

	if nodo.izq == nil && nodo.der == nil {
		return nodo.dato
	}

	izq := abb.modificar(nodo.izq)
	der := abb.modificar(nodo.der)
	nodo.dato = izq + der

	return nodo.dato
}
