package guiaDYC

func ParteEnteraRaiz(n int) int {

	return _raiz(n, 0, n)
}

func _raiz(n, desde, hasta int) int {

	if desde == hasta {
		return desde
	}
	medio := (desde + hasta) / 2
	prox := medio + 1

	if medio*medio <= n && prox*prox > n {
		return medio
	}
	if medio*medio < n {
		return _raiz(n, medio+1, hasta)
	} else {
		return _raiz(n, desde, medio-1)
	}
}
