package guiaDYC

func PosicionPico(v []int, ini, fin int) int {
	if ini == fin {
		return ini
	}

	medio := (ini + fin) / 2

	if v[medio] > v[medio-1] && v[medio] > v[medio+1] {
		return medio
	} else if v[medio] < v[medio+1] {
		return PosicionPico(v, medio, fin)
	} else {
		return PosicionPico(v, ini, medio)
	}

}
