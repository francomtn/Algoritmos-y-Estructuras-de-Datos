package guiaTDA

import TDAPila "tdas/pila"

func balanceado(texto string) bool {
	pilaAux := TDAPila.CrearPilaDinamica[rune]()

	for _, caracter := range texto {
		if caracter == '[' || caracter == '{' || caracter == '(' {
			pilaAux.Apilar(caracter)
		} else if caracter == ']' {
			if pilaAux.EstaVacia() || pilaAux.Desapilar() != '[' {
				return false
			}
		} else if caracter == '}' {
			if pilaAux.EstaVacia() || pilaAux.Desapilar() != '{' {
				return false
			}
		} else if caracter == ')' {
			if pilaAux.EstaVacia() || pilaAux.Desapilar() != '(' {
				return false
			}
		}
	}
	return pilaAux.EstaVacia()
}
