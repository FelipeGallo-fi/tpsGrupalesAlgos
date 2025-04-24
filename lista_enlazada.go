package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type ListaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func (lista *ListaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

type IteradorListaImplementacion[T any] struct {
	lista    *ListaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (lista *ListaEnlazada[T]) Iterador() IteradorLista[T] {
	return &IteradorListaImplementacion[T]{
		lista:    lista,
		actual:   lista.primero,
		anterior: nil,
	}
}

func (iter *IteradorListaImplementacion[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *IteradorListaImplementacion[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *IteradorListaImplementacion[T]) Siguiente() {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *IteradorListaImplementacion[T]) Insertar(dato T) {
	if iter.anterior == nil {
		iter.lista.InsertarPrimero(dato)
		iter.actual = iter.lista.primero
	} else if iter.actual == nil {
		iter.lista.InsertarUltimo(dato)
		iter.actual = iter.lista.ultimo
	} else {
		nuevo := &nodoLista[T]{dato, iter.actual}
		iter.anterior.siguiente = nuevo
		iter.actual = nuevo
		iter.lista.largo++
	}
}

func (iter *IteradorListaImplementacion[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	var dato T

	if iter.anterior == nil {
		dato = iter.lista.BorrarPrimero()
		iter.actual = iter.lista.primero
	} else {
		dato = iter.actual.dato
		iter.anterior.siguiente = iter.actual.siguiente
		if iter.actual == iter.lista.ultimo {
			iter.lista.ultimo = iter.anterior
		}
		iter.actual = iter.actual.siguiente
		iter.lista.largo--
	}

	return dato
}
func CrearListaEnlazada[T any]() Lista[T] {
	return &ListaEnlazada[T]{}
}

func (l *ListaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func nuevoNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func (l *ListaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := nuevoNodo(dato)
	nuevo.siguiente = l.primero
	l.primero = nuevo
	if l.ultimo == nil {
		l.ultimo = nuevo
	}
	l.largo++
}

func (l *ListaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo := nuevoNodo(dato)
	if l.ultimo != nil {
		l.ultimo.siguiente = nuevo
	} else {
		l.primero = nuevo
	}
	l.ultimo = nuevo
	l.largo++

}

func (l *ListaEnlazada[T]) panicVacia() {
	if l.EstaVacia() {
		panic("La lista esta vacia =(")
	}
}
func (l *ListaEnlazada[T]) BorrarPrimero() T {
	l.panicVacia()
	dato := l.primero.dato
	if l.largo > 1 {
		l.primero = l.primero.siguiente
	} else {
		l.primero = nil
		l.ultimo = nil
	}
	l.largo--
	return dato
}

func (l *ListaEnlazada[T]) VerPrimero() T {
	l.panicVacia()
	return l.primero.dato
}
func (l *ListaEnlazada[T]) VerUltimo() T {
	l.panicVacia()
	return l.ultimo.dato
}

func (l *ListaEnlazada[T]) Largo() (dato int) {
	return l.largo
}
