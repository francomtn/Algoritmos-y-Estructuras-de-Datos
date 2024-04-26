package anteKUlt

func (lista *ListaEnlazada[T]) AnteKUltimo(k int) T {

	nodoAdelantado := lista.prim
	nodoAtrasado := lista.prim

	for i := 0; i < k; i++ {
		if nodoAdelantado.prox != nil {
		}
		nodoAdelantado = nodoAdelantado.prox
	}
	for nodoAdelantado != nil {
		nodoAtrasado = nodoAtrasado.prox
		nodoAdelantado = nodoAdelantado.prox
	}
	return nodoAtrasado.dato
}
