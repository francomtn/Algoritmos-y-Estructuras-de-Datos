package abb

import l "tdas/lista"

func (arbol *abb[K, V]) NivelesInverso(visitar func(clave K, dato V) bool) {

	if arbol == nil {
		return
	}

	cola := l.CrearListaEnlazada[*abb[K, V]]()
	pila := l.CrearListaEnlazada[*abb[K, V]]()

	cola.InsertarUltimo(arbol)
	for !cola.EstaVacia() {
		dato := cola.BorrarPrimero()
		pila.InsertarUltimo(dato)

		if arbol.der != nil {
			cola.InsertarUltimo(dato.der)
		}
		if arbol.izq != nil {
			cola.InsertarUltimo(dato.izq)
		}
	}

	for !pila.EstaVacia() {
		nodo := pila.BorrarPrimero()
		if !visitar(nodo.clave, nodo.dato) {
			return
		}
	}
}
