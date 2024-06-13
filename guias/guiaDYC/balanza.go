package guiaDYC

func EncontrarJoya(joyas []int) int {
	return _encontrarJoya(joyas, 0, len(joyas)-1)
}

func _encontrarJoya(arr []int, ini, fin int) int {

	if ini == fin {
		return ini
	}

	medio := (ini + fin) / 2
	if medio%2 == 0 {
		peso := balanza(arr[ini:medio+1], arr[medio:fin])
		if peso == 1 {
			return _encontrarJoya(arr, ini, medio)
		}
		return _encontrarJoya(arr, medio, fin)
	} else {
		// Ajustar la división para manejar correctamente los casos impares
		peso := balanza(arr[ini:medio], arr[medio+1:fin+1])
		if peso == 0 {
			return medio
		}
		if peso == 1 {
			return _encontrarJoya(arr, ini, medio-1)
		}
		return _encontrarJoya(arr, medio+1, fin)
	}
}

func balanza(izquierda []int, derecha []int) int {
	sumaIzquierda := sumarPeso(izquierda)
	sumaDerecha := sumarPeso(derecha)

	if sumaIzquierda == sumaDerecha {
		return 0
	} else if sumaIzquierda > sumaDerecha {
		return 1
	} else {
		return -1
	}
}

func sumarPeso(joyas []int) int {
	suma := 0
	for _, peso := range joyas {
		suma += peso
	}
	return suma
}
