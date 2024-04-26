package ordenamientos

type Alumno struct {
	nombre     string
	p1, p2, p3 int
}

func obtenerP1(a Alumno) int {
	return a.p1
}
func obtenerP2(a Alumno) int {
	return a.p2
}
func obtenerP3(a Alumno) int {
	return a.p3
}
func promedio(a Alumno) int {
	return (a.p1 + a.p2 + a.p3) / 3
}

func OrdenarAlumnosRadixSort(arr []Alumno) []Alumno {
	CountingSort(arr, 11, obtenerP3)
	CountingSort(arr, 11, obtenerP2)
	CountingSort(arr, 11, obtenerP1)
	CountingSort(arr, 11, promedio)
	return arr
}

func CountingSort(elementos []Alumno, rango int, f func(Alumno) int) []Alumno {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]Alumno, len(elementos))

	for _, elem := range elementos {
		valor := f(elem)
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}

	for _, elem := range elementos {
		valor := f(elem)
		pos := sumasAcumuladas[valor]
		resultado[pos] = elem
		sumasAcumuladas[valor]++
	}

	for i := 0; i < len(resultado); i++ {
		elementos[i] = resultado[i]
	}
	return elementos
}
