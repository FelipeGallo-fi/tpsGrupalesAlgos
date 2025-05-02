package diccionario

import (
	"bytes"
	"encoding/binary"
	"reflect"
)


type estadoCelda int 

const (
	VACIA estadoCelda = iota
	OCUPADA
	BORRADA
)

type DiccionarioHash[K comparable, V any] struct {
    tabla []hashElem[K,V]
    cantidad int      
    capacidad int   
}

type hashElem[K comparable, V any] struct{
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

func (d *DiccionarioHash[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _,elemento := range d.tabla{
		if elemento.estado == OCUPADA{
			if !f(elemento.clave ,elemento.valor){
				return
			}
		}
	}
	
}

///funcion Iterador externo

func (d *DiccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {
	posicionOpcupada :=0 
	for (posicionOpcupada < len(d.tabla)) && (d.tabla[posicionOpcupada].estado != OCUPADA){
		posicionOpcupada ++
	}
	
	return &IterDiccionarioImplementacion[K , V]{
		diccionario: d,
		posicion: posicionOpcupada,
	} 
}



///funciones Iterador

func (i *IterDiccionarioImplementacion[K, V]) PanicVacia()  {
	if !i.HaySiguiente(){
		panic("El iterador termino de iterar")
	}
}


func (i *IterDiccionarioImplementacion[K, V]) HaySiguiente() bool {

	for i.posicion < len(i.diccionario.tabla) && i.diccionario.tabla[i.posicion].estado != OCUPADA {
		i.posicion++  
	}
	return i.posicion < len(i.diccionario.tabla) && i.diccionario.tabla[i.posicion].estado  == OCUPADA
	
}

func (i *IterDiccionarioImplementacion[K, V]) VerActual() (K, V) {
	i.PanicVacia()
	return i.diccionario.tabla[i.posicion].clave ,i.diccionario.tabla[i.posicion].valor
}


func (i *IterDiccionarioImplementacion[K, V]) Siguiente() {
	i.PanicVacia()
	i.posicion++

}


//Funcion de hasing sacada de ChatGPT

func Hash(data interface{}) uint64 {
	const (
		prime64   = 1099511628211
		offset64  = 14695981039346656037
		mixPrime1 = 0x100000001b3
		mixPrime2 = 0xC6A4A7935BD1E995
	)

	buf := new(bytes.Buffer)
	writeToBuffer(buf, data)

	var hash uint64 = offset64
	for _, b := range buf.Bytes() {
		hash ^= uint64(b)
		hash *= prime64
	}

	
	hash ^= hash >> 33
	hash *= mixPrime2
	hash ^= hash >> 29
	hash *= mixPrime1
	hash ^= hash >> 32

	return hash
}


func writeToBuffer(buf *bytes.Buffer, val interface{}) {
	v := reflect.ValueOf(val)

	switch v.Kind() {
	case reflect.String:
		buf.WriteString(v.String())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		binary.Write(buf, binary.LittleEndian, v.Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		binary.Write(buf, binary.LittleEndian, v.Uint())
	case reflect.Float32, reflect.Float64:
		binary.Write(buf, binary.LittleEndian, v.Float())
	case reflect.Bool:
		var b byte = 0
		if v.Bool() {
			b = 1
		}
		buf.WriteByte(b)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			writeToBuffer(buf, v.Index(i).Interface())
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			writeToBuffer(buf, v.Field(i).Interface())
		}
	case reflect.Ptr, reflect.Interface:
		if !v.IsNil() {
			writeToBuffer(buf, v.Elem().Interface())
		}
	default:
		
	}
}





