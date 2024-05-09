package diccionario

import (
	"fmt"
)

const (
	COEFICIENTE_REDIMENSION float64 = 0.7
	FACTOR_REDIMENSION      int     = 2
	TAMANO                  int     = 7
)

type Estado int

const (
	VACIO Estado = iota
	OCUPADO
	BORRADO
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado Estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int // Solo hace referencia a ocupados
	tam      int
	borrados int
}

type iterDiccionario[K comparable, V any] struct {
	hash *hashCerrado[K, V]
	pos  int
}

func crearCeldaHash[K comparable, V any]() celdaHash[K, V] {
	nuevaCelda := new(celdaHash[K, V])
	nuevaCelda.estado = VACIO
	return *nuevaCelda
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {

	nuevoHash := new(hashCerrado[K, V])
	nuevoHash.tam = TAMANO
	nuevoHash.cantidad = 0
	nuevoHash.borrados = 0

	nuevoHash.tabla = crearTabla[K, V](nuevoHash.tam)

	return nuevoHash
}

func crearTabla[K comparable, V any](capacidad int) []celdaHash[K, V] {
	nuevaTabla := make([]celdaHash[K, V], capacidad)
	for i := 0; i < capacidad; i++ {
		nuevaCelda := crearCeldaHash[K, V]()
		nuevaTabla[i] = nuevaCelda
	}
	return nuevaTabla
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

	hash.redimensionar(FACTOR_REDIMENSION * hash.tam)

	posicion := hash.buscar(clave)
	celdaActual := &hash.tabla[posicion]

	if celdaActual.estado == VACIO {
		celdaActual.estado = OCUPADO
		celdaActual.clave = clave
		hash.cantidad++
	}
	celdaActual.dato = dato
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {

	posicion := hash.buscar(clave)
	return hash.tabla[posicion].estado != VACIO
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.buscar(clave)
	celdaActual := hash.tabla[posicion]
	if celdaActual.estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}
	return celdaActual.dato

}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {

	hash.redimensionar(FACTOR_REDIMENSION * hash.tam)

	posicion := hash.buscar(clave)
	celdaActual := &hash.tabla[posicion]

	if celdaActual.estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}
	celdaActual.estado = BORRADO
	hash.cantidad--
	hash.borrados++
	return celdaActual.dato
}

func (hash *hashCerrado[K, V]) Iterar(visitar func(clave K, valor V) bool) {

	for _, elem := range hash.tabla {

		if elem.estado == VACIO || elem.estado == BORRADO {
			continue
		}
		if !visitar(elem.clave, elem.dato) {
			break
		}

	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {

	for ind, elemento := range hash.tabla {
		if elemento.estado == OCUPADO {
			return &iterDiccionario[K, V]{hash: hash, pos: ind}
		}
	}

	return &iterDiccionario[K, V]{hash: hash, pos: hash.tam}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {

	return iter.pos != iter.hash.tam
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	clave := iter.hash.tabla[iter.pos].clave
	dato := iter.hash.tabla[iter.pos].dato
	return clave, dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {

	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	iter.pos++

	for iter.pos < iter.hash.tam {
		if iter.hash.tabla[iter.pos].estado == OCUPADO {
			break
		}
		iter.pos++
	}

	iter.pos = iter.hash.tam
}

func (hash *hashCerrado[K, V]) redimensionar(nuevaCapacidad int) {

	factorCarga := float64((hash.cantidad + hash.borrados)) / float64(hash.tam)

	if factorCarga >= COEFICIENTE_REDIMENSION {
		tablaAnterior := hash.tabla
		hash.tabla = crearTabla[K, V](nuevaCapacidad)
		hash.tam = nuevaCapacidad
		hash.borrados = 0
		hash.cantidad = 0

		for _, elem := range tablaAnterior {
			if elem.estado == OCUPADO {
				K, V := elem.clave, elem.dato
				hash.Guardar(K, V)
			}
		}
	}
}

func convertirABytes[K comparable](clave K) []byte {

	return []byte(fmt.Sprintf("%v", clave))
}

func sdbmHash(data []byte) uint64 {
	var hash uint64

	for _, b := range data {
		hash = uint64(b) + (hash << 6) + (hash << 16) - hash
	}

	return hash
}

func (hash *hashCerrado[K, V]) hashear(clave K) int {
	claveByte := convertirABytes(clave)
	hashing := sdbmHash(claveByte)
	resultado := hashing % uint64(hash.tam)
	return int(resultado)

}

func (hash hashCerrado[K, V]) buscar(clave K) int {

	posicion := hash.hashear(clave)

	for (hash.tabla[posicion].estado != VACIO) && !(hash.tabla[posicion].estado == OCUPADO && hash.tabla[posicion].clave == clave) {
		posicion = (posicion + 1) % hash.tam
	}
	return posicion

}
