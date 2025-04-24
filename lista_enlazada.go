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

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for actual := l.primero; actual != nil; {
		if !visitar(actual.dato) {
			break
		}
		actual = actual.siguiente
	}
}

type iteradorLista[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorLista[T]{
		lista:    l,
		actual:   l.primero,
		anterior: nil,
	}
}

func (iter *iteradorLista[T]) VerActual() T {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iteradorLista[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iteradorLista[T]) Siguiente() {
	if iter.actual == nil {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter *iteradorLista[T]) Insertar(dato T) {
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

func (iter *iteradorLista[T]) Borrar() T {
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
