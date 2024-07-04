package finales

/*
2. Implementar un algoritmo que obtenga la parte entera de la raíz de un número n entero en O (log n). Justificar la
complejidad de la primitiva implementada.
*/

func RaizDyC(n int) int {

	return _raiz(n, 0, n)
}

func _raiz(n, ini, fin int) int {

	if ini == fin {
		return ini
	}
	medio := (ini + fin) / 2
	sig := medio + 1
	if medio*medio <= n && sig*sig > n {
		return medio
	} else if medio*medio < n {
		return _raiz(n, medio+1, n)
	} else {
		return _raiz(n, ini, medio-1)
	}
}

/*
3. Implementar una primitiva para una Cola implementada como una estructura en arreglo (como la vista en clase),
Filtrar[T](func condicion(T) bool) Cola[T] que devuelva una nueva cola para la cual los elementos de la cola
original dan true en la función condicion pasada por parámetro. La cola original debe quedar intacta, y los elementos
de la final deben tener el orden relativo que tenían en la original. Indicar y justificar la complejidad del algoritmo
implementado.
*/
type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

/*
func (cola *colaEnlazada[T]) Filtrar(visitar func(T) bool) Cola[T] {

	aux := CrearColaEnlazada[T]()
	prim := cola.primero
	for prim != nil {
		if visitar(prim.dato) {
			aux.Encolar(prim)
		}
		prim = prim.prox
	}
	cola.ultimo = prim
	return aux
}
*/
