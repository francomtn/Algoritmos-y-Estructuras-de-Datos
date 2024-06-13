package guiaDYC

func raiz(f func(int) int, a, b int) int {

	if a == b || a+1 == b {
		if f(a) == 0 {
			return a
		}
		if f(b) == 0 {
			return b
		}
	}
	medio := (a + b) / 2
	if f(medio) == 0 {
		return medio
	} else if f(medio) < 0 {
		return raiz(f, b, medio)
	} else {
		return raiz(f, a, medio)
	}
}
