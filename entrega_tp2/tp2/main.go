package main

import (
	"bufio"
	"os"
	"tp2/comandos"
)

func main() {
	comandos.InicializarEstructuras()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		comandos.EjecutarComando(linea)
	}
}