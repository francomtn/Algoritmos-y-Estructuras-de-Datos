package cola

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	nodo.prox = nil
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (cola colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola colaEnlazada[T]) VerPrimero() T {
	// Devuelve el valor del primer elemento de la cola sin desencolarlo.
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}
func (cola *colaEnlazada[T]) Encolar(dato T) {
	nuevo := crearNodo(dato)
	if cola.EstaVacia() {
		cola.primero = nuevo
	} else {
		cola.ultimo.prox = nuevo
	}
	cola.ultimo = nuevo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	// Desencola y devuelve el primer elemento de la cola.
	// PRECONDICION: la cola no puede estar vacia.
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		// La cola quedo vacia al desencolar
		cola.ultimo = nil
	}
	return dato
}
