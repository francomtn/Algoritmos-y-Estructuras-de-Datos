package parcialito2

import "math"

/*import (
	"tdas/diccionario"
	"math"
)*/

/*
1. Implementar para el Hash Cerrado la primitiva Limpieza(), la cual se encarga de eliminar todos los borrados,
asegurándose de dejar al Hash en un estado correcto (pista: pensar en las búsquedas de los elementos que efectivamente
se encuentran en el hash). Indicar y justificar la complejidad de la primitiva implementada.
*/
/*
func (hash *hashCerrado[K, V]) Limpieza() {

	nueva := crearTabla[K,V](hash.tam * 2)
	anterior := hash.tabla
	hash.tabla = nueva
	hash.borrados = 0

	for _, elem := range anterior {
		if elem.estado == "OCUPADO" {
			hash.Guardar(elem.clave, elem.dato)
		}
	}

}
*/

/*
3. Implementar en Go una primitiva para el árbol binario func (ab *Arbol[T]) EsCompleto() que determine si el mismo
es un árbol completo. Indicar y justificar la complejidad de la primitiva. A fines del ejercicio, considerar que el árbol
está definido como:
type ab[T any] struct {
dato T
izq *ab
der *ab
}*/

func (ab *abb[K, V]) esCompleto() bool {

	cantidad := ab.raiz.CantidadHijos()
	altura := ab.raiz.Altura()

	if cantidad == int(math.Pow(2, float64(altura))-1) {
		return true
	}
	return false
}

func (ab *nodoAbb[K, V]) CantidadHijos() int {

	if ab == nil {
		return 0
	}
	return 1 + ab.izq.CantidadHijos() + ab.der.CantidadHijos()

}

func (ab *nodoAbb[K, V]) Altura() int {

	if ab == nil {
		return 0
	}
	h_izq := ab.izq.Altura()
	h_der := ab.der.Altura()
	return maximo(h_izq, h_der) + 1
}

func maximo(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
