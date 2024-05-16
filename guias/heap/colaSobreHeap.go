package heap

type elemento[T any] struct {
	dato      T
	prioridad int
}
type colaSobreHeap[T any] struct {
	heap     ColaPrioridad[elemento[T]]
	contador int
}

func CrearColaSobreColaPrioridad[T any]() Cola[T] {

	return &colaSobreHeap[T]{heap: CrearHeap(func(a, b elemento[T]) int { return b.prioridad - a.prioridad })}
}

func (cola *colaSobreHeap[T]) Encolar(dato T) {

	elem := elemento[T]{dato: dato, prioridad: cola.contador}
	cola.heap.Encolar(elem)
	cola.contador++
}

func (cola *colaSobreHeap[T]) Desencolar() T {

	elem := cola.heap.Desencolar()
	return elem.dato

}

func (cola *colaSobreHeap[T]) EstaVacia() bool {
	return cola.heap.EstaVacia()
}

func (cola *colaSobreHeap[T]) VerPrimero() T {

	return cola.heap.VerMax().dato
}
