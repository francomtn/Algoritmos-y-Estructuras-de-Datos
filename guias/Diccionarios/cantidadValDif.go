package Diccionarios

import "tdas/diccionario"

func (dicc *hashCerrado[K, V]) CantidadValoresDistintos() int {

	aux := diccionario.CrearHash[V, bool]()

	for _, celda := range dicc.tabla {
		if celda.estado == OCUPADO {
			if !aux.Pertenece(celda.dato) {
				aux.Guardar(celda.dato, true)
			}
		}
	}
	return aux.Cantidad()
}
