package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN int = 10000

// Cuando se crea una lista vacia, esta debe comportarse como tal.
func TestListaVacia(t *testing.T) {
	t.Log("Pruebas con lista vacia")
	lista := TDALista.CrearListaEnlazada[bool]()

	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "Al ver primero de lista vacia no devuelve un panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "Al ver último en lista vacia no devuelve un panic")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "Al borrar último en lista vacia no devuelve un panic")
}

func IteradorListaVacia(t *testing.T) {
	t.Log("Pruebas de iterador con lista vacia")
	lista := TDALista.CrearListaEnlazada[bool]()
	iteradorL := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorL.Borrar() }, "Al querer borrar con iterador en lista vacia no devuelve un panic")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iteradorL.VerActual() }, "Al querer ver actual con iterador en lista vacia no devuelve un panic")
}

// Insertar un elemento mediante el iterador en la posición en la que se crea, efectivamente lo añade primero y es equivalente a InsertarPrimero.
func InsertarAlInicio(t *testing.T) {
	{
		t.Log("Prueba insertar al inicio con el iterador")
		lista1 := TDALista.CrearListaEnlazada[string]()
		lista2 := TDALista.CrearListaEnlazada[string]()
		iteradorL2 := lista2.Iterador()

		elem := "prueba"

		lista1.InsertarPrimero(elem)
		iteradorL2.Insertar(elem)

		require.Equal(t, lista1.VerPrimero(), iteradorL2.VerActual())
	}
	{
		lista1 := TDALista.CrearListaEnlazada[int]()
		lista2 := TDALista.CrearListaEnlazada[int]()
		iteradorL2 := lista2.Iterador()

		elem := 12

		lista1.InsertarPrimero(elem)
		iteradorL2.Insertar(elem)

		require.Equal(t, lista1.VerPrimero(), iteradorL2.VerActual())
	}
	{
		lista1 := TDALista.CrearListaEnlazada[bool]()
		lista2 := TDALista.CrearListaEnlazada[bool]()
		iteradorL2 := lista2.Iterador()

		elem := false

		lista1.InsertarPrimero(elem)
		iteradorL2.Insertar(elem)

		require.Equal(t, lista1.VerPrimero(), iteradorL2.VerActual())
	}
}

// Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
func InsertarAlFinal(t *testing.T) {
	{
		t.Log("Insertar al final con el iterador equivale a insertar al final.")
		lista1 := TDALista.CrearListaEnlazada[string]()
		lista2 := TDALista.CrearListaEnlazada[string]()
		iteradorL2 := lista2.Iterador()

		elementosBase := [5]string{"hola", "como", "estas", "todo", "bien"}
		elementoNuevo := "Messi"

		for i := range elementosBase {
			lista1.InsertarUltimo(elementosBase[i])
			lista2.InsertarUltimo(elementosBase[i])
		}

		for iteradorL2.HaySiguiente() {
			iteradorL2.Siguiente()
		}

		iteradorL2.Insertar(elementoNuevo)
		lista1.InsertarUltimo(elementoNuevo)

		require.Equal(t, lista1.VerUltimo(), iteradorL2.VerActual())
	}
	{
		lista1 := TDALista.CrearListaEnlazada[int]()
		lista2 := TDALista.CrearListaEnlazada[int]()
		iteradorL2 := lista2.Iterador()

		elementoNuevo := 11

		for i := range 11 {
			lista1.InsertarUltimo(i)
			lista2.InsertarUltimo(i)
		}

		for iteradorL2.HaySiguiente() {
			iteradorL2.Siguiente()
		}

		iteradorL2.Insertar(elementoNuevo)
		lista1.InsertarUltimo(elementoNuevo)

		require.Equal(t, lista1.VerUltimo(), iteradorL2.VerActual())
	}
}

// Insertar un elemento en el medio se hace en la posición correcta. Por definición, el elemento que se encontraba en el medio antes debe ser el siguiente del nuevo que se inserta.
func InsertarEnMedio(t *testing.T) {
	t.Log("Prueba insertar en el medio")
	lista := TDALista.CrearListaEnlazada[int]()
	iteradorL := lista.Iterador()
	elementosBase := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	elementoNuevo := 1000

	for i := range elementosBase {
		lista.InsertarUltimo(elementosBase[i])
	}

	medio := (lista.Largo() / 2)

	for i := 0; i < medio; i++ {
		iteradorL.Siguiente()
	}

	iteradorL.Insertar(elementoNuevo)
	iteradorL.Siguiente()

	require.Equal(t, elementosBase[medio-1], iteradorL.VerActual())
}

func BorrarPrimerElemento(t *testing.T) {
	t.Log("Prueba borrar el primero con el iterador y actualiza el primero de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	iteradorL := lista.Iterador()

	for i := range 11 {
		lista.InsertarUltimo(i)
	}

	primerElementoAntes := iteradorL.VerActual()
	iteradorL.Borrar()
	primerElementoDespues := iteradorL.VerActual()

	require.NotEqual(t, primerElementoAntes, primerElementoDespues)
}

func CambiaUltimoElementoAlRemover(t *testing.T) {
	t.Log("Prueba borrar el último con el iterador, cambia el último de la lista")
	lista := TDALista.CrearListaEnlazada[int]()
	iteradorL := lista.Iterador()

	for i := range 11 {
		lista.InsertarUltimo(i)
	}

	for iteradorL.HaySiguiente() {
		iteradorL.Siguiente()
	}

	ultimoElementoAntes := iteradorL.VerActual()
	iteradorL.Borrar()
	ultimoElementoDespues := iteradorL.VerActual()

	require.NotEqual(t, ultimoElementoAntes, ultimoElementoDespues)
}

// Verifica que al remover un elemento del medio, este no esté.
func NoExisteMedioAlRemoverlo(t *testing.T) {
	t.Log("Prueba que al remover un elemento del medio, este no esté")
	lista := TDALista.CrearListaEnlazada[int]()
	iteradorL := lista.Iterador()

	for i := range 11 {
		lista.InsertarUltimo(i)
	}

	medio := (lista.Largo() / 2)

	for i := 0; i < medio; i++ {
		iteradorL.Siguiente()
	}

	elementoABorrar := iteradorL.VerActual()
	iteradorL.Borrar()
	elementoPostBorrado := iteradorL.VerActual()

	require.NotEqual(t, elementoABorrar, elementoPostBorrado)
}

// Valida guardar un elemento (insertando al inicio), y que primero, ultimo y largo sean correctos. Borrar ese único elemento hace que se comporte como lista vacía
func TestInsertarVerYBorrarPrimero(t *testing.T) {
	t.Log("Prueba guardar un elemento en lista vacia y que se actualice primero y último.")
	lista := TDALista.CrearListaEnlazada[int]()
	elem := 5

	lista.InsertarPrimero(elem)

	require.Equal(t, lista.Largo(), 1)
	require.Equal(t, lista.VerPrimero(), lista.VerUltimo())
}

func TestVolumenLista(t *testing.T) {
	{
		t.Log("Prueba de volumen con lista")
		lista := TDALista.CrearListaEnlazada[int]()
		require.True(t, lista.EstaVacia())
		for i := 0; i < _VOLUMEN; i++ {
			lista.InsertarUltimo(i)
		}
		for i := 0; i < _VOLUMEN; i++ {
			require.Equal(t, i, lista.BorrarPrimero())
		}
		require.True(t, lista.EstaVacia())
	}
	{
		t.Log("Prueba de volumen con el iterador de la lista")
		lista := TDALista.CrearListaEnlazada[int]()
		iter := lista.Iterador()
		cant := 1000

		for i := cant; i > 0; i-- {
			iter.Insertar(i)
			require.Equal(t, i, iter.VerActual())
		}
	}
}

// Probamos que si la funcion visitar devuelve false, el iterador corta
func TestIteradorInternoConCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for i := range 11 {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(v int) bool {
		if v%2 == 0 {
			suma += v
		}
		if v == 7 {
			return false
		}
		return true
	})
}

func TestIteradorInternoSinCorte(t *testing.T) {

	lista := TDALista.CrearListaEnlazada[int]()

	for i := range 11 {
		lista.InsertarUltimo(i)
	}
	suma := 0
	lista.Iterar(func(v int) bool {
		suma += v
		return true
	})
	require.Equal(t, 55, suma)
}
