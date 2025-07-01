package cola

type colaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
}

type nodo[T any] struct {
	dato      T
	siguiente *nodo[T]
}

func CrearColaEnlazada[T any]() Cola[T] {
	return &colaEnlazada[T]{}
}

func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato}
}

func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil
}

func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato
}

func (cola *colaEnlazada[T]) Encolar(elemento T) {
	nuevoNodo := crearNodo(elemento)
	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.siguiente = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}

	elemento := cola.primero.dato
	cola.primero = cola.primero.siguiente

	//Queda vacia
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return elemento
}
