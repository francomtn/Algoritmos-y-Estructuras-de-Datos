package guiaTDA

func Largo[T any](pila Pila[T]) int {
	if pila.EstaVacia() {
		return 0
	}
	elem := pila.Desapilar()

	contador := 1 + Largo(pila)
	pila.Apilar(elem)
	return contador
}
