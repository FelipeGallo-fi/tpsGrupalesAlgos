package cola_prioridad_test

import (
	TDAHeap "tdas/tpsGrupalesAlgos/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	_DESCENDENTE = 0
	_VOLUMEN     = 100000
	_ASCENDENTE  = 1
)

var _VECTOR_DE_PRIORIDADES_ = []int{1, 2, 3, 3, 7, 23, 34, 60, 100, 2300, 22, 99, 34, 69, 34, 4}

func cmpMaxHeap(a, b int) int {
	return a - b
}

func cmpMinHeap(a, b int) int {
	return b - a
}

func MergeSort(orden int, arr []int) []int { // si recibe 0 , de mayor a menor , si recibe 1 de menor a mayor
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(orden, arr[:mid])
	right := MergeSort(orden, arr[mid:])

	return merge(orden, left, right)
}

func merge(orden int, left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	if orden == 1 {
		for i < len(left) && j < len(right) {
			if left[i] <= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		}
	} else {
		for i < len(left) && j < len(right) {
			if left[i] >= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func TestBencharkEncolarYDesencolar(t *testing.T) {
	h := TDAHeap.CrearHeap(cmpMaxHeap)

	for i := 0; i < _VOLUMEN; i++ {
		h.Encolar(i)
	}

	require.Equal(t, _VOLUMEN-1, h.VerMax(), "El maximo deberia ser 99999")

	for i := 0; i < _VOLUMEN; i++ {
		require.Equal(t, _VOLUMEN-1-i, h.Desencolar(), "Desencolar deberia devolver los valores de mayor a menor")
	}

	require.True(t, h.EstaVacia(), "El heap deberia de estar vacio al desencolar todo")
}

func TestHeapVacio(t *testing.T) {
	h := TDAHeap.CrearHeap(cmpMaxHeap)

	require.True(t, h.EstaVacia())

	require.Equal(t, 0, h.Cantidad(), "La cantidad deberia estar vacio el heap al crearse")

	require.Panics(t, func() { h.VerMax() }, "VerMax en un heap  recien creada deberia panickear")
	require.Panics(t, func() { h.Desencolar() }, "Desencolar en un heap  recien creada deberia panickear")

}

func TestCrearHeapConArrayMax(t *testing.T) {

	vectorMax := MergeSort(_DESCENDENTE, _VECTOR_DE_PRIORIDADES_)
	h := TDAHeap.CrearHeapArr(vectorMax, cmpMaxHeap)

	require.Equal(t, len(vectorMax), h.Cantidad(), "La cantidad deberia ser igual que la del largo de mi array")

	require.Equal(t, 2300, h.VerMax(), "El maximo de mi heap deberia ser el mismo que el de mi arreglo ")

	i := 0
	for !h.EstaVacia() {
		elemento := h.Desencolar()
		require.Equal(t, vectorMax[i], elemento, "Los elementos deberian de ser iguales ya que los elementos a desencolar en su raiz deberian de ser Maximos")
		i++
	}

	require.True(t, h.EstaVacia(), "El heap deberia de estar vacio")

	h.Encolar(23)
	h.Encolar(2)

	require.Equal(t, 23, h.VerMax(), "El maximo de mi heap deberia ser 23 ")

	require.Equal(t, 23, h.Desencolar(), "Deberia devolver la raiz")

}

func TestCrearHeapConArrayMin(t *testing.T) {

	vectorMin := MergeSort(_ASCENDENTE, _VECTOR_DE_PRIORIDADES_)
	h := TDAHeap.CrearHeapArr(vectorMin, cmpMinHeap)

	require.Equal(t, len(vectorMin), h.Cantidad(), "La cantidad deberia ser igual que la del largo de mi array")

	i := 0
	for !h.EstaVacia() {
		elemento := h.Desencolar()
		require.Equal(t, vectorMin[i], elemento, "Los elementos deberian de ser iguales ya que los elementos a desencolar en su raiz deberian de ser minimos")
		i++
	}

	require.True(t, h.EstaVacia(), "El heap deberia de estar vacio")

	h.Encolar(1)
	h.Encolar(1000)
	h.Encolar(10000)

	require.Equal(t, 1, h.Desencolar(), "Desencolar deberia devolver 1")
	require.Equal(t, 1000, h.Desencolar(), "Desencolar deberia devolver 1000")
	require.Equal(t, 10000, h.Desencolar(), "Desencolar deberia devolver 10000")

	require.Panics(t, func() { h.Desencolar() }, "Desencolar en un heap vacio deberia panickear")

}

func TestCrearHeapConArraynil(t *testing.T) {
	h := TDAHeap.CrearHeapArr([]int{}, cmpMaxHeap)

	require.True(t, h.EstaVacia(), "El heap deberia de estar vacio")
	require.Equal(t, 0, h.Cantidad(), "La cantida tiene que ser 0")

	require.Panics(t, func() { h.VerMax() }, "VerMax en un heap  recien creada deberia panickear")
	require.Panics(t, func() { h.Desencolar() }, "Desencolar en un heap  recien creada deberia panickear")

}

func TestVerMaxDspsDeEncolar(t *testing.T) {
	h := TDAHeap.CrearHeap(cmpMaxHeap)

	valores := []int{10, 20, 5, 30, 25, 60}
	maximosEsperados := []int{10, 20, 20, 30, 30, 60}

	for i, val := range valores {
		h.Encolar(val)
		require.Equal(t, maximosEsperados[i], h.VerMax(),
			"Después de encolar %d, el máximo debería ser %d", val, maximosEsperados[i])
	}

	valoresOrdenados := MergeSort(_DESCENDENTE, valores)
	for i, esperado := range valoresOrdenados {
		obtenido := h.Desencolar()
		require.Equal(t, esperado, obtenido,
			"En la posición %d el valor desencolado debería ser %d", i, esperado)
	}

	require.True(t, h.EstaVacia(), "El heap debería estar vacío después de desencolar todo")

	require.Panics(t, func() { h.VerMax() }, "VerMax debería panickear si el heap está vacío")
	require.Panics(t, func() { h.Desencolar() }, "Desencolar debería panickear si el heap está vacío")
}

func TestMinHeapConNegativos(t *testing.T) {
	h := TDAHeap.CrearHeap(cmpMinHeap)

	valores := []int{-1, -50, 0, -3}
	valoresEsperados := MergeSort(_ASCENDENTE, valores)

	for i, v := range valores {
		h.Encolar(v)
		esperadoMin := MergeSort(_ASCENDENTE, valores[:i+1])[0]
		require.Equal(t, esperadoMin, h.VerMax(), "El mínimo tras encolar %d debería ser %d", v, esperadoMin)
	}

	for _, esperado := range valoresEsperados {
		obtenido := h.Desencolar()
		require.Equal(t, esperado, obtenido, "Luego de desencolar se espera %d, y se obtuvo %d", esperado, obtenido)
	}

	require.True(t, h.EstaVacia(), "El heap debería estar vacío después de desencolar todo")
}

func TestCompararOrdenamientos(t *testing.T) {
	valores := []int{10, 3, 7, 1, 4, 9, 6, 2, 5, 8}

	esperado := MergeSort(_DESCENDENTE, valores)

	h1 := TDAHeap.CrearHeap(cmpMaxHeap)
	for _, v := range valores {
		h1.Encolar(v)
	}
	heapCreado := make([]int, 0, len(valores))
	for !h1.EstaVacia() {
		heapCreado = append(heapCreado, h1.Desencolar())
	}
	require.Equal(t, esperado, heapCreado, "CrearHeap manual debería generar mismo orden que MergeSort descendente")

	h2 := TDAHeap.CrearHeapArr(valores, cmpMaxHeap)
	obtenidoHeapArr := make([]int, 0, len(valores))
	for !h2.EstaVacia() {
		obtenidoHeapArr = append(obtenidoHeapArr, h2.Desencolar())
	}
	require.Equal(t, esperado, obtenidoHeapArr, "CrearHeapArr debería generar mismo orden que MergeSort descendente")
}
