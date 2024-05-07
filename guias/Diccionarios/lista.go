package Diccionarios

type Lista[T any] interface {

	//EstaVacia devuelve true si la lista se encuentra vacia, de lo contrario, devuelve false.
	EstaVacia() bool

	// InsertarPrimero inserta un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el primer elemento de la lista. En caso de que este vacia, entra en pánico con un mensaje "La lista está vacia"
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. En caso de que este vacia, entra en pánico con un mensaje "La lista está vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. En caso de que este vacia, entra en pánico con un mensaje "La lista está vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos alojados en la lista.
	Largo() int

	// Iterar recorre la lista, aplicandole a cada elemento la función pasada por parámetro hasta encontrar una condición de corte.
	// Si devuelve true, la función itera hasta el final. Si devuelve false, se aplicó una condición de corte.
	Iterar(visitar func(T) bool)

	// Iterador  crea un iterador externo, asociado a la lista actual.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento de la posición actual del iterador.
	// Si no hay un elemento en esa posición (es nil), lanza un panic.
	VerActual() T

	//HaySiguiente devuelve true si hay un dato para ver en la posición del iterador, de lo contrario, devuelve false.
	HaySiguiente() bool

	// Siguiente avanza el iterador a la posición siguiente.
	Siguiente()

	// Insertar inserta el dato pasado por parametro entre el anterior y el actual.
	// El iterador que en la posición del nuevo dato.
	Insertar(T)

	// Borrar quita y devuelve el dato de la posición actual del iterador.
	// El iterador queda apuntando al que era siguiente al dato borrado.
	Borrar() T
}
