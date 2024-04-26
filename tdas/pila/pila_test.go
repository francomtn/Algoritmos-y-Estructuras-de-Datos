package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Prueba si la pila esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Prueba si la pila esta vacia")
}

func TestApilarEnteros(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(7)
	require.Equal(t, 7, pila.VerTope())
	pila.Apilar(3)
	require.Equal(t, 3, pila.VerTope())
	pila.Apilar(5)
	require.Equal(t, 5, pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, 5, pila.Desapilar())
	require.Equal(t, 3, pila.Desapilar())
	require.Equal(t, 7, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestApilarCadenas(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("Hola")
	require.Equal(t, "Hola", pila.VerTope())
	pila.Apilar("Mundo")
	require.Equal(t, "Mundo", pila.VerTope())
	pila.Apilar("Algoritmos")
	require.Equal(t, "Algoritmos", pila.VerTope())
	pila.Apilar("Estructuras")
	require.Equal(t, "Estructuras", pila.VerTope())
	require.False(t, pila.EstaVacia())
	require.Equal(t, "Estructuras", pila.Desapilar())
	require.Equal(t, "Algoritmos", pila.Desapilar())
	require.Equal(t, "Mundo", pila.Desapilar())
	require.Equal(t, "Hola", pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	cant := 1000
	for i := 0; i < cant; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope())
	}
	for i := cant - 1; i >= 0; i-- {
		require.Equal(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaNoVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(11)
	require.False(t, pila.EstaVacia())
}

func TestDesapilarYVerTopeInvalidos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(9)
	pila.Desapilar()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Ver el tope de una pila vacia es invalido")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Desapilar una pila vacia es invalido")
}

func TestRecienCreada(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Prueba si la pila esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Prueba si la pila esta vacia")
	pila.Apilar(8)
	pila.Desapilar()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Prueba si la pila esta vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Prueba si la pila esta vacia")
}
