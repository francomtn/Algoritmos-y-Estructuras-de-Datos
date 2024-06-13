package parcialito1

import (
	"fmt"
	"tdas/pila"
)

/*
1. En la ciudad existe una biblioteca tan grande que todos sus visitantes pierden mucho tiempo buscando libros. Para
solucionar este problema, se adquirió un robot recepcionista que pueda darle a las personas la ubicación del libro que
desean leer. Se necesita implementar la función de búsqueda para que dicho robot funcione. Para ello, se dispone de
información con el siguiente formato:
En donde la lista de autores está ordenada alfabéticamente por autor, y cada lista de libros está ordenada alfabéticamente
por título del libro. La función de búsqueda debe tener la siguiente firma:
func buscarLibro(nombreAutor string, tituloLibro string, autores []Autor) string
La función recibe el nombre del autor, el título del libro y una lista de autores con sus respectivos libros. La función
debe devolver la ubicación del libro en la biblioteca. Además, se requiere una solución eficiente que minimice el tiempo
de espera de los usuarios. Indicar y justificar su orden de complejidad. Puede verse al dorso un ejemplo.
*/

/*
autores := []Autor{
Autor{
Nombre: "J.K. Rowling",
Libros: []Libro{
Libro{Titulo: "HP y la cámara secreta", Ubicacion: "Sección A, biblio. 4"},
},
},
Autor{
Nombre: "George Orwell",
Libros: []Libro{
Libro{Titulo: "1984", Ubicacion: "Sección B, biblio. 1"},
Libro{Titulo: "Rebelión en la granja", Ubicacion: "Sección B, biblio. 1"},
},
},
}
buscarLibro("George Orwell", "1984", autores) // Devuelve "Sección B, biblio. 1"
*/

type Libro struct {
	Titulo    string
	Ubicacion string
}

type Autor struct {
	Nombre string
	Libros []Libro
}

func buscarLibro(nombreAutor string, tituloLibro string, autores []Autor) string {

	ubicacionAutor := busquedaBinariaLibro(0, len(autores)-1, nombreAutor, autores, obtenerNombre)
	libros := autores[ubicacionAutor].Libros
	ubicTitulo := busquedaBinariaLibro(0, len(libros)-1, tituloLibro, libros, obtenerLibro)
	return libros[ubicTitulo].Ubicacion
}

func busquedaBinariaLibro[T any](ini int, fin int, elem string, arr []T, f func(T) string) int {

	if ini == fin {
		if f(arr[ini]) == elem {
			return ini
		}
		return -1
	}
	medio := (ini + fin) / 2
	if f(arr[medio]) == elem {
		return medio
	} else if f(arr[medio]) < elem {
		return busquedaBinariaLibro(medio+1, fin, elem, arr, f)
	}
	return busquedaBinariaLibro(ini, medio, elem, arr, f)
}

func obtenerNombre(autor Autor) string {
	return autor.Nombre
}

func obtenerLibro(libro Libro) string {
	return libro.Titulo
}

/*
2. Dada una implementación de Pila, implementar una función func AgregarAlFondo[T any](p Pila[T], elem T)
recursiva que agregue el elemento al fondo de la pila sin usar TDAs o estructuras auxiliares. Justificar el orden de
complejidad de la función implementada. Nota: El “fondo” de la pila es tal que el elemento sea el último en ser
desapilado
*/

func AgregarAlFondo[T any](p pila.Pila[T], elem T) {

	if p.EstaVacia() {
		p.Apilar(elem)
		return
	}
	dato := p.Desapilar()
	AgregarAlFondo(p, elem)
	p.Apilar(dato)
}

/*
Implementar un algoritmo de ordenamiento lineal que permita definir el orden en una lista de sabores de helado para
presentar en una nueva heladería (de menor a mayor). Los datos (structs) a ordenar cuentan con el nombre del sabor
(string), popularidad (de 1 a 10, siendo 1 la popularidad más alta) (int), tipo de sabor (chocolate, ddl, etc) (string)
y si es apto para celíacos o no (bool). Primero deben ir los sabores sin TACC (aptos para celíacos), y estos deben
ordenarse por popularidad, independientemente del tipo de sabor. Luego deben ir los no-aptos para celíacos, que deben
estar ordenados por popularidad. Indicar y justificar el orden de complejidad.
*/

type Helado struct {
	Sabor       string
	SinTACC     bool
	Popularidad int
}

func ordenarHelados(arr []Helado) []Helado {
	conTacc, sinTacc := bucket(arr)
	conTacc = countingSort(conTacc, 11)
	sinTacc = countingSort(sinTacc, 11)
	sinTacc = append(sinTacc, conTacc...)
	return sinTacc
}

func bucket(arr []Helado) ([]Helado, []Helado) {
	var conTacc []Helado
	var sinTacc []Helado
	for _, elem := range arr {
		if elem.SinTACC {
			conTacc = append(conTacc, elem)
		} else {
			sinTacc = append(sinTacc, elem)
		}
	}
	return conTacc, sinTacc
}
func countingSort(arr []Helado, rango int) []Helado {
	frecuencias := make([]int, rango)
	sumas := make([]int, rango)
	res := make([]Helado, len(arr))

	for _, elem := range arr {
		valor := elem.Popularidad
		frecuencias[valor]++
	}

	for i := rango - 2; i > 0; i-- {
		sumas[i] = sumas[i+1] + frecuencias[i+1]
	}

	for _, elem := range arr {
		valor := elem.Popularidad
		posicion := sumas[valor]
		res[posicion] = elem
		sumas[valor]++
	}
	return res
}

func main() {
	helados := []Helado{
		{"Chocolate", true, 5},
		{"Vainilla", false, 2},
		{"Dulce de leche", true, 1},
		{"Fresa", false, 4},
		{"Menta", true, 3},
	}

	ordenados := ordenarHelados(helados)

	for _, helado := range ordenados {
		fmt.Println(helado)
	}
}
