package diccionarios

func SonIguales[K comparable, V comparable](d1, d2 Diccionario[K, V]) bool {

	if d1.Cantidad() != d2.Cantidad() {
		return false
	}
	iter := d1.Iterador()
	for iter.HaySiguiente() {
		clave, valor := iter.VerActual()
		if d2.Pertenece(clave) {
			if d2.Obtener(clave) != valor {
				return false
			}
		} else {
			return false
		}
		iter.Siguiente()
	}
	return true
}
