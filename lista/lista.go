package lista

type Lista[T any] interface {
	EstaVacia() bool
	InsertarPrimero(T)
	InsertarUltimo(T)
	BorrarPrimero() T
	VerPrimero() T
	VerUltimo() T
	Largo() int

	// Iterar recorre la lista desde el primer elemento hasta el último,
	// ejecutando la función visitar con cada uno de los elementos.
	// La iteración se detiene si visitar devuelve false.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un iterador externo que permite recorrer y modificar
	// la lista desde el primer elemento hasta el último.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// VerActual devuelve el elemento en la posición actual del iterador.
	// Si el iterador ya terminó de recorrer la lista, entra en pánico con el mensaje "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente indica si hay un elemento válido en la posición actual del iterador.
	// Devuelve true si el iterador aún no terminó de recorrer la lista, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza el iterador a la siguiente posición.
	// Si el iterador ya terminó de iterar, entra en pánico con el mensaje "El iterador termino de iterar".
	Siguiente()

	// Insertar inserta un nuevo elemento en la posición actual del iterador.
	// Si el iterador está al principio, el elemento se inserta al inicio de la lista.
	// Si el iterador ya terminó de iterar, el elemento se inserta al final.
	// El iterador queda apuntando al nuevo elemento insertado.
	Insertar(T)

	// Borrar elimina el elemento en la posición actual del iterador y devuelve su valor.
	// Si el iterador ya terminó de iterar, entra en pánico con el mensaje "El iterador termino de iterar".
	// Luego de borrar, el iterador queda apuntando al siguiente elemento.
	Borrar() T
}
