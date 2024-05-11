package abb

import "tdas/lista"

func (abb *abb[K, V]) Mayores(clave K) lista.Lista[K] {

	lista := lista.CrearListaEnlazada[K]()
	mayoresRec(abb, abb.raiz, clave, lista)
	return lista
}

func mayoresRec[K comparable, V any](abb *abb[K, V], nodo *nodoAbb[K, V], clave K, lista lista.Lista[K]) {

	if nodo == nil {
		return
	}
	if abb.cmp(nodo.clave, clave) > 0 {
		mayoresRec(abb, nodo.izquierdo, clave, lista)
		lista.InsertarUltimo(nodo.clave)
	}
	mayoresRec(abb, nodo.derecho, clave, lista)
}
