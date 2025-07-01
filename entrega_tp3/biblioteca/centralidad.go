package biblioteca

import (
	"math"
	"sort"

	"tp3/tdas/grafo"
	TDAHEAP "tp3/tdas/heap"
	"tp3/utilidades"
)

type resultadoCentralidad struct {
	Aeropuerto string
	Valor      int
}

type nodo struct {
	vertice string
	peso    float64
}

func Centralidad(g grafo.Grafo[string, utilidades.Arista]) map[string]int {
	centralidades := make(map[string]int)
	for _, v := range g.Vertices() {
		centralidades[v] = 0
	}

	for _, origen := range g.Vertices() {
		distancias, padres := CaminoMinimoDijkstraMultipadre(g, origen)

		conteo := make(map[string]float64)
		for _, v := range g.Vertices() {
			conteo[v] = 0
		}

		ordenados := make([]string, 0, len(distancias))
		for v := range distancias {
			ordenados = append(ordenados, v)
		}
		sort.Slice(ordenados, func(i, j int) bool {
			return distancias[ordenados[i]] > distancias[ordenados[j]]
		})

		for _, w := range ordenados {
			for _, p := range padres[w] {
				conteo[p] += 1 + conteo[w]
			}
		}

		for _, w := range g.Vertices() {
			if w == origen {
				continue
			}
			centralidades[w] += int(conteo[w])
		}
	}

	return centralidades
}

func TopN(centr map[string]int, n int) []string {
	res := make([]resultadoCentralidad, 0, len(centr))
	for k, v := range centr {
		res = append(res, resultadoCentralidad{k, v})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Valor > res[j].Valor
	})

	tope := min(len(res), n)

	salida := make([]string, 0, tope)
	for i := range tope {
		salida = append(salida, res[i].Aeropuerto)
	}
	return salida
}

func cmpNodo(a, b nodo) int {
	if a.peso < b.peso {
		return -1
	} else if a.peso > b.peso {
		return 1
	}
	return 0
}

func CaminoMinimoDijkstraMultipadre(g grafo.Grafo[string, utilidades.Arista], origen string) (map[string]float64, map[string][]string) {

	distancia := make(map[string]float64)
	padres := make(map[string][]string)
	visitado := make(map[string]bool)

	for _, v := range g.Vertices() {
		distancia[v] = math.Inf(1)
		padres[v] = []string{}
	}
	distancia[origen] = 0

	h := TDAHEAP.CrearHeap(cmpNodo)
	h.Encolar(nodo{origen, 0})

	for !h.EstaVacia() {
		actual := h.Desencolar()
		v := actual.vertice
		if visitado[v] {
			continue
		}
		visitado[v] = true

		for _, w := range g.Adyacentes(v) {
			arista, _ := g.ObtenerArista(v, w)
			peso := 1.0 / float64(arista.Frecuencia)
			nuevaDist := distancia[v] + peso

			if nuevaDist < distancia[w] {
				distancia[w] = nuevaDist
				padres[w] = []string{v}
				h.Encolar(nodo{w, nuevaDist})
			} else if nuevaDist == distancia[w] {
				padres[w] = append(padres[w], v)
			}
		}
	}

	return distancia, padres
}
