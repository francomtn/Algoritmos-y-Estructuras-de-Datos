package finales

import "tdas/diccionario"

/*
1. Dados dos arreglos ordenados A y B, donde B tiene “un elemento menos que A”, implementar un algoritmo de división y
conquista que permita obtener el valor faltante de A en B. Ejemplo, si A = {2, 4, 6, 8, 9, 10, 12} y B = {2, 4, 6, 8, 10, 12},
entonces la salida del algoritmo debe ser o bien la posición 4, o el valor 9 (lo que decidan que devuelva).
Indicar y justificar adecuadamente la complejidad del algoritmo implementado.
*/

func EncontrarFaltante(arr1, arr2 []int) int {

	return faltante(arr1, arr2, 0, len(arr1)-1)
}

func faltante(arr1, arr2 []int, ini, fin int) int {

	if ini == fin {
		if arr1[ini] == arr2[ini] {
			return -1
		} else {
			return arr1[ini]
		}
	}
	medio := (ini + fin) / 2
	if arr1[medio] != arr2[medio] {
		return arr1[medio]
	}
	izq := faltante(arr1, arr2, ini, medio-1)
	der := faltante(arr1, arr2, medio+1, fin)
	if izq == -1 {
		return der
	}
	return izq
}

/*
2. Implementar un algoritmo que dado un grafo dirigido y acíclico, obtenga el camino mínimo de un vértice v a otro w,
en tiempo lineal (de vértices y aristas). Justificar la complejidad del algoritmo implementado. Pista: utilizar un
recorrido DFS.
*/

/*
3. Realizar el seguimiento de aplicar BucketSort al siguiente conjunto de equipos de fútbol, ordenando por la cantidad
de descensos que tienen (entre paréntesis se indica la cantidad en cada caso). Para este caso aplicar la versión de
BucketSort que trabaja con valores discretos con baldes unitarios (no por rangos). Implementar dicho algoritmo, e
indicar y justificar la complejidad del mismo.

Olimpo (4) - Boca (0) - Almagro (3) - Rosario Central (4) - Banfield (8) -
Sarmiento (2) - Defensa y Justicia (0) - Plantense (3) - River (1) - Independiente (1) -
Estudiantes LP (2) - Racing (1) - Tigre (8) - Velez (1) - Atlanta (4) - Gimnasia LP (5)
*/
type Equipos struct {
	nombre   string
	cantidad int
}

func OrdenarEquipos(arr []Equipos) []Equipos {

	maxDescensos := 0
	for _, equipo := range arr {
		if equipo.cantidad > maxDescensos {
			maxDescensos = equipo.cantidad
		}
	}

	buckets := make(map[int][]Equipos)

	for i := range maxDescensos {
		buckets[i] = []Equipos{}
	}
	for _, equipo := range arr {
		cant := equipo.cantidad
		buckets[cant] = append(buckets[cant], equipo)
	}
	resultado := []Equipos{}
	for i := 0; i <= maxDescensos; i++ {
		resultado = append(resultado, buckets[i]...)
	}
	return resultado
}

/*
4. Explicar cuál o cuáles son los problemas de la siguiente función de hashing para ser utilizada para implementar un
Hash que funciona únicamente con string como clave (para evitar ambigüedades, suponer que la operación % largo se
aplica luego sobre el resultado de la ejecución de dicha función):
func hashing(clave string) int {
return rand.IntN(2) // devuelve un número al azar, 0 o 1.
}
*/

/*
RTA: al tener una funcion de hashing que solo devuelve 0 y 1 van a ocurrir muchas colisiones lo cual afecta directamente a la complejidad
temporal del TDA. al "mandar" los datos siempre a los mismos lugares, la cpmplejidad pasaria de ser O(1) a O(n), rompiendo justamente con la
idea del hash que es que sea O(1).
Ademas, al hashear la clave de forma aleatoria, una misma clave podria no obtener el mismo resultado al momento de operar, por ejemplo al momento de guardar,
se hashea con 0 y si luego se pregunta si pertenece podria ser hasheada con 1, y la idea del hash es que una misma clave siempre de hashee de una misma manera.
*/

/*-
5. Implementar una primitiva para el heap func (heap *heap[T]) DiferenciaSimetrica(otro *heap[T])
ColaPrioridad[T], que reciba otro Heap y cree un nuevo Heap con los elementos del primero que no se
ecuentren en el segundo, y viceversa (es decir, la diferencia simétrica entre ambos). La función de comparación del
nuevo heap debe ser la del primer heap. Indicar y justificar la complejidad del algoritmo implementado.
*/

type heap[T any] struct {
	datos []T
	cmp   func(T, T) int
}

func (h *heap[T]) DiferenciaSimetrica(otro *heap[T]) heap[T] {

	primero := diccionario.CrearHash[T, T]()
	segundo := diccionario.CrearHash[T, T]()

	for _, elem := range h.datos {
		if !primero.Pertenece(elem) {
			primero.Guardar(elem, elem)
		}
	}
	for _, elem := range otro.datos {
		if !segundo.Pertenece(elem) {
			segundo.Guardar(elem, elem)
		}
	}
	arr := []T{}
	for _, elem := range h.datos {
		if !segundo.Pertenece(elem) {
			arr = append(arr, elem)
		}
	}
	nuevo := CrearHeapArr(arr, h.cmp)
	return nuevo
}

/*
func main() {
	equipos := []Equipos{
		{"Olimpo", 4},
		{"Boca", 0},
		{"Almagro", 3},
		{"Rosario Central", 4},
		{"Banfield", 8},
		{"Sarmiento", 2},
		{"Defensa y Justicia", 0},
		{"Platense", 3},
		{"River", 1},
		{"Independiente", 1},
		{"Estudiantes LP", 2},
		{"Racing", 1},
		{"Tigre", 8},
		{"Velez", 1},
		{"Atlanta", 4},
		{"Gimnasia LP", 5},
	}

	fmt.Println("Equipos antes del ordenamiento:")
	for _, equipo := range equipos {
		fmt.Printf("%s (%d) - ", equipo.nombre, equipo.cantidad)
	}
	fmt.Println()

	equiposOrdenados := OrdenarEquipos(equipos)

	fmt.Println("\nEquipos después del ordenamiento:")
	for _, equipo := range equiposOrdenados {
		fmt.Printf("%s (%d) - ", equipo.nombre, equipo.cantidad)
	}
	fmt.Println()
}
*/
