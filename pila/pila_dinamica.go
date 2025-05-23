package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

const _REDIMENSIONARRIBA = 2
const _REDIMENSIONABAJO = 4
const _ARRAYININICIAL = 10

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, _ARRAYININICIAL), cantidad: 0}
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) vacia() {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
}

func (p *pilaDinamica[T]) VerTope() (dato T) {

	p.vacia()
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) redimensionar(redimension int) {
	pilaRedimensionada := make([]T, redimension)
	copy(pilaRedimensionada, p.datos[:p.cantidad])
	p.datos = pilaRedimensionada
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == cap(p.datos) {
		capacidadPila := (cap(p.datos) * _REDIMENSIONARRIBA)
		p.redimensionar(capacidadPila)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++

}
func (p *pilaDinamica[T]) Desapilar() T {

	p.vacia()
	if p.cantidad*_REDIMENSIONABAJO <= cap(p.datos) && (cap(p.datos) > _ARRAYININICIAL) {
		capacidadPila := cap(p.datos) / 2
		p.redimensionar(capacidadPila)
	}
	p.cantidad--
	dato := p.datos[p.cantidad]

	return dato

}
