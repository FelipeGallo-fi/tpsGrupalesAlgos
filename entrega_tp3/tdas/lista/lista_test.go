package lista_test

import (
	"testing"
	TDALISTA "tp3/tdas/lista"

	"github.com/stretchr/testify/require"
)

const _VOLUMEN = 1000000

func TestListaVacia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	require.Equal(t, 0, lista.Largo(), "El largo de la lista deberia ser 0")
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia")
	require.Panics(t, func() { lista.VerPrimero() }, "Ver el primero en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerUltimo() }, "Ver el ultimo en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")
	lista.InsertarPrimero(1)
	require.Equal(t, 1, lista.VerPrimero(), "El primer elemento de la lista deberia ser 1")
	lista.BorrarPrimero()
	require.Panics(t, func() { lista.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerPrimero() }, "Ver el primero en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerUltimo() }, "Ver el ultimo en una lista vacia deberia panickear")
	iterador := lista.Iterador()
	require.Panics(t, func() { iterador.Siguiente() }, "siguiente con el iterador de una lista vacia deberia panickear")
	require.Panics(t, func() { iterador.Borrar() }, "Borrar con el iterador de una lista vacia deberia panickear")
	require.Panics(t, func() { iterador.VerActual() }, "VerActual con el iterador de una lista vacia deberia panickear")

}

func TestListaVolumen(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	for i := 0; i < _VOLUMEN; i++ {
		if i%2 == 0 {
			lista.InsertarPrimero(i)
			require.Equal(t, i, lista.VerPrimero(), "El primer elemento de la lista deberia ser %d tras insertar con InsertarPrimero", i)
		} else {
			lista.InsertarUltimo(i)
			require.Equal(t, i, lista.VerUltimo(), "El Ultimo elemento de la lista deberia ser %d tras usar InsertarUltimo", i)
		}
		require.Equal(t, i+1, lista.Largo(), "El largo de la lista deberia ser %d", i+1)
		require.False(t, lista.EstaVacia(), "La lista deberia estar vacia")

	}
	require.Equal(t, _VOLUMEN, lista.Largo(), "El largo de  mi lista deberia ser %d", _VOLUMEN)

	for !lista.EstaVacia() {
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia antes de borrar")
		elemEsperado := lista.VerPrimero()
		elemBorrado := lista.BorrarPrimero()
		require.Equal(t, elemEsperado, elemBorrado, "El elemento borrado deberia ser el correcto")
	}

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrarla por completo")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerPrimero() }, "Ver el primero en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerUltimo() }, "Ver el ultimo en una lista vacia deberia panickear")
}

func TestListaVolumenInsertarPirmero(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	for i := 0; i < _VOLUMEN; i++ {
		lista.InsertarPrimero(i)
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia despues de insertar")
	}

	for !lista.EstaVacia() {
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia antes de borrar")
		elemEsperado := lista.VerPrimero()
		elemBorrado := lista.BorrarPrimero()
		require.Equal(t, elemEsperado, elemBorrado, "El elemento borrado deberia ser el correcto")
	}

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrarla por completo")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerPrimero() }, "Ver el primero en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerUltimo() }, "Ver el ultimo en una lista vacia deberia panickear")
}

func TestListaVolumenInsertarUltimo(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	for i := 0; i < _VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia despues de insertar")
	}

	for !lista.EstaVacia() {
		require.False(t, lista.EstaVacia(), "La lista no deberia estar vacia antes de borrar")
		elemEsperado := lista.VerPrimero()
		elemBorrado := lista.BorrarPrimero()
		require.Equal(t, elemEsperado, elemBorrado, "El elemento borrado deberia ser el correcto")
	}

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrarla por completo")
	require.Panics(t, func() { lista.BorrarPrimero() }, "Borrar en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerPrimero() }, "Ver el primero en una lista vacia deberia panickear")
	require.Panics(t, func() { lista.VerUltimo() }, "Ver el ultimo en una lista vacia deberia panickear")
}

func verificarOperacionesIntercaladas[T any](t *testing.T, lista TDALISTA.Lista[T], valores []T) {
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al inicio")

	lista.InsertarPrimero(valores[0])
	require.Equal(t, valores[0], lista.VerPrimero(), "El primer elemento deberia ser el primero insertado")
	require.Equal(t, valores[0], lista.VerUltimo(), "El ultimo elemento deberia ser el primero insertado")
	require.Equal(t, 1, lista.Largo(), "El largo deberia ser 1")

	lista.InsertarUltimo(valores[1])
	require.Equal(t, valores[0], lista.VerPrimero(), "El primer elemento deberia seguir siendo el primero insertado")
	require.Equal(t, valores[1], lista.VerUltimo(), "El ultimo elemento deberia ser el segundo insertado")
	require.Equal(t, 2, lista.Largo(), "El largo deberia ser 2")

	elemBorrado := lista.BorrarPrimero()
	require.Equal(t, valores[0], elemBorrado, "El elemento borrado deberia ser el primero insertado")
	require.Equal(t, valores[1], lista.VerPrimero(), "El primer elemento ahora deberia ser el segundo insertado")
	require.Equal(t, 1, lista.Largo(), "El largo deberia ser 1")

	elemBorrado = lista.BorrarPrimero()
	require.Equal(t, valores[1], elemBorrado, "El elemento borrado deberia ser el segundo insertado")
	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia despues de borrar todo")
}

func TestOperacionesIntercaladasDistintosTipos(t *testing.T) {
	// Enteros
	listaEnteros := TDALISTA.CrearListaEnlazada[int]()
	verificarOperacionesIntercaladas(t, listaEnteros, []int{10, 20})

	// Strings
	listaStrings := TDALISTA.CrearListaEnlazada[string]()
	verificarOperacionesIntercaladas(t, listaStrings, []string{"hola", "mundo"})

	// Booleanos
	listaBooleanos := TDALISTA.CrearListaEnlazada[bool]()
	verificarOperacionesIntercaladas(t, listaBooleanos, []bool{true, false})
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

	require.NotNil(t, iterador, "El iterador no deberia ser nil")
	require.False(t, iterador.HaySiguiente(), "El iterador no deberia tener un siguiente elemento en una lista vacia")
	require.Panics(t, func() { iterador.VerActual() }, "Ver el actual en una lista vacia deberia panickear")
	require.Panics(t, func() { iterador.Siguiente() }, "Siguiente en una lista vacia deberia panickear")
	require.Panics(t, func() { iterador.Borrar() }, "Borrar en una lista vacia deberia panickear")
}

func TestIteradorAlFinal(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	require.Equal(t, 2, lista.VerUltimo(), "El último elemento debería ser 2")

	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	iterador.Insertar(10)
	require.Equal(t, 10, lista.VerUltimo(), "El último elemento de la lista debería ser 10 después de insertar cuando el iterador está al final")
	require.Equal(t, 2, lista.Largo(), "El largo de la lista debería ser 2 después de insertar 2 elementos")
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

	require.Equal(t, 3, lista.Largo(), "El largo de la lista deberia ser 3")

	iterador := iterarHastaMedio(lista)
	iterador.Insertar(4)
	require.Equal(t, 4, lista.Largo(), "El largo de la lista deberia ser 4  ")

	elementosLista := []int{1, 4, 5, 3}
	i := 0
	lista.Iterar(func(elemtno int) bool {
		require.Equal(t, elementosLista[i], elemtno, "El elemento en la posicion %d no es correcto", i)
		i++
		return true
	})
}

func TestRemoverMedio(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	for i := 0; i < 6; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, 6, lista.Largo(), "La lista tendira que ser del largo : 6")

	iterador := iterarHastaMedio(lista)

	medioBorrado := iterador.Borrar()

	require.Equal(t, 3, medioBorrado, "El elemento borrado deberia ser: 3")

	elementosLista := []int{0, 1, 2, 4, 5}
	i := 0
	lista.Iterar(func(elemtno int) bool {
		require.Equal(t, elementosLista[i], elemtno, "El elemento en la posicion %d no es correcto", i)
		i++
		return true
	})

}

func TestRemoverElmento(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	iterador := lista.Iterador()

	require.Equal(t, 1, iterador.Borrar(), "El elemento borrado deberia ser 1")
	require.Equal(t, 2, lista.VerPrimero(), "EL primer numero deberia ser 2 luego de borrar con el nuevo iterador")

}

func TestRemoverUltimoIteradorElmento(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	require.Equal(t, 1, lista.VerPrimero(), "El primer elemento deberia ser 1")
	require.Equal(t, 3, lista.VerUltimo(), "El ultimo elemento deberia ser 3")
	require.Equal(t, 3, lista.Largo(), "El largo inicial de la lista deberia ser 3")

	iter := lista.Iterador()

	for i := 0; i < lista.Largo()-1; i++ {
		iter.Siguiente()
	}

	require.Equal(t, 3, iter.VerActual(), "El ultimo elemento deberia ser 3 antes de borrarlo")

	borrado := iter.Borrar()
	require.Equal(t, 3, borrado, "El elemento borrado deberia ser 3")

	require.Equal(t, 2, lista.VerUltimo(), "El nuevo ultimo elemento deberia ser 2")
	require.Equal(t, 2, lista.Largo(), "El largo de la lista deberia ser 2 despues de borrar el ultimo elemento")

	iter = lista.Iterador()
	for lista.Largo() > 0 {
		elemActual := iter.VerActual()
		elemBorrado := iter.Borrar()
		require.Equal(t, elemActual, elemBorrado, "El elemento borrado deberia ser el actual del iterador")
		if lista.Largo() > 0 {
			require.Equal(t, lista.VerUltimo(), iter.VerActual(), "El nuevo ultimo elemento deberia coincidir con el actual del iterador")
		}
	}

	require.True(t, lista.EstaVacia(), "La lista deberia estar vacia al final")
	require.Equal(t, 0, lista.Largo(), "El largo de la lista deberia ser 0 al final")
}

func TestIterarListaVacia(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	iter := lista.Iterador()

	require.False(t, iter.HaySiguiente(), "El iterador no deberia tener un siguiente elemento en una lista vacia")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.VerActual()
	}, "VerActual deberia panickear en una lista vacia")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Siguiente()
	}, "Siguiente deberia panickear en una lista vacia")

	require.PanicsWithValue(t, "El iterador termino de iterar", func() {
		iter.Borrar()
	}, "Borrar deberia panickear en una lista vacia")

	require.True(t, lista.EstaVacia(), "La lista deberia seguir vacia despues de los panics")
	require.Equal(t, 0, lista.Largo(), "El largo de la lista deberia ser 0 despues de los panics")
}

func TestIteradorInternoPocosElementosSinCorte(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	suma := 0
	lista.Iterar(func(valor int) bool {
		suma += valor
		return true
	})

	require.Equal(t, 6, suma, "La suma de los elementos deberia ser 6")
	require.Equal(t, 3, lista.Largo(), "El largo de la lista deberia seguir siendo 3")
}

func TestIteradorInternoPocosElementosConCorte(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)

	suma := 0
	lista.Iterar(func(valor int) bool {
		suma += valor
		return valor != 2
	})

	require.Equal(t, 3, suma, "La suma de los elementos hasta el corte deberia ser 3")
	require.Equal(t, 3, lista.Largo(), "El largo de la lista deberia seguir siendo 3")
}

func TestIteradorInternoVolumen(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()
	for i := 1; i <= 100000; i++ {
		lista.InsertarUltimo(i)
	}

	suma := 0
	lista.Iterar(func(valor int) bool {
		suma += valor
		return true
	})

	require.Equal(t, 5000050000, suma, "La suma de los primeros 100000 numeros deberia ser 5000050000")
	require.Equal(t, 100000, lista.Largo(), "El largo de la lista deberia seguir siendo 100000")
}

func TestIteradorInternoSinElementosSinCorte(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	contador := 0
	lista.Iterar(func(valor int) bool {
		contador++
		return true
	})

	require.Equal(t, 0, contador, "El contador deberia ser 0 porque la lista esta vacia")
	require.True(t, lista.EstaVacia(), "La lista deberia seguir vacia")
}

func TestIteradorInternoSinElementosConCorte(t *testing.T) {
	lista := TDALISTA.CrearListaEnlazada[int]()

	contador := 0
	lista.Iterar(func(valor int) bool {
		contador++
		return false
	})

	require.Equal(t, 0, contador, "El contador deberia ser 0 porque la lista esta vacia")
	require.True(t, lista.EstaVacia(), "La lista deberia seguir vacia")
}
