package guiaTDA

import TDACola "tdas/cola"

func FiltrarCola[K any](cola Cola[K], filtro func(K) bool) {

	aux := TDACola.CrearColaEnlazada[K]()

	for !cola.EstaVacia() {
		elem := cola.Desencolar()
		if filtro(elem) {
			aux.Encolar(elem)
		}
	}
	for !aux.EstaVacia() {
		cola.Encolar(aux.Desencolar())
	}
}
