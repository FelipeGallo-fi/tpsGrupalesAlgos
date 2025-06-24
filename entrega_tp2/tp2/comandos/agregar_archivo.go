package comandos

import (
	"bufio"
	"fmt"
	"os"
	"time"
	"tp2/TDAvuelo"
	hash "tp2/tdas/hash"
)

func AgregarArchivo(nombreArchivo string) error {
	file, err := os.Open(nombreArchivo)
	if err != nil {
		fmt.Fprintln(os.Stderr, _ErrorAgregarArchivo)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linea := scanner.Text()
		vuelo, err := TDAvuelo.ParsearVuelo(linea)
		if err != nil {
			fmt.Fprintln(os.Stderr, _ErrorAgregarArchivo)
			return err
		}
		if vuelosPorCodigo.Pertenece(vuelo.Codigo) {
			eliminarVueloDeFechaPorCodigo(vuelo.Codigo)
			eliminarVueloDeConexionesPorCodigo(vuelo.Codigo)
			vuelosPorCodigo.Borrar(vuelo.Codigo)
		}

		vuelosPorCodigo.Guardar(vuelo.Codigo, vuelo)
		insertarVueloPorFecha(vuelo)

		if vuelo.Cancelado != _Cancelado {
			insertarVueloEnConexiones(vuelo)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, _ErrorAgregarArchivo)
		return err
	}

	fmt.Println(_MensajeOK)
	return nil
}

func insertarVueloPorFecha(vuelo *TDAvuelo.Vuelo) {
	fecha := vuelo.Fecha
	listaFiltrada := []*TDAvuelo.Vuelo{}

	if vuelosPorFecha.Pertenece(fecha) {
		original := vuelosPorFecha.Obtener(fecha)
		for _, v := range original {
			if v.Codigo != vuelo.Codigo {
				listaFiltrada = append(listaFiltrada, v)
			}
		}
	}

	listaFiltrada = insertarOrdenadoPorFecha(listaFiltrada, vuelo)
	vuelosPorFecha.Guardar(fecha, listaFiltrada)
}

func insertarVueloEnConexiones(vuelo *TDAvuelo.Vuelo) {
	origen := vuelo.Origen
	destino := vuelo.Destino

	var destinos hash.Diccionario[string, []*TDAvuelo.Vuelo]
	if conexiones.Pertenece(origen) {
		destinos = conexiones.Obtener(origen)
	} else {
		destinos = hash.CrearHash[string, []*TDAvuelo.Vuelo]()
		conexiones.Guardar(origen, destinos)
	}

	lista := []*TDAvuelo.Vuelo{}
	if destinos.Pertenece(destino) {
		for _, v := range destinos.Obtener(destino) {
			if v.Codigo != vuelo.Codigo {
				lista = append(lista, v)
			}
		}
	}

	lista = insertarOrdenadoPorFecha(lista, vuelo)
	destinos.Guardar(destino, lista)
}

func insertarOrdenadoPorFecha(lista []*TDAvuelo.Vuelo, vuelo *TDAvuelo.Vuelo) []*TDAvuelo.Vuelo {
	for i, v := range lista {
		if vuelo.Fecha.Before(v.Fecha) || (vuelo.Fecha.Equal(v.Fecha) && vuelo.Codigo < v.Codigo) {
			return append(lista[:i], append([]*TDAvuelo.Vuelo{vuelo}, lista[i:]...)...)
		}
	}
	return append(lista, vuelo)
}

func eliminarVueloDeFechaPorCodigo(codigo string) {
	iter := vuelosPorFecha.Iterador()
	var fechasAEliminar []time.Time

	for iter.HaySiguiente() {
		fecha, lista := iter.VerActual()
		nueva := []*TDAvuelo.Vuelo{}
		eliminado := false

		for _, v := range lista {
			if v.Codigo != codigo {
				nueva = append(nueva, v)
			} else {
				eliminado = true
			}
		}

		if eliminado {
			if len(nueva) == 0 {
				fechasAEliminar = append(fechasAEliminar, fecha)
			} else {
				vuelosPorFecha.Guardar(fecha, nueva)
			}
		}

		iter.Siguiente()
	}

	for _, f := range fechasAEliminar {
		if vuelosPorFecha.Pertenece(f) {
			vuelosPorFecha.Borrar(f)
		}
	}
}

func eliminarVueloDeConexionesPorCodigo(codigo string) {
	iterOrigen := conexiones.Iterador()
	var origenesAEliminar []string

	for iterOrigen.HaySiguiente() {
		origen, destinos := iterOrigen.VerActual()
		iterDestino := destinos.Iterador()
		var destinosAEliminar []string

		for iterDestino.HaySiguiente() {
			destino, lista := iterDestino.VerActual()
			nueva := []*TDAvuelo.Vuelo{}
			eliminado := false

			for _, v := range lista {
				if v.Codigo != codigo {
					nueva = append(nueva, v)
				} else {
					eliminado = true
				}
			}

			if eliminado {
				if len(nueva) == 0 {
					destinosAEliminar = append(destinosAEliminar, destino)
				} else {
					destinos.Guardar(destino, nueva)
				}
			}

			iterDestino.Siguiente()
		}

		for _, d := range destinosAEliminar {
			destinos.Borrar(d)
		}

		if destinos.Cantidad() == 0 {
			origenesAEliminar = append(origenesAEliminar, origen)
		}

		iterOrigen.Siguiente()
	}

	for _, o := range origenesAEliminar {
		conexiones.Borrar(o)
	}
}

func eliminarVueloDeConexiones(vuelo *TDAvuelo.Vuelo) {
	iterOrigen := conexiones.Iterador()
	var origenesAEliminar []string

	for iterOrigen.HaySiguiente() {
		origen, destinos := iterOrigen.VerActual()
		iterDestino := destinos.Iterador()
		var destinosAEliminar []string

		for iterDestino.HaySiguiente() {
			destino, lista := iterDestino.VerActual()
			nueva := []*TDAvuelo.Vuelo{}
			eliminado := false

			for _, v := range lista {
				if v.Codigo != vuelo.Codigo {
					nueva = append(nueva, v)
				} else {
					eliminado = true
				}
			}

			if eliminado {
				if len(nueva) == 0 {
					destinosAEliminar = append(destinosAEliminar, destino)
				} else {
					destinos.Guardar(destino, nueva)
				}
			}

			iterDestino.Siguiente()
		}

		for _, d := range destinosAEliminar {
			destinos.Borrar(d)
		}

		if destinos.Cantidad() == 0 {
			origenesAEliminar = append(origenesAEliminar, origen)
		}

		iterOrigen.Siguiente()
	}

	for _, o := range origenesAEliminar {
		if conexiones.Pertenece(o) {
			conexiones.Borrar(o)
		}
	}
}
