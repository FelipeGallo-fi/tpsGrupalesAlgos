package comandos

import (
	"fmt"
	"os"
	"strconv"
	"time"
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

	desde = desde.Truncate(24 * time.Hour) //PRUEBA
	hasta = hasta.Truncate(24 * time.Hour)	//PRUEBA


	if err1 != nil || err2 != nil || errK != nil || k <= 0 || hasta.Before(desde) {
		fmt.Fprintln(os.Stderr, _ErrorVerTablero)
		return
	}

	modoDesc := (modo == _ModoDesc)
	vuelos := VuelosEnRango(vuelosPorFecha, desde, hasta, modoDesc)


	if len(vuelos) > k {
		vuelos = vuelos[:k]
	}

	for _, v := range vuelos {
		fmt.Printf("%s - %s\n", v.Fecha.Format(_Fecha), v.Codigo)
	}

	fmt.Println(_MensajeOK)
}
