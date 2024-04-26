package guiaTDA

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type ListaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}
