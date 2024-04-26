package cola_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	TDACola "tdas/cola"
)

func TestColaVacia(t *testing.T) {
	t.Log("Pruebas con cola vacia")
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Prueba si la cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Prueba si la cola esta vacia")
}

func TestColaNoVacia(t *testing.T) {
	t.Log("Prueba la primitiva EstaVacia()")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(2)
	require.False(t, cola.EstaVacia())
}

func TestDesencolarYVerPrimeroInvalido(t *testing.T) {
	t.Log("Prueba desencolar una cola vacia")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(7)
	cola.Desencolar()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Prueba si la cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Prueba si la cola esta vacia")
}

func TestEncolarEnteros(t *testing.T) {
	t.Log("Prueba encolar valores del tipo int")
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Encolar(3)
	require.Equal(t, 1, cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, 3, cola.VerPrimero())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}

func TestEncolarCadenas(t *testing.T) {
	t.Log("Prueba encolar valores del tipo string")
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("A")
	require.Equal(t, "A", cola.VerPrimero())
	cola.Encolar("B")
	require.Equal(t, "A", cola.VerPrimero())
	cola.Desencolar()
	require.Equal(t, "B", cola.VerPrimero())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}

func TestVolumen(t *testing.T) {
	t.Log("Pruebas con muchos elementos")
	cola := TDACola.CrearColaEnlazada[int]()
	cant := 1000
	for i := 0; i < cant; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for i := 0; i < cant; i++ {
		require.Equal(t, i, cola.Desencolar())
	}
	require.True(t, cola.EstaVacia())
}

func TestRecienCreada(t *testing.T) {
	t.Log("Prueba que al encolar y desencolar hasta que este vacia, la cola se comporta como recien creada")
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Prueba si la cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Prueba si la cola esta vacia")
	cola.Encolar(8)
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Prueba si la cola esta vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Prueba si la cola esta vacia")
}
