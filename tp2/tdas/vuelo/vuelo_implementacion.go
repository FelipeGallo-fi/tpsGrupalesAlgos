package vuelo

import (
	"fmt"
	prioridad "tp2/tdas/cola_prioridad"
	fechas "tp2/tdas/fechas"
)

type vueloImpl struct{
	codigo string
	fecha *fechas.Fecha
	prioridad *prioridad.ColaPrioridad[int]
	info string
}


func NuevoVuelo(codigo string, fecha *fechas.Fecha, prioridad *prioridad.ColaPrioridad[int], info string) Vuelo {
    return &vueloImpl{
        codigo:    codigo,
        fecha:     fecha,
        prioridad: prioridad,
        info:      info,
    }
}


func (v *vueloImpl) ObtenerCodigo() string {
    return v.codigo
}

func (v *vueloImpl) ObtenerFecha() *fechas.Fecha {
    return v.fecha
}

func (v *vueloImpl) ObtenerPrioridad() *prioridad.ColaPrioridad[int] {
    return v.prioridad
}

func (v *vueloImpl) ObtenerInformacion() string {
    return v.info
}

func (v *vueloImpl) String() string {
    return fmt.Sprintf("Vuelo %s - Fecha: %v - Prioridad: %v - Info: %s", v.codigo, v.fecha, v.prioridad, v.info)
}