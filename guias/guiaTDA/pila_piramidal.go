package guiaTDA

import TDAPila "tdas/pila"

func EsPiramidal(pila Pila[int]) bool {

	condicion := true
	if pila.EstaVacia() {
		return condicion
	}
	aux := TDAPila.CrearPilaDinamica[int]()

	for !pila.EstaVacia() {
		aux.Apilar(pila.Desapilar())
	}
	pila.Apilar(aux.Desapilar())

	for !aux.EstaVacia() {

		if pila.VerTope() <= aux.VerTope() {
			condicion = false
		}
		pila.Apilar(aux.Desapilar())

	}
	return condicion
}
