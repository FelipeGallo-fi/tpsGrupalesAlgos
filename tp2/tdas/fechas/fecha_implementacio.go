package fecha

type fecha struct {
 	year  int
    mes    int
    dia    int
    hora   int
    minuto int
}


func CrearFecha(year , mes , dia , hora , minuto int) Fecha {
	return &fecha{year , mes ,dia ,hora ,minuto}
}

func (f *fecha)
