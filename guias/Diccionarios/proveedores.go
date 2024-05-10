package Diccionarios

import "tdas/diccionario"

func MergeProveedores(prov1, prov2 Diccionario[string, int]) diccionario.Diccionario[string, int] {

	aux := diccionario.CrearHash[string, int]()
	for iter := prov1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !prov2.Pertenece(clave) {
			aux.Guardar(clave, valor)
		} else {
			valor2 := prov2.Obtener(clave)
			if valor < valor2 {
				aux.Guardar(clave, valor)
			} else {
				aux.Guardar(clave, valor2)
			}
		}
	}
	for iter := prov2.Iterador(); iter.HaySiguiente(); iter.Siguiente() { // Repite para guardar los que hayan quedado en d2
		clave, valor := iter.VerActual()
		if !aux.Pertenece(clave) {
			aux.Guardar(clave, valor)
		}
	}
	return aux
}
