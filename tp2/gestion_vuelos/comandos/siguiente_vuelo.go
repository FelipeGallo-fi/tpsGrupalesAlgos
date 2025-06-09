package comandos

import (
	"fmt"
	"os"
	"slices"
	"time"
	TDAvuelo "tp2/gestion_vuelos/TDAvuelo"
)

func SiguienteVuelo(parametros []string) {
	origen := parametros[0]
	destino := parametros[1]
	fechaStr := parametros[2]

	fecha, err := time.Parse(_Fecha, fechaStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, _ErrorSiguienteVuelo)
		return
	}

	if !conexiones.Pertenece(origen) {
		fmt.Fprintln(os.Stderr, _ErrorSiguienteVuelo)
		return
	}

	hashDestino := conexiones.Obtener(origen)
	if !hashDestino.Pertenece(destino) {
		fmt.Fprintln(os.Stderr, _ErrorSiguienteVuelo)
		return
	}

	vuelos := hashDestino.Obtener(destino)

	i := slices.IndexFunc(vuelos, func(v *TDAvuelo.Vuelo) bool {
		return !v.Fecha.Before(fecha)
	})

	for i < len(vuelos) {
		if vuelos[i].Cancelado != _Cancelado {
			fmt.Printf("%s - %s\n", vuelos[i].Codigo, vuelos[i].Fecha.Format(_Fecha))
			fmt.Println(_MensajeOK)
			return
		}
		i++
	}

	fmt.Fprintln(os.Stderr, _ErrorSiguienteVuelo)
}
