package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type iteradorListaImplementacion[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l *listaEnlazada[T]) EstaVacia() bool {
	return l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := crearNodo(dato)
	if l.EstaVacia() {
		l.ultimo = nuevo
	}
	nuevo.siguiente = l.primero
	l.primero = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo := crearNodo(dato)
	if !l.EstaVacia() {
		l.ultimo.siguiente = nuevo
	} else {
		l.primero = nuevo
	}
	l.ultimo = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
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

func (l *listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.primero.dato
}

func (l *listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ultimo.dato
}

func (l *listaEnlazada[T]) Largo() (dato int) {
	return l.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := lista.primero; actual != nil; {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

func crearNodo[T any](dato T) *nodoLista[T] {
	return &nodoLista[T]{dato: dato}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaImplementacion[T]{
		lista:    lista,
		actual:   lista.primero,
		anterior: nil,
	}
}

func (iter *iteradorListaImplementacion[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorListaImplementacion[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorListaImplementacion[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorListaImplementacion[T]) Insertar(dato T) {
	nuevo := crearNodo(dato)
	nuevo.siguiente = iter.actual

	if iter.anterior == nil {
		iter.lista.primero = nuevo
		if iter.lista.ultimo == nil {
			iter.lista.ultimo = nuevo
		}
	} else {
		iter.anterior.siguiente = nuevo
	}

	if iter.actual == nil {
		iter.lista.ultimo = nuevo
	}

	iter.actual = nuevo
	iter.lista.largo++
}

func (iter *iteradorListaImplementacion[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	dato := iter.actual.dato

	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
		iter.actual = iter.lista.primero
		if iter.actual == nil {
			iter.lista.ultimo = nil
		}
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
		iter.actual = iter.actual.siguiente
		if iter.actual == nil {
			iter.lista.ultimo = iter.anterior
		}
	}

	iter.lista.largo--
	return dato
}
