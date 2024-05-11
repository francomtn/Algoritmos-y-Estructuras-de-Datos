package abb

/*type ab struct {
    izq *ab
    der *ab
    dato int
}*/

func (arbol *ab) Quiebres() int {

	if arbol == nil {
		return 0
	}
	izq := arbol.izq.Quiebres()
	der := arbol.der.Quiebres()

	suma := izq + der
	if arbol.der != nil && arbol.der.der == nil && arbol.der.izq != nil {
		suma++
	}
	if arbol.izq != nil && arbol.izq.izq == nil && arbol.izq.der != nil {
		suma++
	}
	return suma
}
