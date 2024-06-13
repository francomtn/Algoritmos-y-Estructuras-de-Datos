package parcialito1

import "tdas/cola"

/*
Implementar una primitiva para una Cola implementada como una estructura en arreglo (como la vista en clase),
Filtrar[T](func condicion(T) bool) Cola[T] que devuelva una nueva cola para la cual los elementos de la cola
original dan true en la función condicion pasada por parámetro. La cola original debe quedar intacta, y los elementos
de la final deben tener el orden relativo que tenían en la original. Indicar y justificar la complejidad del algoritmo
implementado.
*/

func (cola *cola.colaEnlazada[T]) Filtrar(f func(T) bool) cola.Cola() {

	nueva := &colaEnlazada[T]{}
	actualCola := cola.primero
	for actualCola != nil {
		if f(actualCola.dato) {
			nuevo := crearNodo(actualCola.dato)
			if nueva.EstaVacia() {
				nueva.primero = nuevo				
			} else {
				nueva.ultimo.prox = nuevo
			}
			nuevo.ultimo = nuevo
		}
		actualCola = actualCola.prox
	}
	return nueva
}