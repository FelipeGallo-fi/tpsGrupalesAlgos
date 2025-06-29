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
			vueloAntiguo := vuelosPorCodigo.Obtener(vuelo.Codigo)

			eliminarVueloDeFechaPorCodigoEnFecha(vueloAntiguo.Codigo, vueloAntiguo.Fecha)
			eliminarVueloDeConexionesPorCodigoDesde(vueloAntiguo.Codigo, vueloAntiguo.Origen, vueloAntiguo.Destino)
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
	fecha := TDAvuelo.NormalizarFecha(vuelo.Fecha)
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

func eliminarVueloDeFechaPorCodigoEnFecha(codigo string, fecha time.Time) {
	fecha = TDAvuelo.NormalizarFecha(fecha)

	if !vuelosPorFecha.Pertenece(fecha) {
		return
	}
	listaOriginal := vuelosPorFecha.Obtener(fecha)
	nueva := make([]*TDAvuelo.Vuelo, 0, len(listaOriginal))

	for _, vuelo := range listaOriginal {
		if vuelo.Codigo != codigo {
			nueva = append(nueva, vuelo)
		}
	}

	if len(nueva) == 0 {
		vuelosPorFecha.Borrar(fecha)
	} else {
		vuelosPorFecha.Guardar(fecha, nueva)
	}
}

func eliminarVueloDeConexionesPorCodigoDesde(codigo, origen, destino string) {
	if !conexiones.Pertenece(origen) {
		return
	}
	destinos := conexiones.Obtener(origen)
	if !destinos.Pertenece(destino) {
		return
	}
	lista := destinos.Obtener(destino)
	nueva := []*TDAvuelo.Vuelo{}
	for _, v := range lista {
		if v.Codigo != codigo {
			nueva = append(nueva, v)
		}
	}
	if len(nueva) == 0 {
		destinos.Borrar(destino)
	} else {
		destinos.Guardar(destino, nueva)
	}
	if destinos.Cantidad() == 0 {
		conexiones.Borrar(origen)
	}
}
