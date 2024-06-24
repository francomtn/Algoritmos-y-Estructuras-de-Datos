package guiaTDA

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (pila *pilaDinamica[T]) Transformar(aplicar func(T) T) *pilaDinamica[T] {
	// para crear una nueva pila en este ejercicio usar, según la capacidad deseada:
	// nueva := &pilaDinamica[T]{make([]T, capacidadDeseada), capacidadDeseada}

	capacidadDeseada := pila.cantidad
	nueva := &pilaDinamica[T]{make([]T, capacidadDeseada), 0}

	for i := 0; i < pila.cantidad; i++ {
		dato := aplicar(pila.datos[i])
		nueva.datos[nueva.cantidad] = dato
		nueva.cantidad++
	}

	return nueva
}
