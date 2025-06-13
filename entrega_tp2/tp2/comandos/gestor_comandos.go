package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"tp2/TDAvuelo"
	abb "tp2/tdas/abb"
	hash "tp2/tdas/hash"
)

const (
	_Fecha                = "2006-01-02T15:04:05"
	_ErrorAgregarArchivo  = "Error en comando agregar_archivo"
	_ErrorInfoVuelo       = "Error en comando info_vuelo"
	_ErrorVerTablero      = "Error en comando ver_tablero"
	_ErrorPrioridadVuelos = "Error en comando prioridad_vuelos"
	_ErrorSiguienteVuelo  = "Error en comando siguiente_vuelo"
	_ErrorBorrar          = "Error en comando borrar"
	_MensajeOK            = "OK"
	_Cancelado            = 1
	_ModoAsc              = "asc"
	_ModoDesc             = "desc"
)

var (
	vuelosPorCodigo hash.Diccionario[string, *TDAvuelo.Vuelo]
	vuelosPorFecha  abb.DiccionarioOrdenado[time.Time, []*TDAvuelo.Vuelo]
	conexiones      hash.Diccionario[string, hash.Diccionario[string, []*TDAvuelo.Vuelo]]
)

func InicializarEstructuras() {
	vuelosPorCodigo = hash.CrearHash[string, *TDAvuelo.Vuelo]()
	vuelosPorFecha = abb.CrearABB[time.Time, []*TDAvuelo.Vuelo](func(a, b time.Time) int {
		if a.Before(b) {
			return -1
		} else if a.After(b) {
			return 1
		}
		return 0
	})
	conexiones = hash.CrearHash[string, hash.Diccionario[string, []*TDAvuelo.Vuelo]]()
}

func EjecutarComando(linea string) {
	campos := strings.Fields(linea)
	if len(campos) == 0 {
		return
	}

	switch campos[0] {
	case "agregar_archivo":
		if len(campos) != 2 {
			fmt.Fprintln(os.Stderr, _ErrorAgregarArchivo)
			return
		}
		AgregarArchivo(campos[1])

	case "ver_tablero":
		if len(campos) != 5 {
			fmt.Fprintln(os.Stderr, _ErrorVerTablero)
			return
		}
		_, err := strconv.Atoi(campos[1])
		if err != nil {
			fmt.Fprintln(os.Stderr, _ErrorVerTablero)
			return
		}
		VerTablero(campos[1:5])

	case "info_vuelo":
		if len(campos) != 2 {
			fmt.Fprintln(os.Stderr, _ErrorInfoVuelo)
			return
		}
		infoVuelo(campos[1])

	case "prioridad_vuelos":
		if len(campos) != 2 {
			fmt.Fprintln(os.Stderr, _ErrorPrioridadVuelos)
			return
		}
		PrioridadVuelos([]string{campos[1]})

	case "siguiente_vuelo":
		if len(campos) != 3 {
			fmt.Fprintln(os.Stderr, _ErrorSiguienteVuelo)
			return
		}
		SiguienteVuelo([]string{campos[1], campos[2]})

	case "borrar":
		if len(campos) != 3 {
			fmt.Fprintln(os.Stderr, _ErrorBorrar)
			return
		}
		Borrar([]string{campos[1], campos[2]})

	default:
		fmt.Fprintln(os.Stderr, "Comando desconocido")
	}
}
