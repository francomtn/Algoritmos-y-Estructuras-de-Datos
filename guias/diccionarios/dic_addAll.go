package diccionarios

type DiccionarioAddAll[K comparable] interface {
	Diccionario[K, int]

	Add(clave K, valor int)

	AddAll(valor int)
}
