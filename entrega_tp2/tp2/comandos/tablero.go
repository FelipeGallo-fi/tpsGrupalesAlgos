package comandos

import (
	"fmt"
	"os"
	"strconv"
	"time"
	TDAvuelo "tp2/TDAvuelo"
)

func VerTablero(parametros []string) {
	if len(parametros) != 4 {
		fmt.Fprintln(os.Stderr, _ErrorVerTablero)
		return
	}

	kStr := parametros[0]
	modo := parametros[1]
	desdeStr := parametros[2]
	hastaStr := parametros[3]

	desde, err1 := time.Parse(_Fecha, desdeStr)
	hasta, err2 := time.Parse(_Fecha, hastaStr)
	k, errK := strconv.Atoi(kStr)

	if err1 != nil || err2 != nil || errK != nil || k <= 0 || hasta.Before(desde) {
		fmt.Fprintln(os.Stderr, _ErrorVerTablero)
		return
	}

	desde = TDAvuelo.NormalizarFecha(desde)
	hasta = TDAvuelo.NormalizarFecha(hasta)

	modoDesc := (modo == _ModoDesc)
	vuelos := VuelosEnRango(vuelosPorFecha, desde, hasta, modoDesc, k)
	vistos := make(map[string]bool)

	for _, v := range vuelos {
		clave := v.Codigo + "_" + v.Fecha.Format(time.RFC3339)
		if vistos[clave] {
			continue
		}
		vistos[clave] = true

		fmt.Printf("%s - %s\n", v.Fecha.Format(_Fecha), v.Codigo)
	}

	fmt.Println(_MensajeOK)
}
