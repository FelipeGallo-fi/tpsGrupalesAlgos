package grafo

type grafoImpl[V comparable, T any] struct {
	adyacentes map[V]map[V]T
	dirigido   bool
}

func CrearGrafo[V comparable, T any](dirigido bool) Grafo[V, T] {
	return &grafoImpl[V, T]{
		adyacentes: make(map[V]map[V]T),
		dirigido:   dirigido,
	}
}

func (g *grafoImpl[V, T]) AgregarVertice(v V) {
	if _, ok := g.adyacentes[v]; !ok {
		g.adyacentes[v] = make(map[V]T)
	}
}

func (g *grafoImpl[V, T]) EliminarVertice(v V) {
	delete(g.adyacentes, v)
	for origen := range g.adyacentes {
		delete(g.adyacentes[origen], v)
	}
}

func (g *grafoImpl[V, T]) AgregarArista(origen, destino V, datos T) {
	g.AgregarVertice(origen)
	g.AgregarVertice(destino)
	g.adyacentes[origen][destino] = datos
	if !g.dirigido {
		g.adyacentes[destino][origen] = datos
	}
}

func (g *grafoImpl[V, T]) EliminarArista(origen, destino V) {
	delete(g.adyacentes[origen], destino)
	if !g.dirigido {
		delete(g.adyacentes[destino], origen)
	}
}

func (g *grafoImpl[V, T]) ExisteVertice(v V) bool {
	_, ok := g.adyacentes[v]
	return ok
}

func (g *grafoImpl[V, T]) ExisteArista(origen, destino V) bool {
	_, ok := g.adyacentes[origen][destino]
	return ok
}

func (g *grafoImpl[V, T]) ObtenerArista(origen, destino V) (T, bool) {
	val, ok := g.adyacentes[origen][destino]
	if ok {
		return val, true
	}
	var cero T
	return cero, false
}

func (g *grafoImpl[V, T]) Adyacentes(v V) []V {
	var ady []V
	for destino := range g.adyacentes[v] {
		ady = append(ady, destino)
	}
	return ady
}

func (g *grafoImpl[V, T]) Vertices() []V {
	vs := make([]V, 0, len(g.adyacentes))
	for v := range g.adyacentes {
		vs = append(vs, v)
	}
	return vs
}
