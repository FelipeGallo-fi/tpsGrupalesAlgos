// Maneja las conexiones entre aeropuertos (origen → destino → vuelos).
// Usa un doble diccionario para guardar y recuperar vuelos directos ordenados por fecha.
// Se usa exclusivamente para el comando `siguiente_vuelo`.
package comandos

import (
	"slices"
	"time"
	TDAvuelo "tp2/TDAvuelo"
	hash "tp2/tdas/hash"
)

func BuscarSiguienteVuelo(
	conexiones hash.Diccionario[string, hash.Diccionario[string, []*TDAvuelo.Vuelo]],
	origen, destino string,
	fecha time.Time,
) *TDAvuelo.Vuelo {
	if !conexiones.Pertenece(origen) {
		return nil
	}
	destinos := conexiones.Obtener(origen)
	if !destinos.Pertenece(destino) {
		return nil
	}
	vuelos := destinos.Obtener(destino)

	i := slices.IndexFunc(vuelos, func(v *TDAvuelo.Vuelo) bool {
		return !v.Fecha.Before(fecha)
	})

	for i < len(vuelos) {
		if vuelos[i].Cancelado != _Cancelado {
			return vuelos[i]
		}
		i++
	}

	return nil
}

func EliminarVuelo(
	conexiones hash.Diccionario[string, hash.Diccionario[string, []*TDAvuelo.Vuelo]],
	vuelo *TDAvuelo.Vuelo,
) {
	if !conexiones.Pertenece(vuelo.Origen) {
		return
	}

	destinos := conexiones.Obtener(vuelo.Origen)
	if !destinos.Pertenece(vuelo.Destino) {
		return
	}

	lista := destinos.Obtener(vuelo.Destino)
	var nuevaLista []*TDAvuelo.Vuelo
	for _, v := range lista {
		if v.Codigo != vuelo.Codigo {
			nuevaLista = append(nuevaLista, v)
		}
	}

	if len(nuevaLista) > 0 {
		destinos.Guardar(vuelo.Destino, nuevaLista)
	} else {
		destinos.Borrar(vuelo.Destino)
	}


	if destinos.Cantidad() > 0 {
		conexiones.Guardar(vuelo.Origen, destinos)
	}else {
		conexiones.Borrar(vuelo.Origen)
	}
	
	
}
