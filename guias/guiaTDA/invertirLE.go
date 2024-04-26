package guiaTDA

func (lista *ListaEnlazada[T]) Invertir() {
	if lista.primero == nil {
		return
	}
	var ant *nodoLista[T]
	act := lista.primero
	sig := act.siguiente
	for act != nil {
		act.siguiente = ant
		ant = act
		act = sig
		if sig != nil {
			sig = sig.siguiente
		}
	}
	lista.primero = ant
}
