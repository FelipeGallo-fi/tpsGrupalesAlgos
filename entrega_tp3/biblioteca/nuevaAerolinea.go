package biblioteca

import (
	"sort"
	"tp3/tdas/grafo"
	"tp3/tdas/unionFind"
	"tp3/utilidades"
)

func NuevaAerolinea(g grafo.Grafo[string, utilidades.Arista]) []utilidades.AristaCSV {
	gNoDirigido := construirGrafoNoDirigido(g)
	aristas := ObtenerAristas(gNoDirigido)
	sort.Slice(aristas, func(i, j int) bool {
		return aristas[i].Precio < aristas[j].Precio
	})
	return kruskal(gNoDirigido, aristas)
}

func ObtenerAristas(g grafo.Grafo[string, utilidades.Arista]) []utilidades.AristaCSV {
	aristas := []utilidades.AristaCSV{}
	vistos := make(map[string]bool)
	for _, origen := range g.Vertices() {
		vistos[origen] = true
		for _, destino := range g.Adyacentes(origen) {
			if vistos[destino] {
				continue
			}
			a, _ := g.ObtenerArista(origen, destino)
			aristas = append(aristas, utilidades.AristaCSV{
				Origen:     origen,
				Destino:    destino,
				Tiempo:     a.Tiempo,
				Precio:     a.Precio,
				Frecuencia: a.Frecuencia,
			})
		}
	}
	return aristas
}

func construirGrafoNoDirigido(g grafo.Grafo[string, utilidades.Arista]) grafo.Grafo[string, utilidades.Arista] {
	gNoDirigido := grafo.CrearGrafo[string, utilidades.Arista](false)

	for _, v := range g.Vertices() {
		gNoDirigido.AgregarVertice(v)
	}

	vistos := make(map[[2]string]bool)
	for _, v := range g.Vertices() {
		for _, w := range g.Adyacentes(v) {
			par := normalizarPar(v, w)
			if vistos[par] {
				continue
			}
			vistos[par] = true
			a, _ := g.ObtenerArista(v, w)
			gNoDirigido.AgregarArista(v, w, a)
		}
	}

	return gNoDirigido
}

func normalizarPar(a, b string) [2]string {
	if a < b {
		return [2]string{a, b}
	}
	return [2]string{b, a}
}

func kruskal(g grafo.Grafo[string, utilidades.Arista], aristas []utilidades.AristaCSV) []utilidades.AristaCSV {
	uf := unionFind.CrearUnionFind(g.Vertices())
	res := []utilidades.AristaCSV{}
	for _, a := range aristas {
		if uf.Unir(a.Origen, a.Destino) {
			res = append(res, a)
		}
	}
	return res
}
