package ordenamientos

type Persona struct {
	nombre       string
	nacionalidad Nacionalidad
	edad         int
}

// OrdenarFila ordena una fila de personas según las reglas especificadas
func OrdenarFila(persona []Persona) []Persona {

	// Separar niños de no niños
	esNino, noNino := dividirPorEdad(persona)
	esNino = countingSort2(esNino, 13, obtenerEdad)
	// Ordenar los no niños primero por nacionalidad y luego por edad
	noNino = countingSort2(noNino, 120, obtenerEdad)        // Ordenar no niños por edad primero
	noNino = countingSort2(noNino, 33, obtenerNacionalidad) // Luego ordenar no niños por nacionalidad

	// Concatenar los niños y los no niños
	return append(esNino, noNino...)
}

// Función para dividir personas en niños y no niños
func dividirPorEdad(personas []Persona) ([]Persona, []Persona) {
	var esNino []Persona
	var noNino []Persona

	for _, persona := range personas {
		if persona.edad <= 12 {
			esNino = append(esNino, persona)
		} else {
			noNino = append(noNino, persona)
		}
	}
	return esNino, noNino
}

// Función para obtener la nacionalidad de una persona como entero
func obtenerNacionalidad(persona Persona) int {
	return int(persona.nacionalidad)
}

// Función para obtener la edad de una persona
func obtenerEdad(persona Persona) int {
	return persona.edad
}

// Implementación de Counting Sort para ordenar personas según un criterio
func countingSort2(arr []Persona, rango int, f func(Persona) int) []Persona {
	frecuencias := make([]int, rango)
	sumas := make([]int, rango)
	res := make([]Persona, len(arr))

	// Contar las frecuencias
	for _, elem := range arr {
		valor := f(elem)
		frecuencias[valor]++
	}

	// Calcular las posiciones acumuladas
	for i := 1; i < rango; i++ {
		sumas[i] = sumas[i-1] + frecuencias[i-1]
	}

	// Colocar los elementos en la posición correcta
	for _, elem := range arr {
		valor := f(elem)
		posicion := sumas[valor]
		res[posicion] = elem
		sumas[valor]++
	}

	return res
}
