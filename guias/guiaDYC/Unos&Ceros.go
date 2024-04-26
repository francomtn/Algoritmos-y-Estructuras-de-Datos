package guiaDYC

func IndicePrimeroCero(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	if arr[0] == 0 {
		return 0
	}
	return primerCero(arr, 0, len(arr)-1)
}

func primerCero(arr []int, inicio, fin int) int {
	if inicio > fin {
		return -1
	}
	medio := (inicio + fin) / 2
	if arr[medio] == 0 {
		// Si el elemento en el medio es 0, chequeamos si el anterior es 1.
		// Si es así, hemos encontrado el primer cero.
		if medio == 0 || arr[medio-1] == 1 {
			return medio
		}
		// Si no, seguimos buscando en la mitad izquierda.
		return primerCero(arr, inicio, medio-1)
	} else {
		// Si el elemento en el medio es 1, seguimos buscando en la mitad derecha.
		return primerCero(arr, medio+1, fin)
	}
}
