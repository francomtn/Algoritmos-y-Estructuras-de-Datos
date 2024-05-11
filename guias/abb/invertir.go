package abb

/*type ab struct {
    izq  *ab
    der  *ab
    dato int
}*/

func (arbol *ab) Invertir() {

	if arbol == nil {
		return
	}
	arbol.izq.Invertir()
	arbol.der.Invertir()
	arbol.der, arbol.izq = arbol.izq, arbol.der

}

/*
T(n) = 2 T(n/2) + 2 O(1)
     = 2 T(n/2) + O(n0)

Por teorema maestro:
A = 2  B = 2  C = 0
logB(A) = 1 > C  -> T(n) = O(n^(LogB(A)))

T(n) = O(n)
*/
