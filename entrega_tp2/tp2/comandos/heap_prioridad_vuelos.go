package comandos

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"tp2/TDAvuelo"
	heap "tp2/tdas/heap"
)

type vueloPrioritario struct {
	vuelo *TDAvuelo.Vuelo
}

func cmpVuelos(a, b vueloPrioritario) int {
	if a.vuelo.Prioridad != b.vuelo.Prioridad {
		return a.vuelo.Prioridad - b.vuelo.Prioridad
	}
	return strings.Compare(b.vuelo.Codigo, a.vuelo.Codigo)
}

func PrioridadVuelos(parametros []string) {
	kStr := parametros[0]
	k, err := strconv.Atoi(kStr)
	if err != nil || k <= 0 {
		fmt.Fprintln(os.Stderr, _ErrorPrioridadVuelos)
		return
	}

	h := heap.CrearHeap(cmpVuelos)

	vuelosPorCodigo.Iterar(func(_ string, v *TDAvuelo.Vuelo) bool {
		h.Encolar(vueloPrioritario{vuelo: v})
		return true
	})

	for i := 0; i < k && !h.EstaVacia(); i++ {
		v := h.Desencolar().vuelo
		fmt.Printf("%d - %s\n", v.Prioridad, v.Codigo)
	}

	fmt.Println(_MensajeOK)
}
