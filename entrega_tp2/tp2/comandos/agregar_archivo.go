package comandos

import (
	"bufio"
	"fmt"
	"os"
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
			viejo := vuelosPorCodigo.Obtener(vuelo.Codigo)
			eliminarVueloDeFecha(viejo)
			eliminarVueloDeConexiones(viejo)
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
	listaFiltrada := []*TDAvuelo.Vuelo{}
	if vuelosPorFecha.Pertenece(vuelo.Fecha) {
		original := vuelosPorFecha.Obtener(vuelo.Fecha)
		for _, v := range original {
			if v.Codigo != vuelo.Codigo {
				listaFiltrada = append(listaFiltrada, v)
			}
		}
	}
	listaFiltrada = insertarOrdenadoPorFecha(listaFiltrada, vuelo)
	vuelosPorFecha.Guardar(vuelo.Fecha, listaFiltrada)
}

func insertarVueloEnConexiones(vuelo *TDAvuelo.Vuelo) {
	var destinos hash.Diccionario[string, []*TDAvuelo.Vuelo]
	if conexiones.Pertenece(vuelo.Origen) {
		destinos = conexiones.Obtener(vuelo.Origen)
	} else {
		destinos = hash.CrearHash[string, []*TDAvuelo.Vuelo]()
		conexiones.Guardar(vuelo.Origen, destinos)
	}

	var lista []*TDAvuelo.Vuelo
	if destinos.Pertenece(vuelo.Destino) {
		lista = destinos.Obtener(vuelo.Destino)
	}
	lista = insertarOrdenadoPorFecha(lista, vuelo)
	destinos.Guardar(vuelo.Destino, lista)
}

func insertarOrdenadoPorFecha(lista []*TDAvuelo.Vuelo, vuelo *TDAvuelo.Vuelo) []*TDAvuelo.Vuelo {
	for i, v := range lista {
		if vuelo.Fecha.Before(v.Fecha) || (vuelo.Fecha.Equal(v.Fecha) && vuelo.Codigo < v.Codigo) {
			return append(lista[:i], append([]*TDAvuelo.Vuelo{vuelo}, lista[i:]...)...)
		}
	}
	return append(lista, vuelo)
}

func eliminarVueloDeFecha(vuelo *TDAvuelo.Vuelo) {
	if !vuelosPorFecha.Pertenece(vuelo.Fecha) {
		return
	}
	lista := vuelosPorFecha.Obtener(vuelo.Fecha)
	nueva := []*TDAvuelo.Vuelo{}
	for _, v := range lista {
		if v.Codigo != vuelo.Codigo {
			nueva = append(nueva, v)
		}
	}
	if len(nueva) == 0 {
		vuelosPorFecha.Borrar(vuelo.Fecha)
	} else {
		vuelosPorFecha.Guardar(vuelo.Fecha, nueva)
	}
}

func eliminarVueloDeConexiones(vuelo *TDAvuelo.Vuelo) {
	if !conexiones.Pertenece(vuelo.Origen) {
		return
	}
	destinos := conexiones.Obtener(vuelo.Origen)
	if !destinos.Pertenece(vuelo.Destino) {
		return
	}
	lista := destinos.Obtener(vuelo.Destino)
	nueva := []*TDAvuelo.Vuelo{}
	for _, v := range lista {
		if v.Codigo != vuelo.Codigo {
			nueva = append(nueva, v)
		}
	}
	if len(nueva) > 0 {
		destinos.Guardar(vuelo.Destino, nueva)
	} else {
		destinos.Borrar(vuelo.Destino)
		if destinos.Cantidad() == 0 {
			conexiones.Borrar(vuelo.Origen)
		}
	}
}
