package heap

import (
	heap "tdas/cola_prioridad"
	"tdas/pila"
)

type Elem[T any] struct {
	dato      T
	prioridad int
}

type pilaSobreHeap[T any] struct {
	heap     ColaPrioridad[Elem[T]]
	contador int
}

func CrearPilaSobreColaPrioridad[T any]() pila.Pila[T] {
	return &pilaSobreHeap[T]{heap: heap.CrearHeap(func(a, b Elem[T]) int { return a.prioridad - b.prioridad })}
}

func (pila *pilaSobreHeap[T]) Apilar(dato T) {

	elem := Elem[T]{dato: dato, prioridad: pila.contador}
	pila.heap.Encolar(elem)
	pila.contador++
}

func (pila *pilaSobreHeap[T]) Desapilar() T {
	if pila.heap.EstaVacia() {
		panic("La pila esta vacia")
	}
	elem := pila.heap.Desencolar()
	return elem.dato
}

func (pila *pilaSobreHeap[T]) VerTope() T {
	if pila.heap.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.heap.VerMax().dato

}

func (pila *pilaSobreHeap[T]) EstaVacia() bool {
	if pila.heap.EstaVacia() {
		return true
	} else {
		return false
	}
}
