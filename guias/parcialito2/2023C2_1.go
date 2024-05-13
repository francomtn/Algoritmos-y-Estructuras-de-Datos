package parcialito2

import (
	"tdas/diccionario"
	"tdas/lista"
)

//ejercicio 2

func DictMedio[K comparable, T any](dict diccionario.Diccionario[K, lista.Lista[T]]) diccionario.Diccionario[K, T] {

	res := diccionario.CrearHash[K, T]()

	for iter := dict.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !res.Pertenece(clave) {
			indice := (valor.Largo() / 2) - 1
			iterLista := valor.Iterador()
			for i := 0; i < indice; i++ {
				iterLista.Siguiente()
			}
			dato := iterLista.VerActual()
			res.Guardar(clave, dato)
		}
	}
	return res
}

//ejerico 3
