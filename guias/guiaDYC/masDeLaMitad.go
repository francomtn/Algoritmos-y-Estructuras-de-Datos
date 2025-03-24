package guiaDYC

func MasDeLaMitad(arr []int) bool {
	return true
}

/*
func MasDeLaMitad(arr []int) bool {
	if len(arr) == 0 {
		return false
	}
	majorityCandidate := findMajority(arr, 0, len(arr)-1)
	count := contador(arr, majorityCandidate)
	return count > len(arr)/2
}

// Función recursiva para encontrar el candidato a elemento mayoritario en un rango
func findMajority(arr []int, ini, fin int) int {
	// Caso base: si el segmento tiene un solo elemento
	if ini == fin {
		return arr[ini]
	}
	mid := (ini + fin) / 2
	mayoriaIzq := findMajority(arr, ini, mid)
	mayoriaDer := findMajority(arr, mid+1, fin)
	// Si los dos candidatos son iguales, ese es el elemento mayoritario
	if mayoriaIzq == mayoriaDer {
		return mayoriaIzq
	}
	// Contar ocurrencias de ambos candidatos en el rango actual
	izq := contadorInRange(arr, mayoriaIzq, ini, fin)
	der := contadorInRange(arr, mayoriaDer, ini, fin)
	// Retornar el candidato con mayor ocurrencia
	if izq > der {
		return mayoriaIzq
	} else {
		return mayoriaDer
	}
}

// Función para contar las ocurrencias de un elemento en un rango dado
func contadorInRange(arr []int, num, ini, fin int) int {
	cont := 0
	for i := ini; i <= fin; i++ {
		if arr[i] == num {
			cont++
		}
	}
	return cont
}

// Función para contar las ocurrencias de un elemento en todo el arreglo
func contador(arr []int, num int) int {
	count := 0
	for _, v := range arr {
		if v == num {
			count++
		}
	}
	return count
}
*/
