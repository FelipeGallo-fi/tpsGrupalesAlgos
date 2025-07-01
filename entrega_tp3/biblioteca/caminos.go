package biblioteca

import (
	"math"

	TDACOLA "tp3/tdas/cola"
	"tp3/tdas/grafo"
	TDAHEAP "tp3/tdas/heap"
	"tp3/utilidades"
)

type nodoCamino struct {
	Aeropuerto string
	Prioridad  float64
}

func cmpNodoCamino(a, b nodoCamino) int {
	if a.Prioridad < b.Prioridad {
		return -1
	}
	if a.Prioridad > b.Prioridad {
		return 1
	}
	return 0
}

func CaminoMinimoDijkstra(g grafo.Grafo[string, utilidades.Arista], origen string, usarPrecio bool) (map[string]float64, map[string]string) {
	dist := make(map[string]float64)
	padre := make(map[string]string)

	for _, v := range g.Vertices() {
		dist[v] = math.Inf(1)
	}
	dist[origen] = 0

	heap := TDAHEAP.CrearHeap(cmpNodoCamino)
	heap.Encolar(nodoCamino{Aeropuerto: origen, Prioridad: 0})

	for !heap.EstaVacia() {
		nodo := heap.Desencolar()
		actual := nodo.Aeropuerto
		costo := nodo.Prioridad

		if costo > dist[actual] {
			continue
		}

		for _, vecino := range g.Adyacentes(actual) {
			arista, ok := g.ObtenerArista(actual, vecino)
			if !ok {
				continue
			}

			var peso float64
			if usarPrecio {
				peso = arista.Precio
			} else {
				peso = arista.Tiempo
			}

			if dist[actual]+peso < dist[vecino] {
				dist[vecino] = dist[actual] + peso
				padre[vecino] = actual
				heap.Encolar(nodoCamino{Aeropuerto: vecino, Prioridad: dist[vecino]})
			}
		}
	}

	return dist, padre
}

func ReconstruirCamino(destino string, padre map[string]string, origen string) []string {
	camino := []string{}
	actual := destino
	for {
		camino = append([]string{actual}, camino...)
		if actual == origen {
			break
		}
		p, ok := padre[actual]
		if !ok {
			return nil
		}
		actual = p
	}
	return camino
}

func CaminoMinimoEscalas(g grafo.Grafo[string, utilidades.Arista], origen string) (map[string]int, map[string]string) {
	dist := make(map[string]int)
	padre := make(map[string]string)
	visitado := make(map[string]bool)

	q := TDACOLA.CrearColaEnlazada[string]()
	q.Encolar(origen)
	visitado[origen] = true
	dist[origen] = 0

	for !q.EstaVacia() {
		actual := q.Desencolar()

		for _, vecino := range g.Adyacentes(actual) {
			if !visitado[vecino] {
				visitado[vecino] = true
				dist[vecino] = dist[actual] + 1
				padre[vecino] = actual
				q.Encolar(vecino)
			}
		}
	}
	return dist, padre
}
