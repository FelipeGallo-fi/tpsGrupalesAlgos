package vuelo

type Vuelo interface {
    ObtenerCodigo() string // obtengo el codigo del vuelo 
    ObtenerFecha() Fecha        // fecha y hora
    ObtenerPrioridad() int // prioridad de vuelo 
    ObtenerInformacion() string // o devolver los campos por separado
}
