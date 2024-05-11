package diccionarios

import "tdas/lista"

func (hash *hashAbierto[K, V]) Claves() lista.Lista[K] {
	claves := lista.CrearListaEnlazada[K]()

	for _, celdas := range hash.tabla {
		lista := celdas.Iterador()
		for lista.HaySiguiente() {
			par := lista.VerActual()
			claves.InsertarUltimo(par.clave)
			lista.Siguiente()
		}
	}
	return claves
}
