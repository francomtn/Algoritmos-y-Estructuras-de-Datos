package abb

import l "tdas/lista"

func balancear[K comparable](claves []K) l.Lista[K] {

	lista := l.CrearListaEnlazada[K]()
	reconstruir(claves, lista)
	return lista
}

func reconstruir[K comparable](claves []K, lista l.Lista[K]) {

	if len(claves) == 0 {
		return
	}
	medio := len(claves) / 2
	lista.InsertarUltimo(claves[medio])
	reconstruir(claves[:medio], lista)
	reconstruir(claves[medio:+1], lista)
}
