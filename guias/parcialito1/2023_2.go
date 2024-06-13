package parcialito1

/*
1. Implementar un algoritmo que reciba un arreglo de enteros de tamaño n, ordenado descendentemente,
y determine en tiempo O(log n) si existe algún valor i tal que arr[i] = - nˆ2. Justificar la complejidad del algoritmo.
*/

func EncontarValor(arr []int) bool {

	n := len(arr)

	return encontrarAux(arr, 0, len(arr)-1, n)
}

func encontrarAux(arr []int, ini, fin, n int) bool {

	if ini == fin {
		if arr[fin] == -(n * n) {
			return true
		}
		return false
	}
	medio := (ini + fin) / 2
	if arr[medio] == -(n * n) {
		return true
	} else if arr[medio] < -(n * n) {
		return encontrarAux(arr, ini, medio-1, n)
	}
	return encontrarAux(arr, medio+1, fin, n)
}

//COMPLEJIDAD:

/*
T(n)= A T(n/B) + O(n^C)
	= 1 T(n/2) + O(1) ----> A = 1, B = 2, C = 0
logB(A) = log2(1) = 0 = C ---> O(n^0 log n) == O(log n)
*/

/*
2. Implementar la primitiva de listaEnlazada Extend[T any](otra *listaEnlazada[T]) que extien-
da la lista con todos los elementos de la otra lista que se pasa por parámetro. Indicar y justificar el
orden del algoritmo.
*/

func (l *listaEnlazada[T]) Extend(otra *listaEnlazada[T]) {

	if l.EstaVacia() {
		nuevo := crearNodo(otra.primero.dato, nil)
		l.primero = nuevo
		l.ultimo = nuevo
		otra.primero = otra.primero.siguiente
	}

	actual := l.ultimo
	prim := otra.primero
	for prim != nil {
		nuevo := crearNodo(prim.dato, nil)
		actual.siguiente = nuevo
		prim = prim.siguiente
		l.largo++
	}
	l.ultimo = prim
}

/*

3. Indicar Verdadero o Falso justificando su respuesta:
a) Radix sort puede no ser estable dependiendo de la implementación.
b) Si corremos Mergesort sobre un arreglo ordenado, la complejidad va a ser O(1).
c) La complejidad de Counting Sort es menor que la de Radix Sort.
*/
