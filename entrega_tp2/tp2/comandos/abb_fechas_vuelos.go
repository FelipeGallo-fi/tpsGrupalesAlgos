package comandos

import (
	"sort"
	"time"
	TDAvuelo "tp2/TDAvuelo"
	abb "tp2/tdas/abb"
)

func VuelosEnRango(abb abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo], desde, hasta time.Time, esDescendente bool, k int) []*TDAvuelo.Vuelo {
	var resultado []*TDAvuelo.Vuelo

	desde = TDAvuelo.NormalizarFecha(desde)
	hasta = TDAvuelo.NormalizarFecha(hasta)

	abb.IterarRango(&desde, &hasta, func(_ time.Time, lista []*TDAvuelo.Vuelo) bool {
		for _, vuelo := range lista {
			if vuelosPorCodigo.Pertenece(vuelo.Codigo) {
				resultado = append(resultado, vuelo)
			}
		}
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

func EliminarVuelosEnRango(vuelosPorFecha abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo], desde, hasta time.Time, procesarVuelo func(v *TDAvuelo.Vuelo, fecha time.Time)) {
	var vuelosAEliminar []struct {
		vuelo *TDAvuelo.Vuelo
		fecha time.Time
	}

	desde = TDAvuelo.NormalizarFecha(desde)
	hasta = TDAvuelo.NormalizarFecha(hasta)

	vuelosPorFecha.IterarRango(&desde, &hasta, func(fecha time.Time, lista []*TDAvuelo.Vuelo) bool {

		fecha = TDAvuelo.NormalizarFecha(fecha)

		for _, v := range lista {
			vuelosAEliminar = append(vuelosAEliminar, struct {
				vuelo *TDAvuelo.Vuelo
				fecha time.Time
			}{v, fecha})
		}
		return true
	})

	sort.SliceStable(vuelosAEliminar, func(i, j int) bool {
		if vuelosAEliminar[i].fecha.Equal(vuelosAEliminar[j].fecha) {
			return vuelosAEliminar[i].vuelo.Codigo < vuelosAEliminar[j].vuelo.Codigo
		}
		return vuelosAEliminar[i].fecha.Before(vuelosAEliminar[j].fecha)
	})

	for _, item := range vuelosAEliminar {
		procesarVuelo(item.vuelo, item.fecha)
	}
}
