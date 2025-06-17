// Lógica auxiliar para manejar los vuelos ordenados por fecha de despegue.
// Permite insertar vuelos al ABB por fecha, recorrerlos por rango y usarlos en ver_tablero o borrar.
package comandos

import (
	"sort"
	"time"
	TDAvuelo "tp2/TDAvuelo"
	abb "tp2/tdas/abb"
)

func VuelosEnRango(abb abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo], desde, hasta time.Time, esDescendente bool) []*TDAvuelo.Vuelo {
	desde = desde.Truncate(24 * time.Hour)
	hasta = hasta.Truncate(24 * time.Hour)
	var resultado []*TDAvuelo.Vuelo
	abb.IterarRango(&desde, &hasta, func(_ time.Time, lista []*TDAvuelo.Vuelo) bool {
		resultado = append(resultado, lista...)
		return true
	})

	
	sort.SliceStable(resultado, func(i, j int) bool {
		if resultado[i].Fecha.Equal(resultado[j].Fecha) {
			return resultado[i].Codigo < resultado[j].Codigo
		}
		if esDescendente {
			return resultado[i].Fecha.After(resultado[j].Fecha)
		}
		return resultado[i].Fecha.Before(resultado[j].Fecha)
	})


	return resultado
}

func EliminarVuelosEnRango(
	vuelosPorFecha abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo],
	desde, hasta time.Time,
	procesarVuelo func(v *TDAvuelo.Vuelo),
) {
	var clavesABorrar []time.Time

	vuelosPorFecha.IterarRango(&desde, &hasta, func(fecha time.Time, lista []*TDAvuelo.Vuelo) bool {
		for _, v := range lista {
			procesarVuelo(v)
		}
		clavesABorrar = append(clavesABorrar, fecha)
		return true
	})

	for _, fecha := range clavesABorrar {
		vuelosPorFecha.Borrar(fecha)
	}
}
