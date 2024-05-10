package Diccionarios

import "tdas/diccionario"

func DiferenciaSimetricaDict[K comparable, V any](d1, d2 Diccionario[K, V]) diccionario.Diccionario[K, V] {

	aux := diccionario.CrearHash[K, V]()
	for iter := d1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !d2.Pertenece(clave) {
			aux.Guardar(clave, valor)
		}
	}
	for iter := d2.Iterador(); iter.HaySiguiente(); iter.Siguiente() { // Repite para guardar los que hayan quedado en d2
		clave, valor := iter.VerActual()
		if !d1.Pertenece(clave) {
			aux.Guardar(clave, valor)
		}
	}
	return aux
}
