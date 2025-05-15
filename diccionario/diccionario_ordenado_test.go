package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/tpsGrupalesAlgos/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ORDENADO = []int{12500, 25000, 50000, 100000, 200000, 400000}


func cmpString(a, b string) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}

func cmpInt(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	return 0
}


func TestABBDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que ABB vacío no tiene claves")
	abb := TDADiccionario.CrearABB[string, string](cmpString)
	require.Equal(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Borrar("A") })
}

func TestABBUnElemento(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")

	abb := TDADiccionario.CrearABB[string, string](cmpString)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](cmpInt)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abbNum.Borrar(0) })
}

func TestABBGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el abb, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](cmpString)
	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestABBReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}
func TestABBResmpalzaoDatoHopscotch(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	abb:= TDADiccionario.CrearABB[int, int](cmpInt)
	for i := 0; i < 500; i++ {
		abb.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}
func TestABBBorrado(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb:= TDADiccionario.CrearABB[string, string](cmpString)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al abb", func() { abb.Obtener(claves[1]) })
}

func TestAbbReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	clave := "hola"
	abb.Guardar(clave, "mundo!")
	abb.Borrar(clave)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(clave))
	abb.Guardar(clave, "mundooo!")
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, "mundooo!", abb.Obtener(clave))
}

func TestABBConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb:= TDADiccionario.CrearABB[int, string](cmpInt)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}
/*
func TestABBConClavesStructs(t *testing.T) {
	t.Log("Valida que tambien funcione con estructuras mas complejas")
	type basico struct {
		a string
		b int
	}
	type avanzado struct {
		w int
		x basico
		y basico
		z string
	}

	abb:= TDADiccionario.CrearABB[avanzado, int]()

	a1 := avanzado{w: 10, z: "hola", x: basico{a: "mundo", b: 8}, y: basico{a: "!", b: 10}}
	a2 := avanzado{w: 10, z: "aloh", x: basico{a: "odnum", b: 14}, y: basico{a: "!", b: 5}}
	a3 := avanzado{w: 10, z: "hello", x: basico{a: "world", b: 8}, y: basico{a: "!", b: 4}}

	abb.Guardar(a1, 0)
	abb.Guardar(a2, 1)
	abb.Guardar(a3, 2)

	require.True(t, abb.Pertenece(a1))
	require.True(t, abb.Pertenece(a2))
	require.True(t, abb.Pertenece(a3))
	require.EqualValues(t, 0, abb.Obtener(a1))
	require.EqualValues(t, 1, abb.Obtener(a2))
	require.EqualValues(t, 2, abb.Obtener(a3))
	abb.Guardar(a1, 5)
	require.EqualValues(t, 5, abb.Obtener(a1))
	require.EqualValues(t, 2, abb.Obtener(a3))
	require.EqualValues(t, 5, abb.Borrar(a1))
	require.False(t, abb.Pertenece(a1))
	require.EqualValues(t, 2, abb.Obtener(a3))

}
	*/
func TestABBClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}
func TestABBValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb:= TDADiccionario.CrearABB[string, *int](cmpString)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestABBCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		abb.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, abb.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = abb.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestABBGuardarYBorrarRepetidasVeces(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	abb:= TDADiccionario.CrearABB[int, int](cmpInt)
	for i := 0; i < 1000; i++ {
		abb.Guardar(i, i)
		require.True(t, abb.Pertenece(i))
		abb.Borrar(i)
		require.False(t, abb.Pertenece(i))
	}
}

func buscarAbb(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestAbbIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb:= TDADiccionario.CrearABB[string, *int](cmpString)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1,buscarAbb(cs[0], claves))
	require.NotEqualValues(t, -1,buscarAbb(cs[1], claves))
	require.NotEqualValues(t, -1,buscarAbb(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb:= TDADiccionario.CrearABB[string, int](cmpString)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestAbbIteradorInternoValoresConBorrados(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb:= TDADiccionario.CrearABB[string, int](cmpString)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	abb:= TDADiccionario.CrearABB[string, int](cmpString)

	claves := make([]string, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = fmt.Sprintf("%08d", i)
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que devuelva los valores correctos //
	ok := true
	for i := 0; i < n; i++ {
		ok = abb.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = abb.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	// Verifica que borre y devuelva los valores correctos 
	for i := 0; i < n; i++ {
		ok = abb.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !abb.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, abb.Cantidad())
}

func BenchmarkABB(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioVacioABB(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	abb:= TDADiccionario.CrearABB[string, int](cmpString)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestABBIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1,buscarAbb(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1,buscarAbb(segundo, claves))
	require.EqualValues(t, valores[buscarAbb(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1,buscarAbb(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestAbbIteradorNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	claves := []string{"A", "B", "C"}
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	abb.Iterador()
	iter2 := abb.Iterador()
	iter2.Siguiente()
	iter3 := abb.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1,buscarAbb(primero, claves))
	require.NotEqualValues(t, -1,buscarAbb(segundo, claves))
	require.NotEqualValues(t, -1,buscarAbb(tercero, claves))
}

func TestAbbPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	abb:= TDADiccionario.CrearABB[string, string](cmpString)
	abb.Guardar(clave1, "")
	abb.Guardar(clave2, "")
	abb.Guardar(clave3, "")
	abb.Borrar(clave1)
	abb.Borrar(clave2)
	abb.Borrar(clave3)
	iter := abb.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	abb.Guardar(clave1, "A")
	iter = abb.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIteradorABB(b *testing.B, n int) {
	abb:= TDADiccionario.CrearABB[string, *int](cmpString)

	claves := make([]string, n)
	valores := make([]int, n)

	// Inserta 'n' parejas en el hash 
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		abb.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := abb.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkAbbIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestABBVolumenIteradorCorte(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	abb:= TDADiccionario.CrearABB[int, int](cmpInt)

	// Inserta 'n' parejas en el hash //
	for i := 0; i < 10000; i++ {
		abb.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	abb.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

//POSIBLES PRUEBAS CONCRETAS DE ABB

func TestABBMinMax(t *testing.T) {
	t.Log("Verifica que el recorrido InOrder incluya correctamente los elementos mínimo y máximo del ABB")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{5, 3, 8, 1, 4, 7, 9}
	for _, c := range claves {
		abb.Guardar(c, fmt.Sprintf("valor%d", c))
	}

	obtenidoMin := false
	obtenidoMax := false
	abb.Iterar(func(clave int, _ string) bool {
		if clave == 1 {
			obtenidoMin = true
		}
		if clave == 9 {
			obtenidoMax = true
		}
		return true
	})

	require.True(t, obtenidoMin, "El recorrido no incluyó el elemento mínimo")
	require.True(t, obtenidoMax, "El recorrido no incluyó el elemento máximo")
}

func TestABBOInOrder(t *testing.T) {
	t.Log("Verifica que el recorrido interno InOrder respete el orden creciente de claves")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{5, 3, 8, 1, 4, 7, 9}
	for _, c := range claves {
		abb.Guardar(c, fmt.Sprintf("valor%d", c))
	}

	var orden []int
	abb.Iterar(func(clave int, _ string) bool {
		orden = append(orden, clave)
		return true
	})

	require.Equal(t, []int{1, 3, 4, 5, 7, 8, 9}, orden, "El recorrido no es InOrder")
}

func TestABBOInOrderConCorte(t *testing.T) {
	t.Log("Verifica que se pueda cortar la iteración interna antes de recorrer todo el ABB")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{5, 3, 8, 1, 4, 7, 9}
	for _, c := range claves {
		abb.Guardar(c, fmt.Sprintf("valor%d", c))
	}

	var primeros []int
	abb.Iterar(func(clave int, _ string) bool {
		primeros = append(primeros, clave)
		return len(primeros) < 3
	})

	require.Equal(t, []int{1, 3, 4}, primeros, "La iteración no cortó en el momento esperado")
}

func TestABBIteradorExterno(t *testing.T) {
	t.Log("Verifica que el iterador externo recorra todos los elementos en orden")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{5, 3, 8, 1, 4, 7, 9}
	for _, c := range claves {
		abb.Guardar(c, fmt.Sprintf("valor%d", c))
	}

	iter := abb.Iterador()
	var resultado []int
	for iter.HaySiguiente() {
		c, _ := iter.VerActual()
		resultado = append(resultado, c)
		iter.Siguiente()
	}

	require.Equal(t, []int{1, 3, 4, 5, 7, 8, 9}, resultado, "El recorrido del iterador externo no es InOrder")
}

func TestABBIteradorExternoCorte(t *testing.T) {
	t.Log("Verifica que se pueda cortar el recorrido externo antes de iterar todo el ABB")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	claves := []int{5, 3, 8, 1, 4, 7, 9}
	for _, c := range claves {
		abb.Guardar(c, fmt.Sprintf("valor%d", c))
	}

	iter := abb.Iterador()
	var resultado []int
	for iter.HaySiguiente() && len(resultado) < 3 {
		c, _ := iter.VerActual()
		resultado = append(resultado, c)
		iter.Siguiente()
	}

	require.Equal(t, []int{1, 3, 4}, resultado, "El iterador externo no cortó correctamente")
}

func TestABBIterador(t *testing.T) {
	t.Log("Verifica el funcionamiento completo del iterador externo: orden, contenido y panics al final")
	abb := TDADiccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	abb.Guardar(2, "dos")
	abb.Guardar(1, "uno")
	abb.Guardar(3, "tres")

	iter := abb.Iterador()
	require.True(t, iter.HaySiguiente(), "Debe haber al menos un elemento al comenzar")

	c1, v1 := iter.VerActual()
	require.Equal(t, 1, c1, "Primera clave incorrecta")
	require.Equal(t, "uno", v1, "Primer valor incorrecto")
	iter.Siguiente()

	c2, v2 := iter.VerActual()
	require.Equal(t, 2, c2, "Segunda clave incorrecta")
	require.Equal(t, "dos", v2, "Segundo valor incorrecto")
	iter.Siguiente()

	c3, v3 := iter.VerActual()
	require.Equal(t, 3, c3, "Tercera clave incorrecta")
	require.Equal(t, "tres", v3, "Tercer valor incorrecto")
	iter.Siguiente()

	require.False(t, iter.HaySiguiente(), "El iterador debería estar al final")
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}
