package cola_prioridad

const (
	CAPACIDAD_INICIAL int = 10
	COEF_REDIMENSION  int = 4
	VALOR_REDIMENSION int = 2
)

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	nuevo := make([]T, CAPACIDAD_INICIAL)
	return &colaConPrioridad[T]{datos: nuevo, cant: 0, cmp: funcion_cmp}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	var datos []T
	if len(arreglo) == 0 {
		return CrearHeap(funcion_cmp)
	} else {
		datos = make([]T, len(arreglo))
		copy(datos, arreglo)
		heapify(datos, funcion_cmp)
	}
	return &colaConPrioridad[T]{datos: datos, cant: len(arreglo), cmp: funcion_cmp}
}

func (heap *colaConPrioridad[T]) Cantidad() int {
	return heap.cant
}

func (heap *colaConPrioridad[T]) EstaVacia() bool {
	return heap.cant == 0
}

func (heap *colaConPrioridad[T]) VerMax() T {

	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	return heap.datos[0]
}

func (heap *colaConPrioridad[T]) Encolar(dato T) {

	cap := len(heap.datos)
	if heap.cant == cap {
		heap.redimensionar(heap.cant * VALOR_REDIMENSION)
	}

	heap.datos[heap.cant] = dato
	heap.cant++
	upHeap(heap.cant-1, heap.datos, heap.cmp)

}

func (heap *colaConPrioridad[T]) Desencolar() T {

	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}

	dato := heap.datos[0]
	heap.cant--
	if heap.cant > 0 {
		heap.datos[0] = heap.datos[heap.cant]
		downHeap(0, heap.datos[:heap.cant], heap.cmp)
	}

	cap := len(heap.datos)
	if heap.cant*COEF_REDIMENSION <= cap && cap > CAPACIDAD_INICIAL {
		heap.redimensionar(cap / VALOR_REDIMENSION)
	}

	return dato
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {

	heapify(elementos, funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		swap(&elementos[0], &elementos[i])
		downHeap(0, elementos[:i], funcion_cmp)
	}
}

func swap[T any](dato1 *T, dato2 *T) {
	*dato1, *dato2 = *dato2, *dato1
}

func downHeap[T any](i int, arr []T, func_cmp func(T, T) int) {

	if i >= len(arr) || i < 0 {
		return
	}

	hijoIzq := 2*i + 1
	hijoDer := 2*i + 2
	mayor := i

	if hijoIzq < len(arr) && func_cmp(arr[hijoIzq], arr[mayor]) > 0 {
		mayor = hijoIzq
	}

	if hijoDer < len(arr) && func_cmp(arr[hijoDer], arr[mayor]) > 0 {
		mayor = hijoDer
	}

	if mayor != i {
		swap(&arr[i], &arr[mayor])
		downHeap(mayor, arr, func_cmp)
	}
}

func upHeap[T any](i int, arr []T, func_cmp func(T, T) int) {

	padre, iPadre, tienePadre := obtenerPadre(i, arr)

	if !tienePadre || func_cmp(*padre, arr[i]) > 0 {
		return
	}

	swap(&arr[i], &arr[iPadre])
	upHeap(iPadre, arr, func_cmp)
}

func heapify[T any](arr []T, func_cmp func(T, T) int) {

	for i := len(arr) - 1; i >= 0; i-- {
		downHeap(i, arr, func_cmp)
	}
}

func obtenerPadre[T any](i int, arr []T) (*T, int, bool) {

	iPadre := (i - 1) / 2

	if i == 0 {
		return &arr[0], 0, false
	}

	return &arr[iPadre], iPadre, true
}

func (heap *colaConPrioridad[T]) redimensionar(nuevaCapacidad int) {

	nuevo := make([]T, nuevaCapacidad)
	copy(nuevo, heap.datos)
	heap.datos = nuevo
}
