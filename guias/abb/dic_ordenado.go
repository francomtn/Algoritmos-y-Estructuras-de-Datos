package abb

type DiccionarioOrdenado[K comparable, V any] interface {
	Diccionario[K, V]

	// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
	// incluyéndolos en caso de encontrarse
	IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)

	Mayores(clave K) Lista[K]
}
