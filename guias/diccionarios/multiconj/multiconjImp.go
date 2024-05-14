package multiconj

import "tdas/diccionario"

type multiConj[K comparable] struct {
	dic diccionario.Diccionario[K, int]
}

func CrearMulticonjunto[K comparable]() MultiConjunto[K] {
	return &multiConj[K]{dic: diccionario.CrearHash[K, int]()}
}

func (conj multiConj[K]) Guardar(elem K) {

	if !conj.dic.Pertenece(elem) {
		conj.dic.Guardar(elem, 1)
	} else {
		cant := conj.dic.Obtener(elem)
		conj.dic.Guardar(elem, cant+1)
	}
}

func (conj multiConj[K]) Pertenece(elem K) bool {
	return conj.dic.Pertenece(elem)
}

func (conj multiConj[K]) Borrar(elem K) {

	if !conj.dic.Pertenece(elem) {
		panic("Elemento no esta en el multiconjunto")
	}
	cant := conj.dic.Obtener(elem)
	if cant == 1 {
		conj.dic.Borrar(elem)
	} else {
		conj.dic.Guardar(elem, cant-1)
	}
}
