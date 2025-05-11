package lista

type Lista[T any] interface {
	//Comprueba el largo de la lista
	//si es  0 esta vacia retorna un true , sino  false
	EstaVacia() bool
	//Guarda el elemento  , y comprueba si esta vacia la lista, si esta vacia inserta al final
	//Si la lista tiene algun elemento  este va a apuntar a el que era el primero , y el primero va a ser el nuevo
	InsertarPrimero(T)
	//Guarda el elemento  , si la lista esta vacia lo inserta al final , sino inserta en primer lugar y este va a apuntar al segundo
	InsertarUltimo(T)
	//Comprueba que la lista no esta vacia , sino entre en panic
	//Guarda el dato , y cambia la direccion del primero al siguiente
	BorrarPrimero() T
	//Comprueba que la lista no esta vacia ,sino entra en panic con mensaje: "La lista esta vacia" , si no esta vaica deuvelve el primer dato
	VerPrimero() T
	//Comprueba que la lista no esta vacia ,sino entra en panic con mensaje: "La lista esta vacia , si no esta vacia devuelve el ultimo dato
	VerUltimo() T
	//Devuelve un dato de tipo int  con la cantidad de elementos que tiene la lista
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
