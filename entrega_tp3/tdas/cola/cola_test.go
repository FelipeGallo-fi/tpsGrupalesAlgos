package cola_test

import (
	"testing"
	TDACola "tp3/tdas/cola"

	"github.com/stretchr/testify/require"
)

func TestColaVaciaInts(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	require.Panics(t, func() { cola.VerPrimero() }, "La cola esta vacia")
	require.Panics(t, func() { cola.Desencolar() }, "La cola esta vacia")
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "La cola esta vacia")
	require.Panics(t, func() { cola.Desencolar() }, "La cola esta vacia")
}

func TestColaVaciaStrings(t *testing.T) {
	colaCadenas := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaCadenas.EstaVacia())
	require.Panics(t, func() { colaCadenas.VerPrimero() }, "La cola esta vacia")
	require.Panics(t, func() { colaCadenas.Desencolar() }, "La cola esta vacia")

	colaCadenas.Encolar("hola")
	colaCadenas.Encolar("mundo")
	colaCadenas.Encolar("go")

	require.Equal(t, "hola", colaCadenas.VerPrimero())
	require.Equal(t, "hola", colaCadenas.Desencolar())

	require.False(t, colaCadenas.EstaVacia())

	require.Equal(t, "mundo", colaCadenas.VerPrimero())
	require.Equal(t, "mundo", colaCadenas.Desencolar())

	require.Equal(t, "go", colaCadenas.VerPrimero())
	require.Equal(t, "go", colaCadenas.Desencolar())

	require.True(t, colaCadenas.EstaVacia())
	require.Panics(t, func() { colaCadenas.VerPrimero() }, "La cola esta vacia")
	require.Panics(t, func() { colaCadenas.Desencolar() }, "La cola esta vacia")
}

func TestEncolarYDesencolarInts(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)

	require.False(t, cola.EstaVacia())

	require.Equal(t, 1, cola.VerPrimero())
	require.Equal(t, 1, cola.Desencolar())

	require.Equal(t, 2, cola.VerPrimero())
	require.Equal(t, 2, cola.Desencolar())

	require.False(t, cola.EstaVacia())

	require.Equal(t, 3, cola.VerPrimero())
	require.Equal(t, 3, cola.Desencolar())

	require.Equal(t, 4, cola.VerPrimero())
	require.Equal(t, 4, cola.Desencolar())

	require.True(t, cola.EstaVacia())
}

func TestEncolarYDesencolarStrings(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("hola")
	cola.Encolar("mundo")
	cola.Encolar("go")

	require.False(t, cola.EstaVacia())

	require.Equal(t, "hola", cola.VerPrimero())
	require.Equal(t, "hola", cola.Desencolar())

	require.Equal(t, "mundo", cola.VerPrimero())
	require.Equal(t, "mundo", cola.Desencolar())

	require.False(t, cola.EstaVacia())

	require.Equal(t, "go", cola.VerPrimero())
	require.Equal(t, "go", cola.Desencolar())

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "La cola esta vacia")
	require.Panics(t, func() { cola.Desencolar() }, "La cola esta vacia")
}

func TestVolumen(t *testing.T) {
	const volumen = 100000
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())

	for i := range volumen {
		cola.Encolar(i)
		require.Equal(t, 0, cola.VerPrimero())
	}

	require.False(t, cola.EstaVacia())

	for i := range volumen {
		require.Equal(t, i, cola.VerPrimero())
		require.Equal(t, i, cola.Desencolar())
	}

	require.True(t, cola.EstaVacia())
	require.Panics(t, func() { cola.VerPrimero() }, "La cola esta vacia")
	require.Panics(t, func() { cola.Desencolar() }, "La cola esta vacia")
}
