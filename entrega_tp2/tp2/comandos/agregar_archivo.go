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

		if vuelo.Cancelado == _Cancelado {
			continue
		}

		guardarVueloPorCodigo(vuelo)
		insertarVueloPorFecha(vuelo)
		insertarVueloEnConexiones(vuelo)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, _ErrorAgregarArchivo)
		return err
	}

	fmt.Println(_MensajeOK)
	return nil
}

func guardarVueloPorCodigo(vuelo *TDAvuelo.Vuelo) {
	vuelosPorCodigo.Guardar(vuelo.Codigo, vuelo)
}

func insertarVueloPorFecha(vuelo *TDAvuelo.Vuelo) {
	var lista []*TDAvuelo.Vuelo

	if vuelosPorFecha.Pertenece(vuelo.Fecha) {
		lista = vuelosPorFecha.Obtener(vuelo.Fecha)
	} else {
		lista = []*TDAvuelo.Vuelo{}
	}

	lista = insertarOrdenadoPorFecha(lista, vuelo)
	vuelosPorFecha.Guardar(vuelo.Fecha, lista)
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
	} else {
		lista = []*TDAvuelo.Vuelo{}
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
