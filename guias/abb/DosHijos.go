package abb

func (arbol *ab) DosHijos() int {

	if arbol == nil {
		return 0
	}
	cant := 0
	if arbol.der != nil && arbol.izq != nil {
		cant = 1
	}
	return arbol.der.DosHijos() + arbol.izq.DosHijos() + cant
}

//O(n)
