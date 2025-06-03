package fecha

import "fmt"

type fecha struct {
 	year  int
    mes    int
    dia    int
    hora   int
    minuto int
}


func CrearFecha(year, mes, dia, hora, minuto int) Fecha {
    return &fecha{year, mes, dia, hora, minuto}
}

func (f *fecha) EsAnterior(otra Fecha) bool {
    o := otra.(*fecha)
    if f.year != o.year {
        return f.year < o.year
    }
    if f.mes != o.mes {
        return f.mes < o.mes
    }
    if f.dia != o.dia {
        return f.dia < o.dia
    }
    if f.hora != o.hora {
        return f.hora < o.hora
    }
    return f.minuto < o.minuto
}

func (f *fecha) EsIgual(otra Fecha) bool {
    o := otra.(*fecha)
    return f.year == o.year && f.mes == o.mes && f.dia == o.dia &&
           f.hora == o.hora && f.minuto == o.minuto
}

func (f *fecha) String() string {
    return fmt.Sprintf("%04d-%02d-%02d %02d:%02d", f.year, f.mes, f.dia, f.hora, f.minuto)
}