package parcialito2

import "tdas/diccionario"

/*5. Implementar una primitiva de árbol binario de búsqueda que devuelva un diccionario en el cual las claves sean los
niveles (int) y los datos sean listas de todos las claves del ABB que se encuentran en dicho nivel. Indicar y justificar la
complejidad del algoritmo implementado.*/

func (ab *abb[K, V]) identificarNiveles() diccionario.Diccionario[int, []K] {

	dic := diccionario.CrearHash[int, []K]()
	ab.identificarAux(ab.raiz, dic, 1)
	return dic
}

func (ab *abb[K, V]) identificarAux(nodo *nodoAbb[K, V], dic diccionario.Diccionario[int, []K], nivel int) {

	if nodo == nil {
		return
	}
	if !dic.Pertenece(nivel) {
		arr := []K{}
		arr = append(arr, nodo.clave)
		dic.Guardar(nivel, arr)
	} else {
		arr := dic.Obtener(nivel)
		arr = append(arr, nodo.clave)
		dic.Guardar(nivel, arr)
	}
	ab.identificarAux(nodo.izq, dic, nivel+1)
	ab.identificarAux(nodo.der, dic, nivel+1)

}
