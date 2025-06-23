// Lógica auxiliar para manejar los vuelos ordenados por fecha de despegue.
// Permite insertar vuelos al ABB por fecha, recorrerlos por rango y usarlos en ver_tablero o borrar.
package comandos

import (
	"sort"
	"time"
	TDAvuelo "tp2/TDAvuelo"
	abb "tp2/tdas/abb"
)

func VuelosEnRango(abb abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo], desde, hasta time.Time, esDescendente bool, k int) []*TDAvuelo.Vuelo {
	var resultado []*TDAvuelo.Vuelo

	abb.IterarRango(&desde, &hasta, func(_ time.Time, lista []*TDAvuelo.Vuelo) bool {
		resultado = append(resultado, lista...)
		return true
	})

	if esDescendente {
		sort.SliceStable(resultado, func(i, j int) bool {
			if resultado[i].Fecha.Equal(resultado[j].Fecha) {
				return resultado[i].Codigo > resultado[j].Codigo
			}
			return resultado[i].Fecha.After(resultado[j].Fecha)
		})
	} else {
		sort.SliceStable(resultado, func(i, j int) bool {
			if resultado[i].Fecha.Equal(resultado[j].Fecha) {
				return resultado[i].Codigo < resultado[j].Codigo
			}
			return resultado[i].Fecha.Before(resultado[j].Fecha)
		})
	}
	if len(resultado) > k {
		resultado = resultado[:k]
	}

	return resultado
}

func EliminarVuelosEnRango(vuelosPorFecha abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo], desde, hasta time.Time, procesarVuelo func(v *TDAvuelo.Vuelo)) {
	var vuelosAEliminar []*TDAvuelo.Vuelo
	var clavesABorrar []time.Time

	vuelosPorFecha.IterarRango(&desde, &hasta, func(fecha time.Time, lista []*TDAvuelo.Vuelo) bool {
		vuelosAEliminar = append(vuelosAEliminar, lista...)
		clavesABorrar = append(clavesABorrar, fecha)
		return true
	})

	sort.SliceStable(vuelosAEliminar, func(i, j int) bool {
		if vuelosAEliminar[i].Fecha.Equal(vuelosAEliminar[j].Fecha) {
			return vuelosAEliminar[i].Codigo < vuelosAEliminar[j].Codigo
		}
		return vuelosAEliminar[i].Fecha.Before(vuelosAEliminar[j].Fecha)
	})

	for _, v := range vuelosAEliminar {
		procesarVuelo(v)
	}

	for _, fecha := range clavesABorrar {
		if vuelosPorFecha.Pertenece(fecha) {
			vuelosPorFecha.Borrar(fecha)
		}
	}
}
