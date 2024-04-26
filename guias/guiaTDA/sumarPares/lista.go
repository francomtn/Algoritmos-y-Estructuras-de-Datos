package sumapares

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el dato al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista. Si la lista se encontraba vacía, entra el pánico con el
	// mensaje 'La lista esta vacia'.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al inicio de la lista (el primero). Si la lista se encontraba vacía, entra en
	// pánico con el mensaje 'La lista esta vacia'.
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista (el ultimo). Si la lista se encontraba vacía, entra em
	// pánico con el mensaje 'La lista esta vacia'.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar aplica la funcion pasada por parametro a todos los elementos de la lista, hasta que no hayan más
	// elementos, o la función en cuestión devualva false.
	Iterar(func(T) bool)
}
