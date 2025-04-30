package diccionario

import "fmt"


type estadoCelda int 

const (
	VACIA estadoCelda = iota
	OCUPADA
	BORRADA
)

type DiccionarioHash[K comparable, V any] struct {
    tabla []entrada[K,V]
    cantidad int      
    capacidad int   
}

type entrada[K comparable, V any] struct{
	clave K
	valor V
	estado estadoCelda
}

type IterDiccionarioImplementacion[K comparable, V any] struct{
	diccionario *DiccionarioHash[K, V]
	posicion 	int 
}

func CrearHash[K comparable, V any]() Diccionario[K, V]{
	return &DiccionarioHash[K,V]{}
}

func (d *DiccionarioHash[K, V]) Guardar(K, V){

}

func (d *DiccionarioHash[K, V]) Pertenece(clave K) bool{

}

func (d *DiccionarioHash[K, V]) Obtener(clave K) V{

}


func (d *DiccionarioHash[K, V]) Borrar(clave K) V{

}

func (d *DiccionarioHash[K, V]) Cantidad() int{

}

///funcion Iterador Interno

func (d *DiccionarioHash[K, V]) Iterar(func(clave K, dato V) bool) {
	
}

///funcion Iterador externo

func (d *DiccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {

}

///funciones Iterador
func (i *DiccionarioHash[K, V]) VerActual() (K, V) {
	
}


func (i *DiccionarioHash[K, V]) HaySiguiente() bool {
	
}

func (i *DiccionarioHash[K, V]) Siguiente() {
	
}


//funcion de hasing generica dada por la catedra , HAY QUE CAMBIARLA
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}







