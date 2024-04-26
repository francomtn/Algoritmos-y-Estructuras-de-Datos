package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func crearNodo[T any](dato T) nodoLista[T] {

	return nodoLista[T]{dato: dato}
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = &nuevoNodo
	} else {
		nuevoNodo.siguiente = lista.primero

	}
	lista.primero = &nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevoNodo := crearNodo(dato)

	if lista.EstaVacia() {
		lista.primero = &nuevoNodo
	} else {
		lista.ultimo.siguiente = &nuevoNodo

	}
	lista.ultimo = &nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	} else {
		return lista.primero.dato
	}

}

func (lista *listaEnlazada[T]) VerUltimo() T {

	if lista.EstaVacia() {
		panic("La lista esta vacia")
	} else {
		return lista.ultimo.dato
	}
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {

	actual := lista.primero
	iterarHastaFinal := true

	for actual != nil || !iterarHastaFinal {
		iterarHastaFinal = visitar(actual.dato)
		actual = actual.siguiente
	}
}

type iterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: lista, actual: lista.primero}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("Fin de la iteracion")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("Fin de la iteracion")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nuevoNodo := crearNodo(dato)

	//la lista esta vacia
	if iterador.anterior == nil && iterador.actual == nil {
		iterador.actual = &nuevoNodo
	}
	iterador.anterior.siguiente = &nuevoNodo
	nuevoNodo.siguiente = iterador.actual
	iterador.actual = &nuevoNodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if iterador.HaySiguiente() {
		panic("Fin de la iteracion")
	}
	if iterador.anterior == nil {
		datoBorrado := iterador.actual.dato
		iterador.actual = nil
		iterador.lista.largo--
		return datoBorrado
	} else {
		datoBorrado := iterador.actual.dato
		iterador.anterior.siguiente = iterador.actual.siguiente
		iterador.actual = iterador.actual.siguiente
		if iterador.actual.siguiente == nil {
			iterador.lista.ultimo = nil
		}
		iterador.lista.largo--
		return datoBorrado

	}
}
