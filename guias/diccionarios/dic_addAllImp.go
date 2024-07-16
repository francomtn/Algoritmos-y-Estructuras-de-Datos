package diccionarios

type dicAddAll[K comparable] struct {
	dic    Diccionario[K, int]
	addSum int
}

/*func CrearDicAddAll[K comparable]() DiccionarioAddAll[K] {
	return &dicAddAll[K]{dic: CrearHash[K, int](), addSum: 0}
}*/

func (dic *dicAddAll[K]) Guardar(clave K, valor int) {
	dic.dic.Guardar(clave, valor-dic.addSum)
}

func (dic *dicAddAll[K]) Pertenece(clave K) bool {
	return dic.dic.Pertenece(clave)
}

func (dic *dicAddAll[K]) Borrar(clave K) int {
	if !dic.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	dato := dic.dic.Borrar(clave)
	return dato
}

func (dic *dicAddAll[K]) Obtener(clave K) int {
	if !dic.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	return dic.dic.Obtener(clave) + dic.addSum
}

func (dic *dicAddAll[K]) Cantidad() int {
	return dic.dic.Cantidad()
}

func (dic *dicAddAll[K]) Add(clave K, valor int) {
	if !dic.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	dato := dic.Obtener(clave)
	dic.Guardar(clave, valor+dato)
}

func (dic *dicAddAll[K]) AddAll(valor int) {
	dic.addSum += valor
}

func (dic *dicAddAll[K]) Iterar(visitar func(clave K, dato int) bool) {
	dic.dic.Iterar(func(clave K, dato int) bool { return visitar(clave, dato+dic.addSum) })
}

func (dic *dicAddAll[K]) Iterador() IterDiccionario[K, int] {
	return dic.dic.Iterador()
}
