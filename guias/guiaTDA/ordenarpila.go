package guiaTDA

import TDApila "tdas/pila"

func Ordenar(pila Pila[int]) {
	aux := TDApila.CrearPilaDinamica[int]()

	for !pila.EstaVacia() {
		maximo := pila.Desapilar()

		for !aux.EstaVacia() && aux.VerTope() < maximo {
			pila.Apilar(aux.Desapilar())
		}
		aux.Apilar(maximo)
	}
	for !aux.EstaVacia() {
		pila.Apilar(aux.Desapilar())
	}
}
