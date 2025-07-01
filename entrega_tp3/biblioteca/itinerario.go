package biblioteca

import (
	"tp3/tdas/cola"
	"tp3/tdas/grafo"
)

func OrdenTopologico(ciudades []string, restricciones [][2]string) ([]string, bool) {
	g := grafo.CrearGrafo[string, struct{}](true)
	grados := make(map[string]int)

	for _, ciudad := range ciudades {
		g.AgregarVertice(ciudad)
		grados[ciudad] = 0
	}
	for _, restric := range restricciones {
		desde := restric[0]
		hasta := restric[1]
		if !g.ExisteVertice(desde) {
			g.AgregarVertice(desde)
			grados[desde] = 0
		}
		if !g.ExisteVertice(hasta) {
			g.AgregarVertice(hasta)
			grados[hasta] = 0
		}
		g.AgregarArista(desde, hasta, struct{}{})
		grados[hasta]++
	}

	q := cola.CrearColaEnlazada[string]()
	for v, grado := range grados {
		if grado == 0 {
			q.Encolar(v)
		}
	}

	var orden []string
	for !q.EstaVacia() {
		v := q.Desencolar()
		orden = append(orden, v)
		for _, w := range g.Adyacentes(v) {
			grados[w]--
			if grados[w] == 0 {
				q.Encolar(w)
			}
		}
	}

	if len(orden) != len(grados) {
		return nil, false // hay ciclo
	}
	return orden, true
}
