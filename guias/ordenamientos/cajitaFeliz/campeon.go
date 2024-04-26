package cajitaFeliz

type Persona struct {
	nombre       string
	nacionalidad Nacionalidad
	edad         int
}

func OrdenarFila(persona []Persona) []Persona {
	esNiño, noNiño := DividirPorEdad(persona)
	countingSort(esNiño, 13, obtenerEdad)
	countingSort(noNiño, 120, obtenerEdad)
	countingSort(noNiño, 33, obtenerNacionalidad)
	esNiño = append(esNiño, noNiño...)
	return esNiño
}

func DividirPorEdad(persona []Persona) ([]Persona, []Persona) {

	var esNiño []Persona
	var noNiño []Persona

	for _, elem := range persona {
		if elem.edad < 13 {
			esNiño = append(esNiño, elem)
		} else {
			noNiño = append(noNiño, elem)
		}
	}
	return esNiño, noNiño
}

func obtenerNacionalidad(persona Persona) int {
	return int(persona.nacionalidad)

}
func obtenerEdad(persona Persona) int {
	return persona.edad
}

func countingSort(arr []Persona, rango int, f func(Persona) int) []Persona {
	frecuencias := make([]int, rango)
	sumas := make([]int, rango)
	res := make([]Persona, len(arr))

	for _, elem := range arr {
		valor := f(elem)
		frecuencias[valor]++
	}
	for i := 1; i < rango; i++ {
		sumas[i] = sumas[i-1] + frecuencias[i-1]
	}
	for _, elem := range arr {
		valor := f(elem)
		posicion := sumas[valor]
		res[posicion] = elem
		sumas[valor]++
	}

	return res
}
