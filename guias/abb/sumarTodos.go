package abb

/*type ab struct {
    izq  *ab
    der  *ab
    dato int
}*/

func (arbol *ab) Suma() int {
	if arbol == nil {
		return 0
	}
	return arbol.dato + arbol.izq.Suma() + arbol.der.Suma()
}

//O(n)
