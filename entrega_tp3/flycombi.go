package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"tp3/biblioteca"
	"tp3/tdas/grafo"
	"tp3/utilidades"
)

var ultimaRuta [][]string

func parsearCSVLine(line string) ([]string, error) {
	reader := csv.NewReader(strings.NewReader(line))
	reader.TrimLeadingSpace = true
	return reader.Read()
}

func CargarDatos(aeropuertos, vuelos string) (grafo.Grafo[string, utilidades.Arista], map[string][]string, map[string][2]float64, error) {
	ciudades := make(map[string][]string)
	coordenadas := make(map[string][2]float64)
	g := grafo.CrearGrafo[string, utilidades.Arista](false)

	if err := cargarAeropuertos(aeropuertos, g, ciudades, coordenadas); err != nil {
		return nil, nil, nil, err
	}
	if err := cargarVuelos(vuelos, g); err != nil {
		return nil, nil, nil, err
	}
	return g, ciudades, coordenadas, nil
}

func cargarAeropuertos(path string, g grafo.Grafo[string, utilidades.Arista], ciudades map[string][]string, coords map[string][2]float64) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	for _, linea := range records {
		if len(linea) < 4 {
			continue
		}
		ciudad := linea[0]
		codigo := linea[1]
		lat, _ := strconv.ParseFloat(linea[2], 64)
		lon, _ := strconv.ParseFloat(linea[3], 64)

		ciudades[ciudad] = append(ciudades[ciudad], codigo)
		coords[codigo] = [2]float64{lat, lon}
		g.AgregarVertice(codigo)
	}
	return nil
}

func cargarVuelos(path string, g grafo.Grafo[string, utilidades.Arista]) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		return err
	}
	for _, linea := range records {
		if len(linea) < 5 {
			continue
		}
		origen := linea[0]
		destino := linea[1]
		tiempo, _ := strconv.ParseFloat(linea[2], 64)
		precio, _ := strconv.ParseFloat(linea[3], 64)
		frecuencia, _ := strconv.Atoi(linea[4])

		g.AgregarArista(origen, destino, utilidades.Arista{
			Tiempo:     tiempo,
			Precio:     precio,
			Frecuencia: frecuencia,
		})
	}

	return nil
}

func ejecutarComandos(g grafo.Grafo[string, utilidades.Arista], ciudades map[string][]string, coordenadas map[string][2]float64) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		linea = strings.TrimSpace(linea)
		if linea == "" {
			continue
		}

		campos := strings.Fields(linea)
		if len(campos) == 0 {
			continue
		}

		comando := campos[0]
		args := strings.TrimPrefix(linea, comando+" ")

		switch comando {
		case "camino_mas":
			comandoCaminoMas(args, g, ciudades)

		case "camino_escalas":
			comandoCaminoEscalas(args, g, ciudades)

		case "centralidad":
			n, err := strconv.Atoi(args)
			if err != nil {
				fmt.Println("Error: parámetro inválido para centralidad")
				continue
			}
			comandoCentralidad(g, n)

		case "nueva_aerolinea":
			aristas := biblioteca.NuevaAerolinea(g)
			if err := guardarAristasCSV(args, aristas); err != nil {
				fmt.Println("Error al guardar archivo:", err)
			} else {
				fmt.Println("OK")
			}

		case "itinerario":
			comandoItinerario(args, g, ciudades)

		case "exportar_kml":
			if err := biblioteca.ExportarKML(args, ultimaRuta, coordenadas); err != nil {
				fmt.Println("Error al exportar KML:", err)
			} else {
				fmt.Println("OK")
			}
		}
	}
}

func guardarAristasCSV(nombre string, aristas []utilidades.AristaCSV) error {
	f, err := os.Create(nombre)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	for _, a := range aristas {
		linea := []string{
			a.Origen,
			a.Destino,
			strconv.Itoa(int(a.Tiempo)),
			strconv.Itoa(int(a.Precio)),
			strconv.Itoa(a.Frecuencia),
		}
		if err := w.Write(linea); err != nil {
			return err
		}
	}
	w.Flush()
	return w.Error()
}

func comandoCaminoMas(args string, g grafo.Grafo[string, utilidades.Arista], ciudades map[string][]string) {
	partes, err := parsearCSVLine(args)
	if err != nil || len(partes) != 3 {
		fmt.Println("Error: se esperan 3 parámetros")
		return
	}
	tipo := strings.TrimSpace(partes[0])
	origen := strings.TrimSpace(partes[1])
	destino := strings.TrimSpace(partes[2])

	listaOrigenes, ok1 := ciudades[origen]
	listaDestinos, ok2 := ciudades[destino]

	if !ok1 || !ok2 {
		fmt.Println("Error: ciudad no encontrada")
		return
	}

	usarPrecio := tipo == "barato"
	mejorCamino := []string{}
	mejorCosto := math.Inf(1)

	for _, o := range listaOrigenes {
		dist, padres := biblioteca.CaminoMinimoDijkstra(g, o, usarPrecio)
		for _, d := range listaDestinos {
			costo, ok := dist[d]
			if !ok || math.IsInf(costo, 1) {
				continue
			}
			camino := biblioteca.ReconstruirCamino(d, padres, o)
			if len(camino) == 0 || camino[0] != o {
				continue
			}
			if costo < mejorCosto {
				mejorCosto = costo
				mejorCamino = camino
			}
		}
	}

	ultimaRuta = [][]string{mejorCamino}

	if len(mejorCamino) == 0 {
		fmt.Println("No se encontró camino")
		return
	}

	fmt.Println(strings.Join(mejorCamino, " -> "))
}

func comandoCaminoEscalas(args string, g grafo.Grafo[string, utilidades.Arista], ciudades map[string][]string) {
	partes := strings.Split(args, ",")
	if len(partes) != 2 {
		fmt.Println("Error: se esperan 2 parámetros")
		return
	}
	origen := partes[0]
	destino := partes[1]

	listaOrigenes, ok1 := ciudades[origen]
	listaDestinos, ok2 := ciudades[destino]
	if !ok1 || !ok2 {
		fmt.Println("Error: ciudad no encontrada")
		return
	}

	mejorCamino := []string{}
	mejorEscalas := -1

	for _, o := range listaOrigenes {
		dist, padres := biblioteca.CaminoMinimoEscalas(g, o)
		for _, d := range listaDestinos {
			escalas, ok := dist[d]
			if !ok {
				continue
			}
			camino := biblioteca.ReconstruirCamino(d, padres, o)
			if len(camino) == 0 || camino[0] != o {
				continue
			}

			if mejorEscalas < 0 || escalas < mejorEscalas {
				mejorEscalas = escalas
				mejorCamino = camino
			}
		}
	}

	ultimaRuta = [][]string{mejorCamino}

	if len(mejorCamino) == 0 {
		fmt.Println("No se encontró camino")
		return
	}
	fmt.Println(strings.Join(mejorCamino, " -> "))
}

func comandoItinerario(path string, g grafo.Grafo[string, utilidades.Arista], ciudades map[string][]string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("Error abriendo archivo")
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if !scanner.Scan() {
		fmt.Println("Archivo vacío")
		return
	}

	linea := scanner.Text()
	ciudadesOrden := strings.Split(linea, ",")
	restricciones := [][2]string{}
	for scanner.Scan() {
		partes := strings.Split(scanner.Text(), ",")
		if len(partes) == 2 {
			restricciones = append(restricciones, [2]string{partes[0], partes[1]})
		}
	}

	orden, ok := biblioteca.OrdenTopologico(ciudadesOrden, restricciones)
	if !ok {
		fmt.Println("No se puede resolver el itinerario (ciclo)")
		return
	}
	fmt.Println(strings.Join(orden, ", "))

	for i := 0; i < len(orden)-1; i++ {
		o := orden[i]
		d := orden[i+1]
		encontrado := false
		for _, ori := range ciudades[o] {
			_, padres := biblioteca.CaminoMinimoEscalas(g, ori)
			for _, dest := range ciudades[d] {
				camino := biblioteca.ReconstruirCamino(dest, padres, o)
				if len(camino) > 0 && camino[0] == ori {
					fmt.Println(strings.Join(camino, " -> "))
					encontrado = true
					break
				}
			}
			if encontrado {
				break
			}
		}
	}
}

func comandoCentralidad(g grafo.Grafo[string, utilidades.Arista], n int) {
	centralidades := biblioteca.Centralidad(g)
	top := biblioteca.TopN(centralidades, n)
	fmt.Println(strings.Join(top, ", "))
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Uso: ./flycombi aeropuertos.csv vuelos.csv")
		return
	}

	g, ciudades, coordenadas, err := CargarDatos(os.Args[1], os.Args[2])

	if err != nil {
		fmt.Println("Error al cargar datos:", err)
		return
	}

	ejecutarComandos(g, ciudades, coordenadas)

}
