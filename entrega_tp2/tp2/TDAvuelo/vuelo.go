// Define la estructura de un vuelo y su parseo desde una línea CSV.
// También provee funciones para mostrar un vuelo en el formato requerido por los comandos.
package TDAvuelo

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	_Fecha = "2006-01-02T15:04:05"
)

type Vuelo struct {
	Codigo      string
	Aerolinea   string
	Origen      string
	Destino     string
	Matricula   string
	Prioridad   int
	Fecha       time.Time
	Retraso     int
	TiempoVuelo int
	Cancelado   int
}

func ParsearVuelo(linea string) (*Vuelo, error) {
	partes := strings.Split(strings.TrimSpace(linea), ",")
	if len(partes) != 10 {
		return nil, fmt.Errorf("línea inválida: %s", linea)
	}

	prioridad, err := strconv.Atoi(partes[5])
	if err != nil {
		return nil, err
	}
	fecha, err := time.Parse(_Fecha, partes[6])
	
	if err != nil {
		return nil, err
	}

	retraso, err := strconv.Atoi(partes[7])
	if err != nil {
		return nil, err
	}

	tiempoVuelo, err := strconv.Atoi(partes[8])
	if err != nil {
		return nil, err
	}

	cancelado, err := strconv.Atoi(partes[9])
	if err != nil {
		return nil, err
	}

	return &Vuelo{
		Codigo:      partes[0],
		Aerolinea:   partes[1],
		Origen:      partes[2],
		Destino:     partes[3],
		Matricula:   partes[4],
		Prioridad:   prioridad,
		Fecha:       fecha,
		Retraso:     retraso,
		TiempoVuelo: tiempoVuelo,
		Cancelado:   cancelado,
	}, nil
}

func (v *Vuelo) String() string {
	return fmt.Sprintf("%s %s %s %s %s %d %s %02d %d %d",
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
