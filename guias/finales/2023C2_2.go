package finales

import (
	"math"
	h "tdas/cola_prioridad"
	"tdas/diccionario"
)

/*
1. Implementar un algoritmo que, dado un árbol binario, determine si el mismo es completo (es decir, que todos los niveles
que tenga estén completos). Indicar y justificar la complejidad del algoritmo implementado.
*/
type funcCmp[K comparable] func(K, K) int

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      funcCmp[K]
}

func (ab *abb[K, V]) EsCompleto() bool {

	cantidad := ab.raiz.CantHijos()
	altura := ab.raiz.Altura()

	return int(math.Pow(2, float64(altura))) == cantidad-1
}

func (ab *nodoAbb[K, V]) Altura() int {

	if ab == nil {
		return 0
	}
	izq := ab.izquierdo.Altura()
	der := ab.derecho.Altura()
	return max(izq, der) + 1

}

func (ab *nodoAbb[K, V]) CantHijos() int {

	if ab == nil {
		return 0
	}
	return 1 + ab.izquierdo.CantHijos() + ab.derecho.CantHijos()
}

/*
2. El 10/10 un nuevo algoritmo de ordenamiento ha sido inventado: el MessiSort. Así como no podemos entender cómo
hace Messi para jugar como lo hace, vamos a asumir que no podemos entender cómo hace este algoritmo para ordenar.
El creador del algoritmo (que nada tiene que ver con el astro argentino), declara que el mismo ordena en tiempo mejor
a O(n log n). ¿Tenés algo para decir sobre su afirmación? Si esto no fuera cierto, ¿podría utilizarse como algoritmo
auxiliar de RadixSort?
*/

/*
RTA: Para que la complejidad sea mejor que O(n log n), el algoritmos NO puede ser comparativo.
tendria que ser un algoritmo no comparativo como lo es counting, radix o bucket sort. Para poder
usar este algotimo tendriamos que saber informacion adicional.
Podria ser utilizado como ordenamiento auxiliar de radix independioentemente de si la afirmacion es correcta o no.
Lo unico de cambiaria es la complejidad final del algoritmo que se use para ordenar, que seria la resultante de usar
radix con MessiSort

por gpt:
El algoritmo MessiSort, si realmente tiene una complejidad mejor que O(nlogn), tendría que ser un algoritmo no comparativo.
Sin embargo, para utilizar este algoritmo como un algoritmo auxiliar en Radix Sort, necesitamos conocer más detalles sobre
cómo MessiSort funciona y en qué condiciones es eficiente.
Independientemente de si la afirmación sobre MessiSort es correcta o no, podríamos intentar usarlo como un algoritmo auxiliar en Radix Sort.
Si MessiSort realmente es más rápido que O(nlogn), entonces la complejidad total de Radix Sort podría mejorar. De lo contrario,
la complejidad sería dominada por la parte más lenta del proceso, ya sea Radix Sort o MessiSort.
*/

/*
3. Implementar un algoritmo que reciba dos arreglos desordenados y determine si ambos arreglos tienen los mismos
elementos (y en mismas cantidades). Indicar y justificar la complejidad del algoritmo implementado.
*/

func MismosElems(arr1, arr2 []int) bool {

	if len(arr1) != len(arr2) {
		return false
	}

	dic1 := diccionario.CrearHash[int, int]()
	dic2 := diccionario.CrearHash[int, int]()

	for _, elem := range arr1 {
		if !dic1.Pertenece(elem) {
			dic1.Guardar(elem, 1)
		} else {
			valor := dic1.Obtener(elem)
			dic1.Guardar(elem, valor+1)
		}
	}
	for _, elem := range arr2 {
		if !dic2.Pertenece(elem) {
			dic2.Guardar(elem, 1)
		} else {
			valor := dic2.Obtener(elem)
			dic2.Guardar(elem, valor+1)
		}
	}

	for iter := dic1.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		if !dic2.Pertenece(clave) {
			return false
		} else {
			dato := dic2.Obtener(clave)
			if valor != dato {
				return false
			}
		}
	}
	return true
}

/*
4. Trabajamos para una escuela primaria muy estructurada. En dicha escuela hay k cursos, cada uno con m alumnos (es
decir, hay un total de n = k · m alumnos). Todas las mañanas hay que armar filas para cantar Aurora en el patio del
colegio. Primero los alumnos se ubican en una fila correspondiente a su curso, de menor a mayor altura para cantar.
Una vez terminado, proceden a entrar a la escuela de a un alumno por vez, pero deben hacerlo de menor a mayor altura.
Es decir, se debe ordenar a todos los alumnos de menor a mayor. Nosotros sabemos que esto es ineficiente (suelen usar
mergesort, así que es O(n log n)), y desaprovechamos que los alumnos ya estaban ordenados por cursos. Implementar
un algoritmo que reciba k filas (arreglos) de alumnos, cada una previamente ordenada de menor a mayor altura, y nos
devuelva un único arreglo con todos los alumnos ordeados por altura en tiempo menor a O(n log n). Indicar y justificar
la complejidad del algoritmo implementado.
*/

type Alumnos struct {
	nombre string
	curso  string
	altura int
}

func Kmerge(arr [][]Alumnos) []Alumnos {

	heap := h.CrearHeapArr[[]Alumnos](arr, func(a, b []Alumnos) int { return b[0].altura - a[0].altura })
	res := make([]Alumnos, len(arr)*len(arr[0]))

	for i := range len(res) {
		primero := heap.Desencolar()
		res[i] = primero[0]
		if len(primero) > 1 {
			heap.Encolar(primero[1:])
		}
	}
	return res
}
