package unionFind

type UnionFind[T comparable] struct {
	padre map[T]T
}

func CrearUnionFind[T comparable](vertices []T) *UnionFind[T] {
	uf := &UnionFind[T]{padre: make(map[T]T)}
	for _, v := range vertices {
		uf.padre[v] = v
	}
	return uf
}

func (uf *UnionFind[T]) Encontrar(v T) T {
	if uf.padre[v] != v {
		uf.padre[v] = uf.Encontrar(uf.padre[v])
	}
	return uf.padre[v]
}

func (uf *UnionFind[T]) Unir(v, w T) bool {
	raizV := uf.Encontrar(v)
	raizW := uf.Encontrar(w)
	if raizV == raizW {
		return false
	}
	uf.padre[raizW] = raizV
	return true
}
