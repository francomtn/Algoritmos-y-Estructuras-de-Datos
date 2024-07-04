package finales

import p "tdas/pila"

/*
1. Un algoritmo iterativo sencillo para obtener la potencia de un número (b^n) tiene complejidad O(n). Tal como vieron en
el secundario (esperamos), sabemos que b^n = (b^(n/2))^2. Utilizar esta propiedad para implementar un algoritmo que calcule b^n,
en tiempo O(log n). Justificar la complejidad del algoritmo implementado. Recordar tener cuidado con el caso que n sea un valor impar.
*/

func PotenciaFormula(b, n int) int {

	if n == 0 {
		return 1
	}
	if n%2 == 0 {
		val := PotenciaFormula(b, n/2)
		return val * val
	} else {
		val := PotenciaFormula(b, (n-1)/2)
		return val * val * b
	}
}

/*	T(n) = AT(n/b) + O(n^c)
	T(n) = 1T(n/2) + O(1) --> A = 1, B = 2, C = 0

	logB(A) = log2(1) = 0 = C --> O(n⁰log n)  = O(log n)
*/

/*
2. Se tiene un AVL de claves números enteros (su función de comparación simplemente compara dichos valores de la
forma tradicional). Su estado inicial puede reconstruirse a partir del Preorder: 23 - 9 - 6 - 11 - 75 - 34. Hacer el
seguimiento de las siguientes inserciones, incluyendo rotaciones intermedias: 106 - 40 - 57 - 28 - 17 - 32 - 36 - 37.


3. Realizar el seguimiento de aplicar CountingSort al siguiente conjunto de equipos de fútbol, ordenando por la cantidad
de descensos que tienen (entre paréntesis se indica la cantidad en cada caso). Implementar dicho algoritmo, e indicar y
justificar la complejidad del mismo.
Olimpo (4) - Boca (0) - Almagro (3) - Rosario Central (4) - Banfield (8) -
Sarmiento (2) - Defensa y Justicia (0) - Plantense (3) - River (1) - Independiente (1) -
Estudiantes LP (2) - Racing (1) - Tigre (8) - Velez (1) - Atlanta (4) - Gimnasia LP (5)


4. Implementar una función que dada una pila, determine si la misma se encuentra ordenada (es decir, se ingresaron los
elementos de menor a mayor). La pila debe quedar en el mismo estado al original al terminar la ejecución de la función.
Indicar y justificar la complejidad de la función.
*/

func EstaOrdenada(pila p.Pila[int]) bool {

	aux := p.CrearPilaDinamica[int]()
	aux.Apilar(pila.Desapilar())
	estado := true
	for !pila.EstaVacia() {
		if pila.VerTope() < aux.VerTope() {
			aux.Apilar(pila.Desapilar())
		} else {
			estado = false
		}
	}
	for !aux.EstaVacia() {
		pila.Apilar(aux.Desapilar())
	}
	return estado
}
