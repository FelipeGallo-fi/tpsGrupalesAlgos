package fecha

type Fecha interface {
    EsAnterior(f Fecha) bool
    EsIgual(f Fecha) bool
    String() string
}