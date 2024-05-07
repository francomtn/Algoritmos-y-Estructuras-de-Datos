package Diccionarios

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}
