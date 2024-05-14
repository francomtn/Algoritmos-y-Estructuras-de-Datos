package multiconj

type MultiConjunto[K comparable] interface {

	// Guardar guarda un elemento en el multiconjunto.
	Guardar(elem K)

	// Pertenece devuelve true si el elemento se encuentra en el multiconjunto
	Pertenece(elem K) bool

	// Borrar borra una aparición del elemento en el Multiconjunto. Para ser completamente
	// borrado, el elemento debe ser borrado tantas veces como se haya guardado (con
	// Guardar). Si el elemento se guardó K veces y se lo borró M veces (con M < K), entonces
	// Pertenece(elem) seguirá dando true. Si se lo borra K veces, Pertenece deberá devolver
	// false. Puede volverse a Guardar nuevamente. Si se borra un elemento que no Pertenece,
	// la primitiva entra en pánico con el mensaje 'Elemento no esta en el multiconjunto'.
	Borrar(elem K)
}
