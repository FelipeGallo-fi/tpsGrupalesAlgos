package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estadoCelda int

const (
	cargaMax        = 0.7
	capacidadMinima = 20
)

const (
	VACIA estadoCelda = iota
	OCUPADA
	BORRADA
)

type DiccionarioHash[K comparable, V any] struct {
	tabla     []hashElem[K, V]
	cantidad  int
	capacidad int
}

type hashElem[K comparable, V any] struct {
	clave  K
	valor  V
	estado estadoCelda
}

type IterDiccionarioImplementacion[K comparable, V any] struct {
	diccionario *DiccionarioHash[K, V]
	posicion    int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &DiccionarioHash[K, V]{}
}

func (d *DiccionarioHash[K, V]) Guardar(clave K, valor V) {
	if d.capacidad == 0 {
		d.capacidad = capacidadMinima
		d.tabla = make([]hashElem[K, V], d.capacidad)
	}

	if float64(d.cantidad+1)/float64(d.capacidad) > cargaMax {
		d.redimensionar(d.capacidad * 2)
	}

	posicion := hash(clave, d.capacidad)

	for {
		elem := &d.tabla[posicion]

		if elem.estado == VACIA || elem.estado == BORRADA {
			break
		}

		if elem.estado == OCUPADA && elem.clave == clave {
			elem.valor = valor
			return
		}

		posicion = (posicion + 1) % d.capacidad
	}

	d.tabla[posicion] = hashElem[K, V]{clave: clave, valor: valor, estado: OCUPADA}
	d.cantidad++
}

func (d *DiccionarioHash[K, V]) Pertenece(clave K) bool {
	if d.capacidad == 0 {
		return false
	}
	start := hash(clave, d.capacidad)
	pos := start

	for {
		elem := d.tabla[pos]
		if elem.estado == VACIA {
			return false
		}
		if elem.estado == OCUPADA && elem.clave == clave {
			return true
		}
		pos = (pos + 1) % d.capacidad
		if pos == start {
			return false
		}
	}
}

func (d *DiccionarioHash[K, V]) Obtener(clave K) V {
	if d.capacidad == 0 {
		panic("La clave no pertenece al diccionario")
	}
	posicion := hash(clave, d.capacidad)
	for {
		elem := d.tabla[posicion]
		if elem.estado == VACIA {
			panic("La clave no pertenece al diccionario")
		}
		if elem.estado == OCUPADA && elem.clave == clave {
			return elem.valor
		}
		posicion = (posicion + 1) % d.capacidad
	}
}

func (d *DiccionarioHash[K, V]) Borrar(clave K) V {
	if d.capacidad == 0 {
		panic("La clave no pertenece al diccionario")
	}

	start := hash(clave, d.capacidad)
	pos := start

	for {
		elem := &d.tabla[pos]
		if elem.estado == VACIA {
			panic("La clave no pertenece al diccionario")
		}
		if elem.estado == OCUPADA && elem.clave == clave {
			valor := elem.valor
			elem.estado = BORRADA
			d.cantidad--

			if d.cantidad == 0 {
				d.tabla = make([]hashElem[K, V], d.capacidad)
			}
			return valor
		}
		pos = (pos + 1) % d.capacidad
		if pos == start {
			panic("La clave no pertenece al diccionario")
		}
	}
}

func (d *DiccionarioHash[K, V]) Cantidad() int {
	return d.cantidad
}

///funcion Iterador Interno

func (d *DiccionarioHash[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, elemento := range d.tabla {
		if elemento.estado == OCUPADA {
			if !f(elemento.clave, elemento.valor) {
				return
			}
		}
	}

}

///funcion Iterador externo

func (d *DiccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {
	posicionOpcupada := 0
	for (posicionOpcupada < len(d.tabla)) && (d.tabla[posicionOpcupada].estado != OCUPADA) {
		posicionOpcupada++
	}

	return &IterDiccionarioImplementacion[K, V]{
		diccionario: d,
		posicion:    posicionOpcupada,
	}
}

///funciones Iterador

func (i *IterDiccionarioImplementacion[K, V]) panicVacia() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (i *IterDiccionarioImplementacion[K, V]) HaySiguiente() bool {

	for i.posicion < len(i.diccionario.tabla) && i.diccionario.tabla[i.posicion].estado != OCUPADA {
		i.posicion++
	}
	return i.posicion < len(i.diccionario.tabla) && i.diccionario.tabla[i.posicion].estado == OCUPADA

}

func (i *IterDiccionarioImplementacion[K, V]) VerActual() (K, V) {
	i.panicVacia()
	return i.diccionario.tabla[i.posicion].clave, i.diccionario.tabla[i.posicion].valor
}

func (i *IterDiccionarioImplementacion[K, V]) Siguiente() {
	i.panicVacia()
	i.posicion++

}

func hash[K comparable](clave K, capacidad int) int {
	h := fnv.New32a()
	switch v := any(clave).(type) {
	case string:
		h.Write([]byte(v))
	case []byte:
		h.Write(v)
	case int:
		h.Write([]byte(fmt.Sprintf("%d", v)))
	default:
		h.Write([]byte(fmt.Sprint(v)))
	}
	return int(h.Sum32() % uint32(capacidad))
}

func (d *DiccionarioHash[K, V]) redimensionar(nuevaCapacidad int) {
	viejaTabla := d.tabla

	d.tabla = make([]hashElem[K, V], nuevaCapacidad)
	d.capacidad = nuevaCapacidad
	d.cantidad = 0

	for _, elem := range viejaTabla {
		if elem.estado == OCUPADA {
			d.Guardar(elem.clave, elem.valor)
		}
	}
}
