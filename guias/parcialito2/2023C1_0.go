package parcialito2

import (
	"tdas/cola"
	"tdas/diccionario"
	"tdas/pila"
)

/*
1. Implementar en Go una primitiva de ABB (DiccionarioOrdenado) que funcione como un iterador interno que haga un
recorrido por niveles inverso. Es decir, que visite los elementos del nivel más inferior hasta la raiz.
*/
type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func (ab *abb[K, V]) RecorrioNivelesInverso(visitar func(clave K, dato V) bool) {

	if ab == nil {
		return
	}
	cola := cola.CrearColaEnlazada[*nodoAbb[K, V]]()
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	cola.Encolar(ab.raiz)
	for !cola.EstaVacia() {
		actual := cola.Desencolar()
		pila.Apilar(actual)
		if actual.izq != nil {
			cola.Encolar(actual.izq)
		}
		if actual.der != nil {
			cola.Encolar(actual.der)
		}
	}
	for !pila.EstaVacia() {
		actual := pila.Desapilar()
		clave, valor := actual.clave, actual.dato
		if !visitar(clave, valor) {
			break
		}
	}
}

/*
2. Para una app de fútbol nos pidieron crear un registro del historial de partidos de la selección Argentina, de modo
de que se puedan ver por cuántos partidos tenemos de hijo ventaja a las otras selecciones. Se pide crear una función
func paternidad(resuls []Resultado) []Resumen que recibe un slice de Resultados, con la info de cada partido y
devuelva un slice de Resumenes con el resumen de victorias por contrincante (ver al dorso la definición de los structs).
Indicar y justificar la complejidad de la función implementada. Ejemplo:
entrada: [["Francia", "v"], ["Croacia", "v"], ["Arabia", "d"], ["Uruguay", "e"],
["Croacia", "d"], ["Francia", "v"]]
salida: [["Francia", 2], ["Croacia", 0], ["Arabia", -1], ["Uruguay", 0]]
*/

type Resultado struct {
	Pais      string
	Resultado string // "v" si ganamos, "d" perdimos y "e" empatamos
}
type Resumen struct {
	pais    string
	ventaja int // de partidos de ventaja sobre el contrincante
}

func Paternidad(results []Resultado) []Resumen {

	var res []Resumen
	dic := diccionario.CrearHash[string, int]()

	for _, elem := range results {
		pais := elem.Pais
		if !dic.Pertenece(pais) {
			valor := obtenerRes(elem.Resultado)
			dic.Guardar(pais, valor)
		} else {
			valor := dic.Obtener(pais)
			valor += obtenerRes(elem.Resultado)
			dic.Guardar(pais, valor)
		}
	}
	for iter := dic.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		res = append(res, Resumen{clave, valor})
	}
	return res
}

func obtenerRes(cad string) int {

	switch cad {
	case "v":
		return 1
	case "d":
		return -1
	}
	return 0
}

/*
3. Implementar en Go una primitiva Invertir() para el Heap, que haga que el heap se comporte con la función de
comparación contraria a la que venía utilizando hasta ese momento. El heap debe quedar en estado correcto para que las
operaciones siguientes sean válidas considerando la función de comparación, ahora invertida. Todas las operaciones deben
seguir teniendo la complejidad esperada de un Heap. No se puede modificar la estructura del heap para implementar
esta primitiva. Es decir, La implementación actual que tienen del heap debería poder trabajar con esta primitiva,
sin tener que modificar ninguna primitiva ni función auxilia ni agregar campos a la estructura. Indicar y justificar la
complejidad de la función implementada.
*/

type colaConPrioridad[T any] struct {
	datos []T
	cant  int
	cmp   func(T, T) int
}

func (h *colaConPrioridad[T]) Invertir() {

	cmpAnterior := h.cmp
	h.cmp = func(a, b T) int {
		return cmpAnterior(b, a) //  - cmpAnterior(a, b)
	}
	//heapify(h.datos, h.cmp) comento para ue no salga error pero SI se hace heapify
}
