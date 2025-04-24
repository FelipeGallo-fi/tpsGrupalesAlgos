package lista_test

import (
	"testing"
)

const _VOLUMEN = 1000000

func TestListaVacia(t *testing.T){
	/*que la lista vacia se comporte como una lista vacia
	ej:
		panic al verprimero , borrar ,etc

	*/
}

func TestListaVolumen(t *testing.T){
	/*para provar de cargar muchos datos
	ej:
		probar que al ir insertando datos del 1 al 1000
		el l.primero.dato se vaya actualizando
		*/
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




