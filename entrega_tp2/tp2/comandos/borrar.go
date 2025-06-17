package comandos

import (
	"fmt"
	"os"
	"time"
	TDAvuelo "tp2/TDAvuelo"
)

func Borrar(parametros []string) {
	if len(parametros) != 2 {
		fmt.Fprintln(os.Stderr, _ErrorBorrar)
		return
	}

	desde, err1 := time.Parse(_Fecha, parametros[0])
	hasta, err2 := time.Parse(_Fecha, parametros[1])
	desde = desde.Truncate(24 * time.Hour)
	hasta = hasta.Truncate(24 * time.Hour)


	if err1 != nil || err2 != nil || hasta.Before(desde) {
		fmt.Fprintln(os.Stderr, _ErrorBorrar)
		return
	}

	EliminarVuelosEnRango(vuelosPorFecha, desde, hasta, func(v *TDAvuelo.Vuelo) {
		ImprimirVuelo(v)
		vuelosPorCodigo.Borrar(v.Codigo)
		EliminarVuelo(conexiones, v)
	})

	fmt.Println(_MensajeOK)
}
