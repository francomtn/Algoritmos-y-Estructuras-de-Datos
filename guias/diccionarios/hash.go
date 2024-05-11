package diccionarios

type estadoParClaveDato int

const (
	VACIO = estadoParClaveDato(iota)
	BORRADO
	OCUPADO
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estadoParClaveDato
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}
