package pila

const (
	CAPACIDAD_INICIAL     int = 5
	VALOR_REDIMENSION     int = 2
	COEFICIENTE_CAPACIDAD int = 4
)

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, CAPACIDAD_INICIAL)
	pila.cantidad = 0
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(valor T) {
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * VALOR_REDIMENSION)
	}
	pila.datos[pila.cantidad] = valor
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	capacidad := len(pila.datos)
	if pila.cantidad*COEFICIENTE_CAPACIDAD <= capacidad && capacidad > CAPACIDAD_INICIAL {
		pila.redimensionar(capacidad / VALOR_REDIMENSION)
	}
	pila.cantidad--
	return pila.datos[pila.cantidad]
}

func (pila *pilaDinamica[T]) redimensionar(nuevaCapacidad int) {
	nueva := make([]T, nuevaCapacidad)
	copy(nueva, pila.datos)
	pila.datos = nueva
}
