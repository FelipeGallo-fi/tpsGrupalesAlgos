package lista

type nodoLista[T any] struct{
	dato	 T
	siguiente *nodoLista[T]
}

type ListaEnlazada[T any] struct{
	
	primero *nodoLista[T]
	ultimo *nodoLista[T]
	largo int
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

func (l *ListaEnlazada[T]) InsertarPrimero(dato T){
	nuevo := nuevoNodo(dato)
	nuevo.siguiente = l.primero
	l.primero =nuevo
	if l.ultimo == nil{
		l.ultimo = nuevo
	} 
	l.largo++
}

func (l *ListaEnlazada[T]) InsertarUltimo(dato T){
	nuevo := nuevoNodo(dato)
	if l.ultimo != nil {
		l.ultimo.siguiente = nuevo
	} else {
		l.primero = nuevo
	}
	l.ultimo = nuevo
	l.largo++

}

func (l *ListaEnlazada[T])panicVacia() {
	if l.EstaVacia(){
		panic("La lista esta vacia =(")
	}
}
func (l *ListaEnlazada[T])BorrarPrimero() T {
	l.panicVacia()
	dato := l.primero.dato
	if l.largo >1 {
		l.primero = l.primero.siguiente
	} else {
		l.primero = nil
		l.ultimo = nil
	}
	l.largo --
	return dato
}

func (l *ListaEnlazada[T])VerPrimero() T{
	l.panicVacia()
	return l.primero.dato
}
func (l *ListaEnlazada[T])VerUltimo() T{
	l.panicVacia()
	return l.ultimo.dato
}


func (l *ListaEnlazada[T]) Largo() (dato int){
	return l.largo
}