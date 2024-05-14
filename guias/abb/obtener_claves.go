package abb

//Implementar una primitiva para el ABB, que devuelva una lista con las claves del mismo,
//ordenadas tal que si insertáramos las claves en un ABB vacío, dicho ABB tendría la misma estructura que el árbol original.

import "tdas/lista"

func (abb *abb[K, V]) Claves() lista.Lista[K] {

	lista := lista.CrearListaEnlazada[K]()
	obtenerClaves(abb.raiz, lista)
	return lista
}

func obtenerClaves[K comparable, V any](nodo *nodoAbb[K, V], lista lista.Lista[K]) {

	if nodo == nil {
		return
	}
	lista.InsertarUltimo(nodo.clave)
	obtenerClaves(nodo.izquierdo, lista)
	obtenerClaves(nodo.derecho, lista)
}

// otra forma

func (nodo *nodoAbb[K, V]) clavesPreOrder(lista lista.Lista[K]) {
	if nodo == nil {
		return
	}
	lista.InsertarUltimo(nodo.clave)
	nodo.izquierdo.clavesPreOrder(lista)
	nodo.derecho.clavesPreOrder(lista)
}

func (nodo *nodoAbb[K, V]) Claves2() lista.Lista[K] {
	reconstruccion := lista.CrearListaEnlazada[K]()
	nodo.clavesPreOrder(reconstruccion)
	return reconstruccion
}

func (abb *abb[K, V]) Claves2() lista.Lista[K] {
	return abb.raiz.Claves2()
}
