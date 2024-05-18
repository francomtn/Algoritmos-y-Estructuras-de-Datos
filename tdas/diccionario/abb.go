package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type funcCmp[K comparable] func(K, K) int

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

type iterABB[K comparable, V any] struct {
	desde *K
	hasta *K
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	arbol *abb[K, V]
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {

	nuevoArbol := new(abb[K, V])

	nuevoArbol.cantidad = 0
	nuevoArbol.cmp = funcion_cmp

	return nuevoArbol
}

func crearNodo[K comparable, V any](clave K, valor V) *nodoAbb[K, V] {

	return &nodoAbb[K, V]{
		clave:     clave,
		dato:      valor,
		izquierdo: nil,
		derecho:   nil,
	}
}

func (arbol *abb[K, V]) buscar(padre **nodoAbb[K, V], clave K) **nodoAbb[K, V] {

	if *padre == nil {
		return padre
	}

	comparacion := arbol.cmp(clave, (*padre).clave)

	if comparacion == 0 {
		return padre
	}

	if comparacion < 0 {
		return arbol.buscar(&((*padre).izquierdo), clave)
	}

	return arbol.buscar(&((*padre).derecho), clave)
}

func (arbol *abb[K, V]) Guardar(clave K, valor V) {
	busqueda := arbol.buscar(&arbol.raiz, clave)

	if *busqueda == nil {
		*busqueda = crearNodo(clave, valor)
		arbol.cantidad++
	} else {
		(*busqueda).dato = valor
	}
}

func (arbol *abb[K, V]) Pertenece(clave K) bool {

	busqueda := arbol.buscar(&arbol.raiz, clave)
	return *busqueda != nil
}

func (arbol *abb[K, V]) Cantidad() int {
	return arbol.cantidad
}

func (arbol *abb[K, V]) Borrar(clave K) V {

	busqueda := arbol.buscar(&arbol.raiz, clave)

	if *busqueda == nil {
		panic("La clave no pertenece al diccionario")
	}

	dato := (*busqueda).dato

	if (*busqueda).izquierdo == nil && (*busqueda).derecho == nil {
		*busqueda = nil
	} else if (*busqueda).izquierdo != nil && (*busqueda).derecho == nil {
		(*busqueda) = (*busqueda).izquierdo
	} else if (*busqueda).izquierdo == nil && (*busqueda).derecho != nil {
		(*busqueda) = (*busqueda).derecho
	} else {
		izquierdoMasDerecho := buscarMasDerecho[K, V](&(*busqueda).izquierdo)
		clave, valor := (**izquierdoMasDerecho).clave, (**izquierdoMasDerecho).dato
		arbol.Borrar((*izquierdoMasDerecho).clave)
		arbol.cantidad++
		(*busqueda).clave, (*busqueda).dato = clave, valor
	}

	arbol.cantidad--
	return dato
}

func (arbol *abb[K, V]) Obtener(clave K) V {

	busqueda := arbol.buscar(&arbol.raiz, clave)

	if *busqueda == nil {
		panic("La clave no pertenece al diccionario")
	}

	return (*busqueda).dato
}

func (arbol *abb[K, V]) Iterar(visitar func(clave K, valor V) bool) {

	arbol.iteradorInterno(arbol.raiz, visitar)
}

func (arbol *abb[K, V]) iteradorInterno(nodoActual *nodoAbb[K, V], visitar func(clave K, valor V) bool) bool {

	if nodoActual != nil {
		if !arbol.iteradorInterno(nodoActual.izquierdo, visitar) {
			return false
		}

		if !visitar(nodoActual.clave, nodoActual.dato) {
			return false
		}

		if !arbol.iteradorInterno(nodoActual.derecho, visitar) {
			return false
		}
	}

	return true

}

func (arbol *abb[K, V]) Iterador() IterDiccionario[K, V] {

	iter := arbol.IteradorRango(nil, nil)
	return iter
}

func (iter *iterABB[K, V]) apilarSiguientes(primerNodo *nodoAbb[K, V]) {

	if primerNodo == nil {
		return
	}

	comparacionDesde := iter.arbol.cmp(*iter.desde, primerNodo.clave)
	comparacionHasta := iter.arbol.cmp(*iter.hasta, primerNodo.clave)

	if comparacionDesde <= 0 && comparacionHasta >= 0 {
		iter.pila.Apilar(primerNodo)
		iter.apilarSiguientes(primerNodo.izquierdo)

	} else if comparacionDesde > 0 {
		iter.apilarSiguientes(primerNodo.derecho)
	} else if comparacionHasta < 0 {
		iter.apilarSiguientes(primerNodo.izquierdo)
	}

}

func (iter *iterABB[K, V]) HaySiguiente() bool {

	return !iter.pila.EstaVacia()
}

func (iter *iterABB[K, V]) Siguiente() {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	nodoActual := iter.pila.Desapilar()
	iter.apilarSiguientes(nodoActual.derecho)
}

func (iter *iterABB[K, V]) VerActual() (K, V) {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	clave := iter.pila.VerTope().clave
	valor := iter.pila.VerTope().dato

	return clave, valor
}

func (arbol *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	if desde == nil {
		minimo := buscarMasIzquierdo[K, V](&arbol.raiz)
		desde = &(*minimo).clave
	}

	if hasta == nil {
		maximo := buscarMasDerecho[K, V](&arbol.raiz)
		hasta = &(*maximo).clave
	}

	arbol._iterarRango(arbol.raiz, desde, hasta, visitar)
}

func (arbol *abb[K, V]) _iterarRango(nodoActual *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool) {

	if nodoActual == nil {
		return
	}

	if arbol.cmp(nodoActual.clave, *desde) >= 0 {
		arbol._iterarRango(nodoActual.izquierdo, desde, hasta, visitar)
	}

	if arbol.cmp(nodoActual.clave, *desde) >= 0 && arbol.cmp(nodoActual.clave, *hasta) <= 0 {
		if !visitar(nodoActual.clave, nodoActual.dato) {
			return
		}
	}

	if arbol.cmp(nodoActual.clave, *hasta) <= 0 {
		arbol._iterarRango(nodoActual.derecho, desde, hasta, visitar)
	}
}

func (arbol *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

	iter := new(iterABB[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.arbol = arbol

	if arbol.raiz == nil {
		return iter
	}

	if desde != nil {
		iter.desde = desde
	} else {
		minimo := buscarMasIzquierdo[K, V](&arbol.raiz)
		iter.desde = &(*minimo).clave
	}

	if hasta != nil {
		iter.hasta = hasta
	} else {
		maximo := buscarMasDerecho[K, V](&arbol.raiz)
		iter.hasta = &(*maximo).clave
	}

	iter.apilarSiguientes(arbol.raiz)
	return iter
}

func buscarMasDerecho[K comparable, V any](padre **nodoAbb[K, V]) **nodoAbb[K, V] {

	if (*padre) == nil || (*padre).derecho == nil {
		return padre
	}
	return buscarMasDerecho(&(*padre).derecho)
}

func buscarMasIzquierdo[K comparable, V any](padre **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*padre) == nil || (*padre).izquierdo == nil {
		return padre
	}
	return buscarMasIzquierdo(&(*padre).izquierdo)
}
