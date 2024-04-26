package sumapares

func SumaPares(lista Lista[*int]) int {
	suma := 0
	visitar := func(n *int) bool {
		if *n%2 == 0 {
			suma += *n
		}
		return true
	}
	lista.Iterar(visitar)
	return suma
}
