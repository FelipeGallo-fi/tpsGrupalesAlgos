package diccionario

import (
	"fmt"
	"hash/fnv"
)

type estadoCelda int

const (
	_CARGA_MAX        = 0.7
	_CAPACIDAD_MINIMA = 20
	_REDIMENSION      = 2
	_REDUCCION        = 0.25
)

const (
	_VACIA estadoCelda = iota
	_OCUPADA
	_BORRADA
)

type diccionarioHash[K comparable, V any] struct {
	tabla     []hashElem[K, V]
	cantidad  int
	borrados  int
	capacidad int
}

type hashElem[K comparable, V any] struct {
	clave  K
	valor  V
	estado estadoCelda
}

type iterDiccionarioImplementacion[K comparable, V any] struct {
	diccionario *diccionarioHash[K, V]
	posicion    int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	d := &diccionarioHash[K, V]{}
	d.inicializarTabla(_CAPACIDAD_MINIMA)
	return d
}

func (d *diccionarioHash[K, V]) crearTablaVacia(capacidad int) {
	if capacidad < _CAPACIDAD_MINIMA {
		capacidad = _CAPACIDAD_MINIMA
	}
	d.capacidad = capacidad
	d.tabla = make([]hashElem[K, V], d.capacidad)
	d.cantidad = 0
	d.borrados = 0
}

func (d *diccionarioHash[K, V]) inicializarTabla(capacidad int) {
	if capacidad < _CAPACIDAD_MINIMA {
		capacidad = _CAPACIDAD_MINIMA
	}
	d.crearTablaVacia(capacidad)
}

func (d *diccionarioHash[K, V]) buscarPos(clave K) (pos int, estado estadoCelda) {
	pos = hash(clave, d.capacidad)
	primerBorrado := -1

	for {
		elem := &d.tabla[pos]

		if elem.estado == _VACIA {
			if primerBorrado != -1 {
				return primerBorrado, _BORRADA
			}
			return pos, _VACIA
		}

		if elem.estado == _BORRADA {
			if primerBorrado == -1 {
				primerBorrado = pos
			}
		} else if elem.estado == _OCUPADA && elem.clave == clave {
			return pos, _OCUPADA
		}

		pos = (pos + 1) % d.capacidad
	}
}
func (d *diccionarioHash[K, V]) Guardar(clave K, valor V) {
	if float64(d.cantidad+d.borrados+1)/float64(d.capacidad) > _CARGA_MAX {
		d.redimensionar(d.capacidad * _REDIMENSION)
	}

	posicion, estado := d.buscarPos(clave)
	if estado == _OCUPADA {
		d.tabla[posicion].valor = valor
		return
	}

	d.tabla[posicion] = hashElem[K, V]{clave: clave, valor: valor, estado: _OCUPADA}
	d.cantidad++
}

func (d *diccionarioHash[K, V]) Pertenece(clave K) bool {
	_, estado := d.buscarPos(clave)
	return estado == _OCUPADA
}

func (d *diccionarioHash[K, V]) Obtener(clave K) V {
	posicion, estado := d.buscarPos(clave)
	if estado != _OCUPADA {
		panic("La clave no pertenece al diccionario")
	}
	return d.tabla[posicion].valor

}

func (d *diccionarioHash[K, V]) Borrar(clave K) V {
	posicion, estado := d.buscarPos(clave)

	if estado != _OCUPADA {
		panic("La clave no pertenece al diccionario")
	}

	valor := d.tabla[posicion].valor
	d.tabla[posicion].estado = _BORRADA
	d.cantidad--
	d.borrados++

	if d.cantidad == 0 {
		d.inicializarTabla(_CAPACIDAD_MINIMA)
		d.borrados = 0
	} else if float64(d.cantidad+d.borrados)/float64(d.capacidad) < _REDUCCION && d.capacidad > _CAPACIDAD_MINIMA {
		nuevaCapacidad := max(d.capacidad/_REDIMENSION, _CAPACIDAD_MINIMA)
		d.redimensionar(nuevaCapacidad)
	}

	return valor
}

func (d *diccionarioHash[K, V]) Cantidad() int {
	return d.cantidad
}

///funcion Iterador Interno

func (d *diccionarioHash[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, elemento := range d.tabla {
		if elemento.estado == _OCUPADA {
			if !f(elemento.clave, elemento.valor) {
				return
			}
		}
	}

}

///funcion Iterador externo

func (d *diccionarioHash[K, V]) proximaPosicionOcupada(pos int) int {
	for pos < len(d.tabla) && d.tabla[pos].estado != _OCUPADA {
		pos++
	}
	return pos
}

func (d *diccionarioHash[K, V]) Iterador() IterDiccionario[K, V] {
	posicionInicial := d.proximaPosicionOcupada(0)
	return &iterDiccionarioImplementacion[K, V]{
		diccionario: d,
		posicion:    posicionInicial,
	}
}

///funciones Iterador

func (i *iterDiccionarioImplementacion[K, V]) panicVacia() {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}

func (i *iterDiccionarioImplementacion[K, V]) HaySiguiente() bool {
	proximo := i.diccionario.proximaPosicionOcupada(i.posicion)
	return proximo < len(i.diccionario.tabla)
}

func (i *iterDiccionarioImplementacion[K, V]) VerActual() (K, V) {
	i.panicVacia()
	return i.diccionario.tabla[i.posicion].clave, i.diccionario.tabla[i.posicion].valor
}

func (i *iterDiccionarioImplementacion[K, V]) Siguiente() {
	i.panicVacia()
	i.posicion = i.diccionario.proximaPosicionOcupada(i.posicion + 1)
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

func (d *diccionarioHash[K, V]) redimensionar(nuevaCapacidad int) {
	viejaTabla := d.tabla

	d.crearTablaVacia(nuevaCapacidad)

	for _, elem := range viejaTabla {
		if elem.estado == _OCUPADA {
			d.Guardar(elem.clave, elem.valor)
		}
	}

	d.borrados = 0
}
