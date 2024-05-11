package abb

type ab struct {
	izq  *ab
	der  *ab
	dato int
}

func (arbol *ab) Altura() int {

	if arbol == nil {
		return 0
	}
	h_izq := arbol.izq.Altura()
	h_der := arbol.der.Altura()
	return maximo(h_izq, h_der) + 1
}

func maximo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
