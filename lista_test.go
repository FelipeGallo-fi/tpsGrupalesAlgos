package lista_test

import (
	TDALISTA "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 1000000

func TestListaVacia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.BorrarPrimero() }, "BorrarPrimero en un lista recien  deberia panickear")
	require.Panics(t, func() { lista.VerPrimero() }, "VerPimero en una lista recien creada deberia panickear")
}

func TestListaVolumen(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	for i := 0; i < _VOLUMEN; i++ {

		if i%2 == 0 {
			lista.InsertarPrimero(i)
			require.Equal(t, i, lista.VerPrimero(), "El primer elemento de la lista deberia ser %d", i)
		} else {
			lista.InsertarUltimo(i)
			require.Equal(t, i, lista.VerUltimo(), "El Ultimo elemento de la lista deberia ser %d", i)
		}
	}
	require.Equal(t, _VOLUMEN, lista.Largo(), "El largo de  mi lista deberia ser %d", _VOLUMEN)

	for !lista.EstaVacia() {
		lista.BorrarPrimero()
	}

	require.True(t, lista.EstaVacia(), "La lista deberia esta vacia despues de borrarla por completo")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")

}

func TestListas(t *testing.T) {
	// Lista de enteros
	listaEnteros := TDALISTA.CrearListaEnlazada[int]()
	listaEnteros.InsertarPrimero(1)
	require.Equal(t, 1, listaEnteros.VerPrimero())
	listaEnteros.InsertarUltimo(4)
	require.Equal(t, 4, listaEnteros.VerUltimo())
	listaEnteros.BorrarPrimero()
	require.Equal(t, 4, listaEnteros.VerPrimero())
	listaEnteros.BorrarPrimero()
	require.PanicsWithValue(t, "La lista esta vacia =(", func() {
		listaEnteros.VerPrimero()
	})

	// Lista de strings
	listaStrings := TDALISTA.CrearListaEnlazada[string]()
	listaStrings.InsertarPrimero("hola")
	require.Equal(t, "hola", listaStrings.VerPrimero())
	listaStrings.InsertarUltimo("mundo")
	require.Equal(t, "mundo", listaStrings.VerUltimo())
	listaStrings.BorrarPrimero()
	require.Equal(t, "mundo", listaStrings.VerPrimero())
	listaStrings.BorrarPrimero()
	require.PanicsWithValue(t, "La lista esta vacia =(", func() {
		listaStrings.VerPrimero()
	})

	// Lista de booleanos
	listaBooleanos := TDALISTA.CrearListaEnlazada[bool]()
	listaBooleanos.InsertarPrimero(true)
	require.Equal(t, true, listaBooleanos.VerPrimero())
	listaBooleanos.InsertarUltimo(false)
	require.Equal(t, false, listaBooleanos.VerUltimo())
	listaBooleanos.BorrarPrimero()
	require.Equal(t, false, listaBooleanos.VerPrimero())
	listaBooleanos.BorrarPrimero()
	require.PanicsWithValue(t, "La lista esta vacia =(", func() {
		listaBooleanos.BorrarPrimero()
	})
}

func TestListaListas(t *testing.T) {
	//Una lista de listas
	listaListas := TDALISTA.CrearListaEnlazada[TDALISTA.Lista[int]]()
	listaListas.InsertarPrimero(TDALISTA.CrearListaEnlazada[int]())
	require.Equal(t, 0, listaListas.VerPrimero().Largo())
	listaListas.VerPrimero().InsertarPrimero(1)
	require.Equal(t, 1, listaListas.VerPrimero().VerPrimero())
	listaListas.InsertarUltimo(TDALISTA.CrearListaEnlazada[int]())
	require.Equal(t, 0, listaListas.VerUltimo().Largo())
	listaListas.BorrarPrimero()
	require.Equal(t, 0, listaListas.VerPrimero().Largo())
	listaListas.BorrarPrimero()
	require.Panics(t, func() { listaListas.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")
}

func TestInsertar(t *testing.T) {
	listaEnteros := TDALISTA.CrearListaEnlazada[int]()

	for i := 0; i < 1000; i++ {
		if i%2 == 0 {
			listaEnteros.InsertarPrimero(i)
			require.Equal(t, i, listaEnteros.VerPrimero())
		} else {
			listaEnteros.InsertarUltimo(i)
			require.Equal(t, i, listaEnteros.VerUltimo())
		}
	}

	require.Equal(t, 999, listaEnteros.VerUltimo())
	require.Equal(t, 998, listaEnteros.VerPrimero())
}

func TestCrearIterador(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	iterador.Insertar(3)

	require.Equal(t, 3, lista.VerPrimero(), "El primer elemento de mi lista deberia ser 3")
	require.Equal(t, 1, lista.Largo(), "El largo de mi lista deberia ser 1")
}

func TestIteradorAlFinal(t *testing.T) {

	lista := TDALISTA.CrearListaEnlazada[int]()

	lista.InsertarUltimo(2)

	require.Equal(t, 2, lista.VerUltimo(), "El ultimo elemento deberia ser 2")
	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	iterador.Insertar(10)

	require.Equal(t, 10, lista.VerUltimo(), "El ultimo elemento de mi lista deberia ser 10 despues de insertar cuando mi iterador esta en el final")

	require.Equal(t, 2, lista.Largo(), "El largo de mi lista deberia ser 2 despues de insertar 2 elementos")
}

func iterarHastaMedio[T any](lista TDALISTA.Lista[T]) TDALISTA.IteradorLista[T] {
	iterador := lista.Iterador()
	largoLista := (lista.Largo() / 2)
	for i := 0; i < largoLista; i++ {
		iterador.Siguiente()
	}
	return iterador
}

func TestInsertarMedio(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(3)

	largo := lista.Largo()
	//el medio deberia ser 5

	iterador := iterarHastaMedio(lista)
	iterador.Insertar(4)
	//el nuevo medio deberia ser  4

	//el nuevo largo deberia ser 1 mas
	require.Equal(t, largo+1, lista.Largo(), "El largo de la lista deberia ser 1 mas largo  ")
	largo = lista.Largo()

	for i := 0; i < (largo/2)-1; i++ {
		lista.BorrarPrimero()
	}
	require.Equal(t, 4, lista.VerPrimero(), "El medio deberia ser: 4")
}

func TestRemoverMedio(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	for i := _VOLUMEN - 1; i >= 0; i-- {
		lista.InsertarPrimero(i)
	}
	require.Equal(t, _VOLUMEN, lista.Largo(), "La lista tendira que ser del largo : %d", _VOLUMEN)

	iterador := iterarHastaMedio(lista)

	medioBorrado := iterador.Borrar()

	require.Equal(t, _VOLUMEN/2, medioBorrado, "El elemento borrado deberia ser: %d", _VOLUMEN/2)

	for i := 0; i < (_VOLUMEN/2)-1; i++ {
		lista.BorrarPrimero()
	}
	medioNuevo := lista.VerPrimero()
	require.Equal(t, (_VOLUMEN/2)-1, medioNuevo, "El nuevo medio deberia ser distiano al anterior")
}

func TestRemoverElmento(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()
	iterador.Borrar()

	require.Equal(t, 2, lista.VerPrimero(), "EL primer numero deberia ser 2 luego de borrar con el nuevo iterador")

}

func TestRemoverUltimoIteradorElmento(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter = lista.Iterador()
	var actual int
	for iter.HaySiguiente() {
		actual = iter.VerActual()
		iter.Siguiente()
	}
	iter = lista.Iterador()
	for iter.HaySiguiente() && iter.VerActual() != actual {
		iter.Siguiente()
	}
	iter.Borrar()
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 2, lista.Largo())
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	})

	require.False(t, iter.HaySiguiente())
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	iterador := lista.Iterador()

	for i := 0; i < 100; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, 100, lista.Largo(), "El largo de  mi lista deberia ser 100")

	for i := 0; iterador.HaySiguiente(); i++ {
		require.Equal(t, i, iterador.VerActual(), "El valor actual del iterador deberia ser %d", i)
		iterador.Siguiente()
	}

}
