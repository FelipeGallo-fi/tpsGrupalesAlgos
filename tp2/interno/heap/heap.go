// Implementación del TDA Heap (cola de prioridad) genérico.
// Se usa como base para la lógica de prioridad_vuelos.

package cola_prioridad

type heap[T any] struct {
	datos       []T
	cantidad    int
	funcion_cmp func(T, T) int
}

const (
	_FACTOR_REDIMENSION_AGRANDAR = 2
	_PILA_VACIA                  = 0
	_MINIMA_CAPACIDAD            = 10
)

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos:       make([]T, _MINIMA_CAPACIDAD),
		cantidad:    0,
		funcion_cmp: funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	capacidad := len(arreglo) * _FACTOR_REDIMENSION_AGRANDAR
	if capacidad < _MINIMA_CAPACIDAD {
		capacidad = _MINIMA_CAPACIDAD
	}
	h := &heap[T]{
		datos:       make([]T, capacidad),
		cantidad:    len(arreglo),
		funcion_cmp: funcion_cmp,
	}
	copy(h.datos, arreglo)
	heapify(h.datos, h.cantidad, funcion_cmp)
	return h
}

func heapify[T any](datos []T, cantidad int, funcion_cmp func(T, T) int) {
	for i := cantidad/2 - 1; i >= 0; i-- {
		downHeapAux(datos, cantidad, i, funcion_cmp)
	}
}

func downHeapAux[T any](datos []T, cantidad int, indice int, funcion_cmp func(T, T) int) {
	hijoIzq := 2*indice + 1
	hijoDer := 2*indice + 2
	mayor := indice

	if hijoIzq < cantidad && funcion_cmp(datos[hijoIzq], datos[mayor]) > 0 {
		mayor = hijoIzq
	}
	if hijoDer < cantidad && funcion_cmp(datos[hijoDer], datos[mayor]) > 0 {
		mayor = hijoDer
	}
	if mayor != indice {
		datos[indice], datos[mayor] = datos[mayor], datos[indice]
		downHeapAux(datos, cantidad, mayor, funcion_cmp)
	}
}

func (h *heap[T]) downHeap(indice int) {
	downHeapAux(h.datos, h.cantidad, indice, h.funcion_cmp)
}

func (h *heap[T]) upHeap(indice int) {
	if indice == 0 {
		return
	}
	padre := (indice - 1) / 2
	if h.funcion_cmp(h.datos[indice], h.datos[padre]) > 0 {
		h.intercambiar(indice, padre)
		h.upHeap(padre)
	}
}

func (h *heap[T]) intercambiar(i, j int) {
	h.datos[i], h.datos[j] = h.datos[j], h.datos[i]
}

func (h *heap[T]) redimensionar(nuevaCapacidad int) {
	nuevos := make([]T, nuevaCapacidad)
	copy(nuevos, h.datos[:h.cantidad])
	h.datos = nuevos
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	heapify(elementos, len(elementos), funcion_cmp)

	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeapAux(elementos, i, 0, funcion_cmp)
	}
}
func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.datos) {
		nuevaCap := len(h.datos) * _FACTOR_REDIMENSION_AGRANDAR
		if nuevaCap == 0 {
			nuevaCap = _MINIMA_CAPACIDAD
		}
		h.redimensionar(nuevaCap)
	}
	h.datos[h.cantidad] = elem
	h.cantidad++
	h.upHeap(h.cantidad - 1)
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	max := h.datos[0]
	h.cantidad--
	if h.cantidad > 0 {
		h.datos[0] = h.datos[h.cantidad]
		h.downHeap(0)
	}

	if h.cantidad > 0 && h.cantidad <= len(h.datos)/4 && len(h.datos)/2 >= _MINIMA_CAPACIDAD {
		h.redimensionar(len(h.datos) / 2)
	}

	return max
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}
