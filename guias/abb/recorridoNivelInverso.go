package abb

import l "tdas/lista"

func (arbol *abb[K, V]) NivelesInverso(visitar func(clave K, dato V) bool) {

	if arbol == nil {
		return
	}

	cola := l.CrearListaEnlazada[*nodoAbb[K, V]]()
	pila := l.CrearListaEnlazada[*nodoAbb[K, V]]()

	cola.InsertarUltimo(arbol.raiz)
	for !cola.EstaVacia() {
		dato := cola.BorrarPrimero()
		pila.InsertarUltimo(dato)

		if dato.derecho != nil {
			cola.InsertarUltimo(dato.derecho)
		}
		if dato.izquierdo != nil {
			cola.InsertarUltimo(dato.izquierdo)
		}
	}

	for !pila.EstaVacia() {
		nodo := pila.BorrarPrimero()
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}
}
