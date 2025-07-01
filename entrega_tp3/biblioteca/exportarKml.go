package biblioteca

import (
	"fmt"
	"os"
)

func ExportarKML(nombre string, rutas [][]string, coords map[string][2]float64) error {
	f, err := os.Create(nombre)
	if err != nil {
		return err
	}
	defer f.Close()

	f.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")
	f.WriteString(`<kml xmlns="http://earth.google.com/kml/2.1">` + "\n")
	f.WriteString(`<Document>` + "\n")
	f.WriteString(`<name>Rutas exportadas</name>` + "\n")
	f.WriteString(`<description>Archivo KML generado por FlyCombi</description>` + "\n")

	for _, ruta := range rutas {
		f.WriteString(`<Placemark><LineString><coordinates>`)
		for _, codigo := range ruta {
			coord, ok := coords[codigo]
			if !ok {
				continue
			}
			// Primero LONGITUD, luego LATITUD
			f.WriteString(fmt.Sprintf("%f,%f ", coord[1], coord[0]))
		}
		f.WriteString(`</coordinates></LineString></Placemark>` + "\n")
	}

	f.WriteString(`</Document></kml>`)
	return nil
}
