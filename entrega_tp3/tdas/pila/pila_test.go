package pila_test

import (
	"testing"
	TDAPila "tp3/tdas/pila"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 100000

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	require.Panics(t, func() { pila.Desapilar() }, "Desapilar una pila vacia deberia dar un panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

	pila.Apilar(6)
	require.False(t, pila.EstaVacia(), "La pila no tendria que estar vacia")

	tope := pila.Desapilar()
	require.Equal(t, 6, tope, "El tope tendira que ser 6")

	require.True(t, pila.EstaVacia(), "La pila tendria que estar vacia")

}

func TestApilarLifo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(25)
	pila.Apilar(10)
	pila.Apilar(2)

	require.False(t, pila.EstaVacia(), "La pila no tendria que estar vacia")

	tope := pila.Desapilar()

	require.Equal(t, 2, tope, "El tope tendira que ser 2")

	tope = pila.Desapilar()
	require.Equal(t, 10, tope, "El tope tendira que ser 10")

	tope = pila.Desapilar()

	require.Equal(t, 25, tope, "El tope tendira que ser 25")

	require.True(t, pila.EstaVacia(), "La pila  tendria que estar vacia")

}

func VolumenPila(t *testing.T) {

	pila := TDAPila.CrearPilaDinamica[int]()

	for i := 1; i <= _VOLUMEN; i++ {
		pila.Apilar(i)

		require.Equal(t, i, pila.VerTope(), "El tope deberia ser %d", i)
	}

	for i := _VOLUMEN; i >= 1; i-- {
		tope := pila.Desapilar()
		require.Equal(t, i, tope, "El tope tendira que ser %d", i)

	}
	require.True(t, pila.EstaVacia(), "La pila deberia esta vacia")

	require.Panics(t, func() { pila.VerTope() }, "VerTope en pila vacia deberia ocurrir un panic")
	require.Panics(t, func() { pila.Desapilar() }, "Desapilar en pila vacia deberia ocurrir un panic")

}

func PilaInicial(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	require.Panics(t, func() { pila.Desapilar() }, "Desapilar una pila recien creada deberia dar panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia")

	pila.Apilar(25)
	pila.Apilar(10)
	pila.Apilar(2)

	pila.Desapilar()
	pila.Desapilar()
	pila.Desapilar()

	require.Panics(t, func() { pila.Desapilar() }, "Desapilar una pila recien creada deberia dar panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

}

func PilaString(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()

	require.Panics(t, func() { pila.Desapilar() }, "Desapilar una pila recien creada deberia dar panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

	palabra1 := "Hola"
	palabra2 := "Mundo"

	pila.Apilar(palabra1)
	pila.Apilar(palabra2)

	require.Equal(t, pila.VerTope(), "Mundo", "El tope deberia ser %s", palabra2)

	pila.Desapilar()
	require.Equal(t, pila.VerTope(), "Hola", "El tope deberia ser %s", palabra1)
	pila.Desapilar()
	require.Panics(t, func() { pila.Desapilar() }, "Desapilar deberia devolver panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

}

func PilaFLoat(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()

	require.Panics(t, func() { pila.Desapilar() }, "Desapilar una pila recien creada deberia dar panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

	flot1 := 12.3
	flot2 := 0.9

	pila.Apilar(flot1)
	pila.Apilar(flot2)

	require.Equal(t, pila.VerTope(), flot2, "El tope deberia ser %s", flot2)

	pila.Desapilar()
	require.Equal(t, pila.VerTope(), flot1, "El tope deberia ser %s", flot1)
	pila.Desapilar()
	require.Panics(t, func() { pila.Desapilar() }, "Desapilar deberia devolver panic")
	require.Panics(t, func() { pila.VerTope() }, "Ver tope de  una pila vacia deberia dar un panic")

}
func PilaDePilas(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()

	pilaInt := TDAPila.CrearPilaDinamica[int]()
	pilaInt.Apilar(73)
	pilaInt.Apilar(37)

	pilaFloar64 := TDAPila.CrearPilaDinamica[float64]()
	pilaFloar64.Apilar(32.01)
	pilaFloar64.Apilar(2.31)

	pilaBoolean := TDAPila.CrearPilaDinamica[bool]()
	pilaBoolean.Apilar(true)
	pilaBoolean.Apilar(false)

	pila.Apilar(pilaInt)
	pila.Apilar(pilaFloar64)
	pila.Apilar(pilaBoolean)

	desapiladaBool, ok := pila.Desapilar().(TDAPila.Pila[bool])
	require.True(t, ok, "El elemento desapilado deberia ser booleano")
	require.Equal(t, false, desapiladaBool.Desapilar(), "El primer elemento desapilado deberia ser false")
	require.Equal(t, true, desapiladaBool.Desapilar(), "El segundo elemento desapilado deberia ser true")
	require.True(t, desapiladaBool.EstaVacia(), "La pila desapilada de bool deberia estar vacia luego de desapilarla dos veces ")

	desapilada, ok := pila.Desapilar().(TDAPila.Pila[float64])
	require.True(t, ok, "El elemento desapilado deberia ser una pila de float64")
	require.Equal(t, 2.31, desapilada.Desapilar(), "El primer elemento desapilado deberia ser 2,31")
	require.Equal(t, 32.01, desapilada.Desapilar(), "El segundo elemento desapilado deberia ser 32,01")
	require.True(t, desapilada.EstaVacia(), "La pila desapilada de float deberia estar vacia luego de desapilarla dos veces")

	desapiladaInt, ok := pila.Desapilar().(TDAPila.Pila[int])
	require.True(t, ok, "El elemento desapilado deberia ser una pila de int")
	require.Equal(t, 37, desapiladaInt.Desapilar(), "El primer elemento desapilado deberia ser 37")
	require.Equal(t, 73, desapiladaInt.VerTope(), "El tope despues de desapilar una vez deberia ser 73")
	desapiladaInt.Desapilar()
	require.True(t, desapiladaInt.EstaVacia(), "La pila de int deberia estar vacia luego de desapilar dos veces")

}

func ApilarNull(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[*int]()

	var valor *int = nil
	pila.Apilar(valor)

	require.False(t, pila.EstaVacia(), "La pila no deberia estar vacia ,deberia tener un nil despues de apilarlo")
	require.Nil(t, pila.VerTope(), "El tope deberia ser nil")
	require.Nil(t, pila.Desapilar(), "El valor desapilado deberia ser nil")

}

func DesapilarIgual(t *testing.T) {
	//lo uso para apilar y desapilar el mismo elementos varias veces

	pila := TDAPila.CrearPilaDinamica[string]()

	var palabra string = "HOLAMUNDO"
	pila.Apilar(palabra)
	pila.Apilar(palabra)
	pila.Apilar(palabra)

	require.Equal(t, palabra, "El primer desapilado deberia ser %s", palabra)
	require.Equal(t, palabra, "El segundo desapilado deberia ser %s", palabra)
	require.Equal(t, palabra, "El tercero desapilado deberia ser %s", palabra)

	require.Panics(t, func() { pila.Desapilar() }, "Desapilar vacio deberia ser panic")

	for i := 0; i < _VOLUMEN; i++ {
		pila.Apilar("HOLA")
		pila.Desapilar()

	}
	require.True(t, pila.EstaVacia(), "La pila deberia estar vacia despues de apilar y desapilar el mismo elemento muchas veces")

}
func PilaStruc(t *testing.T) {
	type Alumno struct {
		Nombre   string
		Apellido string
		Padron   int
	}
	pila := TDAPila.CrearPilaDinamica[Alumno]()
	pila.Apilar(Alumno{"Federico", "Usy", 112694})
	pila.Apilar(Alumno{"Alan", "Turing", 000001})

	require.Equal(t, "Turing", pila.Desapilar().Apellido, "Desapilar el apellido deberia devolver Turing")
	require.Equal(t, 112694, pila.VerTope().Padron, "El tope de padron deberia devolver 112694")
	require.Equal(t, Alumno{"Federico", "Usy", 112694}, pila.Desapilar(), "Desapilar deberia devolver {Federico,Usy,112694}")

}

func PilaBoolean(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[bool]()
	pila.Apilar(true)
	pila.Apilar(true)
	pila.Apilar(false)
	pila.Apilar(true)

	require.Equal(t, true, pila.Desapilar(), "Desapilar deberia devolver True")
	require.Equal(t, false, pila.Desapilar(), "Desapilar deberia devolver False")
	pila.Desapilar()
	require.Equal(t, true, pila.VerTope(), "El tope deberia ser true")
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "La pila tendira que estar vacia")

}
