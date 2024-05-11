package diccionarios

import "tdas/diccionario"

func ParElementosSumanK(arr []int, k int) bool {

	aux := diccionario.CrearHash[int, int]()
	for _, elem := range arr {
		diferencia := k - elem
		if aux.Pertenece(diferencia) {
			return true
		}
		aux.Guardar(elem, 1)
	}
	return false
}
