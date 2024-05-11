package diccionarios

import "tdas/lista"

func Claves[K comparable, V any](dic Diccionario[K, V]) lista.Lista[K] {

	claves := lista.CrearListaEnlazada[K]()
	for celda := dic.Iterador(); celda.HaySiguiente(); celda.Siguiente() {
		clave, _ := celda.VerActual()
		claves.InsertarUltimo(clave)
	}
	return claves
}
