package comandos

import (
	"fmt"
	"os"
	TDAvuelo "tp2/TDAvuelo"
)

func infoVuelo(codigo string) error {
	if !vuelosPorCodigo.Pertenece(codigo) {
		fmt.Fprintln(os.Stderr, _ErrorInfoVuelo)
		return nil
	}

	vuelo := vuelosPorCodigo.Obtener(codigo)
	ImprimirVuelo(vuelo)
	fmt.Println(_MensajeOK)
	return nil
}

func ImprimirVuelo(v *TDAvuelo.Vuelo) {
	fmt.Printf("%s %s %s %s %s %d %s %d %d %d\n",
		v.Codigo,
		v.Aerolinea,
		v.Origen,
		v.Destino,
		v.Matricula,
		v.Prioridad,
		v.Fecha.Format(_Fecha),
		v.Retraso,
		v.TiempoVuelo,
		v.Cancelado,
	)
}
