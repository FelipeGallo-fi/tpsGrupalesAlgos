package diccionario

type DiccionarioOrdenado[K comparable, V any] interface {
	Diccionario[K, V]

	// IterarRango itera sólo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
	// incluyéndolos en caso de encontrarse
	IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool)

	// IteradorRango crea un IterDiccionario que sólo itere por las claves que se encuentren en el rango indicado
	IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
}


type ABB[K comparable, V any] interface {
    Guardar(clave K, dato V)
    Pertenece(clave K) bool
    Obtener(clave K) V
    Borrar(clave K) V
    Cantidad() int
    Iterar(func(clave K, dato V) bool)
    Iterador() IterDiccionario[K, V]
}
