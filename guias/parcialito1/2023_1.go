package parcialito1

import "tdas/pila"

/*
1. Implementar en Go un algoritmo que, utilizando división y conquista, reciba un arreglo de enteros de tamaño n y
determine cuál es el par de números vecinos cuya suma es la mayor. Indicar y justificar el orden del algoritmo.
Ejemplo: para el arreglo [4, 8, 9, 1, 3, 6, 7, 1, 11] el par es (8, 9)
*/

func MayorSuma(arr []int) (int, int) {

	return sumaRec(arr, 0, len(arr)-1)
}

func sumaRec(arr []int, ini, fin int) (int, int) {

	if fin == ini+1 {
		return arr[ini], arr[fin]
	}

	medio := (ini + fin) / 2
	i1, i2 := sumaRec(arr, ini, medio)
	d1, d2 := sumaRec(arr, medio, fin)

	if i1+i2 > d1+d2 {
		return i1, i2
	}
	return d1, d2
}

/*
2. Implementar una función que reciba una Pila genérica y modifique su contenido tal que los elementos sean intercambiados
de a pares. No se pueden usar estructuras auxiliares. Indicar y justificar el orden del algoritmo.
Ejemplo:
Pila original
Tope <- A,B1,B2,C1,C2,D1,D2
Pila tras la función
Tope <- A,B2,B1,C2,C1,D2,D1
*/

func intercalar[T any](pila pila.Pila[T]) {

	dato := pila.Desapilar()
	_intercalar(pila)
	pila.Apilar(dato)
}

func _intercalar[T any](pila pila.Pila[T]) {

	if pila.EstaVacia() {
		return
	}
	dato1 := pila.Desapilar()
	dato2 := pila.Desapilar()
	_intercalar(pila)
	pila.Apilar(dato1)
	pila.Apilar(dato2)
}

/*
3.Se tiene un arreglo con las distancias en años luz entre la Tierra y distintos cuerpos celestes en la Via Lactea y otras
galaxias ceranas descubiertos hasta el momento. Debido a un error cometido por el sysadmin del Planetario, estas
distancias hoy están completamente desordenadas. Esta persona necesita volver a ordenar los registros antes de que
descubran el error para no perder su trabajo. Como no cursó Algoritmos 2 no sabe cómo hacer, por lo que nos pide
ayuda a nosotros. Diseñar un algoritmo de ordenamiento que pueda ordenar un arreglo de CuerpoCeleste donde
cada uno tiene un campo distancia int que funcione lo más rápido posible. Indicar y justificar el orden del
algoritmo.
De los datos, se sabe lo siguiente:
• Hay alrededor de 500 mil de registros
• Las distancias son muy dispares, yendo desde unos pocos años luz (como Alfa Centauri, la estrella mas cercana a
la Tierra que está a 4 años luz) a un poco más de 20 millones (como ADS 7251, una de las estrellas que conforman
la constelación Osa Mayor).
*/
