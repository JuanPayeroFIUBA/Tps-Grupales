package cola_prioridad_test

import (
	TDAHeap "tdas/heap"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmpEnteros(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	t.Log("Verifica comportamiento de heap recién creado")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarYDesencolar(t *testing.T) {
	t.Log("Prueba encolar algunos elementos y desencolarlos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	
	heap.Encolar(5)
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 5, heap.VerMax())
	
	heap.Encolar(3)
	heap.Encolar(8)
	heap.Encolar(1)
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, 8, heap.VerMax())
	
	require.EqualValues(t, 8, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 3, heap.Desencolar())
	require.EqualValues(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestCrearHeapDesdeArreglo(t *testing.T) {
	t.Log("Crea heap desde arreglo y verifica orden de desencolado")
	valores := []int{15, 3, 9, 22, 7, 11, 5}
	heap := TDAHeap.CrearHeapArr(valores, cmpEnteros)
	
	require.EqualValues(t, 7, heap.Cantidad())
	require.EqualValues(t, 22, heap.VerMax())
	
	esperados := []int{22, 15, 11, 9, 7, 5, 3}
	for _, esperado := range esperados {
		require.EqualValues(t, esperado, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	t.Log("Ordena arreglo con HeapSort")
	arr := []int{13, 4, 19, 2, 9, 15, 6}
	TDAHeap.HeapSort(arr, cmpEnteros)
	
	esperado := []int{2, 4, 6, 9, 13, 15, 19}
	require.EqualValues(t, esperado, arr)
}

func TestVolumen(t *testing.T) {
	t.Log("Prueba con muchos elementos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	
	for i := 0; i < 1000; i++ {
		heap.Encolar(i)
	}
	require.EqualValues(t, 1000, heap.Cantidad())
	
	for i := 999; i >= 0; i-- {
		require.EqualValues(t, i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestInvarianteHeap(t *testing.T) {
	t.Log("Verifica que siempre se desencola el maximo")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	numeros := []int{25, 10, 30, 5, 15, 35, 20}
	
	for _, num := range numeros {
		heap.Encolar(num)
	}
	
	anterior := heap.Desencolar()
	for !heap.EstaVacia() {
		actual := heap.Desencolar()
		require.True(t, anterior >= actual)
		anterior = actual
	}
}
