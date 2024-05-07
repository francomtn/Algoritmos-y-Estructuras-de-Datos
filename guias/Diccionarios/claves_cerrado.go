package Diccionarios

import (
	"tdas/lista"
)

func (hash *hashCerrado[K, V]) Claves() lista.Lista[K] {

	lista := lista.CrearListaEnlazada[K]()

	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO {
			lista.InsertarUltimo((celda.clave))
		}
	}
	return lista
}
