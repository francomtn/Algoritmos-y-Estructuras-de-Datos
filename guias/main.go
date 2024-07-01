package main

import (
	//"strings"
	//"tdas/diccionario"
	"fmt"
	"guias/guiaDYC"
	//"tdas/diccionario"
)

func main() {

	arr := []int{1, 2, 3, 4, 1, 2, 3, 4}
	fmt.Println(guiaDYC.Alternar(arr))
}

// arreglo ordenado := []int{1,2,3,5,8,11,17,21,26,30,35,56}
// arreglo desordenado := []int{2, 6, 5, 8, 4, 9, 11, -56, 42, 0, 7, 5}
// arreglo casi ordenado := []int{1,2,3,0,8,11,17,21,26,30,35,56}        []int{1,2,3,4,5,6,7,8,90,10}
// 1s y 0s := [9]int{1,1,1,0,0,1,0,1,0}
// pico := []int{0,1,2,3,4,5,7,5,4,3,2,1} >> 6      []int{2,3,4,5,6,7,6,5,4} >> 5      []int{2,3,4,5,4,3,2,1,0} >> 3
// heap [ 1, 7, 2, 8, 7, 6, 3, 3, 9, 10 ]
// mayor suma vecinos arr := {9}int[4, 8, 9, 1, 3, 6, 7, 1, 11] el par es (8, 9)
/*
pila := pila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	parcialito1.AgregarAlFondo(pila, 5)
	for !pila.EstaVacia() {
		fmt.Println(pila.Desapilar())
	}

*/

/*
helados := []Helado{
		{"Chocolate", 5, "chocolate", true},
		{"Vainilla", 2, "vainilla", false},
		{"Dulce de leche", 1, "ddl", true},
		{"Fresa", 4, "fruta", false},
		{"Menta", 3, "menta", true},
	}

	ordenados := radixHelado(helados)

	for _, helado := range ordenados {
		fmt.Println(helado)
	}
*/
/*
	arbol := diccionario.CrearABB[string, int](strings.Compare)
	arbol.Guardar("hola", 1)
	arbol.Guardar("chau", 1)
	arbol.Guardar("perro", 1)
	arbol.Guardar("gato", 1)
	arbol.Guardar("prueba", 1)
	arbol.Guardar("algo", 1)
	arbol.Guardar("estructura", 1)
	arbol.Guardar("franco", 1)
	arbol.Guardar("rosi", 1)

	dic := arbol.IdentificarNiveles()

	for iter := dic.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		clave, valor := iter.VerActual()
		fmt.Println(clave, valor)
	}

*/
