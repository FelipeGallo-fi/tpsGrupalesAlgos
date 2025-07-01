package grafo

// Grafo representa un grafo dirigido y ponderado.
// V es el tipo del vértice (debe ser comparable para poder indexar mapas).
// T es el tipo de los datos asociados a las aristas.
type Grafo[V comparable, T any] interface {

	// AgregarVertice agrega un vértice al grafo. Si ya existía, no realiza ninguna acción.
	AgregarVertice(v V)

	// EliminarVertice elimina un vértice y todas sus aristas asociadas del grafo.
	EliminarVertice(v V)

	// AgregarArista agrega una arista dirigida desde el vértice origen al destino, con los datos proporcionados.
	// Si ya existía la arista, se actualizan sus datos. Agrega los vértices si no existen.
	AgregarArista(origen, destino V, datos T)

	// EliminarArista elimina la arista dirigida desde el vértice origen al destino.
	EliminarArista(origen, destino V)

	// ExisteVertice indica si el vértice dado se encuentra en el grafo.
	ExisteVertice(v V) bool

	// ExisteArista indica si existe una arista dirigida desde el vértice origen al destino.
	ExisteArista(origen, destino V) bool

	// ObtenerArista devuelve los datos de la arista dirigida desde el vértice origen al destino.
	// Si la arista no existe, devuelve false como segundo valor.
	ObtenerArista(origen, destino V) (T, bool)

	// Adyacentes devuelve una lista de los vértices adyacentes al vértice dado.
	Adyacentes(v V) []V

	// Vertices devuelve una lista con todos los vértices del grafo.
	Vertices() []V
}
