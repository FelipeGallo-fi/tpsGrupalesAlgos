package vuelo

import (
	prioridad "tp2/tdas/cola_prioridad"
	fechas "tp2/tdas/fechas"
)

type Vuelo interface {
    ObtenerCodigo() string // obtengo el codigo del vuelo 
    ObtenerFecha() *fechas.Fecha       // fecha y hora
    ObtenerPrioridad() *prioridad.ColaPrioridad[int] // prioridad de vuelo 
    ObtenerInformacion() string // o devolver los campos por separado
}
