package Diccionarios

import "tdas/diccionario"

func MasDeLaMitad(arr []int) bool {

	if len(arr) < 2 {
		return true
	}
	dic := diccionario.CrearHash[int, int]()
	for _, elem := range arr {
		if dic.Pertenece(elem) {
			valor := dic.Obtener(elem)
			valor++
			dic.Guardar(elem, valor)
			if valor > len(arr)/2 {
				return true
			}
		} else {
			dic.Guardar(elem, 1)
		}

	}
	return false
}
