package ordenamientos

func OrdenarNotacionCientifica(arr []string) []string {
	countingSort(arr, 10, 1)
	countingSort(arr, 10, 0)
	countingSort(arr, 10, 2)
	return arr
}

func countingSort(elementos []string, rango int, digito int) {
	frecuencias := make([]int, rango)
	sumasAcumuladas := make([]int, rango)
	resultado := make([]string, len(elementos))

	for _, elem := range elementos {
		valor := int(elem[digito] - '0')
		frecuencias[valor]++
	}

	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}

	for _, elem := range elementos {
		valor := int(elem[digito] - '0')
		pos := sumasAcumuladas[valor]
		resultado[pos] = elem
		sumasAcumuladas[valor]++
	}

	for i := 0; i < len(resultado); i++ {
		elementos[i] = resultado[i]
	}
}
