package comandos

import (
	"fmt"
	"slices"
	"time"
	TDAvuelo "tp2/TDAvuelo"
)

func SiguienteVuelo(parametros []string) {
	if len(parametros) != 3 {
		fmt.Println(_ErrorSiguienteVuelo)
		return
	}

	origen := parametros[0]
	destino := parametros[1]
	fechaStr := parametros[2]

	fecha, err := time.Parse(_Fecha, fechaStr)
	if err != nil {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fechaStr)
		fmt.Println(_MensajeOK)
		return
	}

	fecha = TDAvuelo.NormalizarFecha(fecha)

	if !conexiones.Pertenece(origen) {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fechaStr)
		fmt.Println(_MensajeOK)
		return
	}

	hashDestino := conexiones.Obtener(origen)
	if !hashDestino.Pertenece(destino) {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fechaStr)
		fmt.Println(_MensajeOK)
		return
	}

	vuelos := hashDestino.Obtener(destino)

	i := slices.IndexFunc(vuelos, func(v *TDAvuelo.Vuelo) bool {
		return !v.Fecha.Before(fecha)
	})

	if i == -1 {
		fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fechaStr)
		fmt.Println(_MensajeOK)
		return
	}

	for i < len(vuelos) {
		if vuelos[i].Cancelado == 0 {
			v := vuelos[i]
			fmt.Println(v.String())
			fmt.Println(_MensajeOK)
			return
		}
		i++
	}

	fmt.Printf("No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fechaStr)
	fmt.Println(_MensajeOK)

}
