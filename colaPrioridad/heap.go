package cola_prioridad

type heap[T any] struct {
	datos       []T
	cantidad    int
	funcion_cmp func(T, T) int
}

const (
	FACTOR_REDIMENSION_AGRANDAR = 2
	PILA_VACIA                  = 0
	MINIMA_CAPACIDAD            = 10
)

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	return &heap[T]{
		datos:       make([]T, MINIMA_CAPACIDAD),
		cantidad:    0,
		funcion_cmp: funcion_cmp,
	}
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) *heap[T] {
	capacidad := len(arreglo) * FACTOR_REDIMENSION_AGRANDAR
	if capacidad < MINIMA_CAPACIDAD {
		capacidad = MINIMA_CAPACIDAD
	}
	h := &heap[T]{
		datos:       make([]T, capacidad),
		cantidad:    len(arreglo),
		funcion_cmp: funcion_cmp,
	}
	copy(h.datos, arreglo)
	for i := len(arreglo)/2 - 1; i >= 0; i-- {
		h.downHeap(i)
	}
	return h
}

func (h *heap[T]) downHeap(indice int) {
	hijoIzq := 2*indice + 1
	hijoDer := 2*indice + 2
	mayor := indice

	if hijoIzq < h.cantidad && h.funcion_cmp(h.datos[hijoIzq], h.datos[mayor]) > 0 {
		mayor = hijoIzq
	}
	if hijoDer < h.cantidad && h.funcion_cmp(h.datos[hijoDer], h.datos[mayor]) > 0 {
		mayor = hijoDer
	}
	if mayor != indice {
		h.intercambiar(indice, mayor)
		h.downHeap(mayor)
	}
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

func (h *heap[T]) ordenar() {
	original := h.cantidad
	for i := h.cantidad/2 - 1; i >= 0; i-- {
		h.downHeap(i)
	}
	for i := h.cantidad - 1; i > 0; i-- {
		h.intercambiar(0, i)
		h.cantidad--
		h.downHeap(0)
	}
	h.cantidad = original
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	copia := make([]T, len(elementos))
	copy(copia, elementos)
	h := CrearHeapArr(copia, funcion_cmp)
	h.ordenar()
	copy(elementos, h.datos[:h.cantidad])
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.datos) {
		nuevaCap := len(h.datos) * FACTOR_REDIMENSION_AGRANDAR
		if nuevaCap == 0 {
			nuevaCap = MINIMA_CAPACIDAD
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
	return max
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}
