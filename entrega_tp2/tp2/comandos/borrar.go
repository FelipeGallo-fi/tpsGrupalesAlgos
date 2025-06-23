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

	desdeStr := parametros[0]
	hastaStr := parametros[1]

	desde, err1 := time.Parse(_Fecha, desdeStr)
	hasta, err2 := time.Parse(_Fecha, hastaStr)

	if err1 != nil || err2 != nil || hasta.Before(desde) {
		fmt.Fprintln(os.Stderr, _ErrorBorrar)
		return
	}

	visto := make(map[string]bool)

	procesar := func(v *TDAvuelo.Vuelo) {
		if !visto[v.Codigo] {
			fmt.Println(v.String())
			visto[v.Codigo] = true
		}
		vuelosPorCodigo.Borrar(v.Codigo)
		eliminarVueloDeConexiones(v)
	}

	EliminarVuelosEnRango(vuelosPorFecha, desde, hasta, procesar)

	fmt.Println(_MensajeOK)
}
