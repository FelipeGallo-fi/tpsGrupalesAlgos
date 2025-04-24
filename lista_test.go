package lista_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	TDALista "tdas/lista"
)

const _VOLUMEN = 1000000

func TestListaVacia(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()

	require.True(t, lista.EstaVacia())

	require.Panics(t, func() { lista.BorrarPrimero() }, "BorrarPrimero en un lista recien  deberia panickear")
	require.Panics(t, func() { lista.VerPrimero() }, "VerPimero en una lista recien creada deberia panickear")
}

func TestListaVolumen(t *testing.T){
	lista := TDALista.CrearListaEnlazada[int]()
	
	for i:=0;i<_VOLUMEN;i++{
		if i%2 == 0 {
			lista.InsertarPrimero(i)
		require.Equal(t,i,lista.VerPrimero(),"El primer elemento deberia cambiar")
		} else {
			lista.InsertarUltimo(i)
			require.Equal(t,i,lista.VerPrimero(),"El Ultimo elemento deberia cambiar")
		}
	}
	require.Equal(t,_VOLUMEN,lista.Largo(),"El largo de  mi lista deberia ser %i",_VOLUMEN)

	for !lista.EstaVacia() {
		lista.BorrarPrimero()
	}

	require.True(t, lista.EstaVacia(),"La lista deberia esta vacia despues de borrarla por completo")
	require.Panics(t, func() {lista.BorrarPrimero()},"Borrar en una lista vacia deberia panickear")

}


func TestListas(t *testing.T){
	/*Este es un test general que pudemos usar para hacer una lista 
	de varios tipos 
	ej: 
		int
		float64
		string	
		bool
	*/
}
func TestListaListas(t *testing.T){
	//Una lista de listas
}

func TestInsertar(t *testing.T){
	/*podemos usarlo para insertar muchos datos intercalando en 
	ej:	
		insertar primero y ver ultimo muchas veces
		ver ultimo muchas veces
		etc
	*/
}

func TestCrearIterador(t *testing.T) {
	/*Al insertar un elemento en la posición en la que se crea el iterador,
	 efectivamente se inserta al principio.*/
}

func TestIteradorAlFinal(t *testing.T) {
	/*Insertar un elemento cuando el iterador está al final
	 efectivamente es equivalente a insertar al final.*/
}

func TestInsertarMedio(t *testing.T) {
	/*Insertar un elemento en el medio se hace en la posición correcta..*/
}
func TestRemoverMedio(t *testing.T) {
	/*Verificar que al remover un elemento del medio, este no está.*/
}

func TestRemoverElmento(t *testing.T) {
	/*Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.*/
}

func TestRemoverUltimoIteradorElmento(t *testing.T) {
	/*Remover el último elemento con el iterador cambia el último de la lista..*/
}


func TestIterarListaVacia(t *testing.T) {
	//iterar lista vacia
}

func TestIteradorInterno(t *testing.T) {
	//probar que se pueda reccorer la lisa con VerActual() y ver siguiente 
}



