// Pruebas unitarias del TDA Vuelo. Verifica el parseo correcto desde CSV
// y la visualización esperada del vuelo.
package TDAvuelo_test

import (
	"testing"

	"tp2/gestion_vuelos/TDAvuelo"

	"github.com/stretchr/testify/require"
)

func TestParsearVueloValido(t *testing.T) {
	linea := "4608,OO,PDX,SEA,N812SK,8,2018-04-10T23:22:55,05,43,0"
	v, err := TDAvuelo.ParsearVuelo(linea)
	require.NoError(t, err)
	require.Equal(t, "4608", v.Codigo)
	require.Equal(t, 8, v.Prioridad)
	require.Equal(t, "SEA", v.Destino)
}
