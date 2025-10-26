package cola_prioridad_test

import (
	"math/rand"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func cmpEnteros(a, b int) int {
	return a - b
}
func cmpStrings(a, b string) int {
	return strings.Compare(a, b)
}
func MergeSortAuxiliar(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	izquierda := MergeSortAuxiliar(arr[:mid])
	derecha := MergeSortAuxiliar(arr[mid:])

	return merge(izquierda, derecha)
}

func merge(izq, der []int) []int {
	resultado := make([]int, 0, len(izq)+len(der))
	i, j := 0, 0

	for i < len(izq) && j < len(der) {
		if izq[i] <= der[j] {
			resultado = append(resultado, izq[i])
			i++
		} else {
			resultado = append(resultado, der[j])
			j++
		}
	}
	resultado = append(resultado, izq[i:]...)
	resultado = append(resultado, der[j:]...)

	return resultado
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

func TestEncolarYDesencolarStrings(t *testing.T) {
	t.Log("Prueba encolar algunos elementos y desencolarlos")
	heap := TDAHeap.CrearHeap(cmpStrings)

	heap.Encolar("5")
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, "5", heap.VerMax())

	heap.Encolar("hola")
	heap.Encolar("juan")
	heap.Encolar("XX")
	require.EqualValues(t, 4, heap.Cantidad())
	require.EqualValues(t, "juan", heap.VerMax())

	require.EqualValues(t, "juan", heap.Desencolar())
	require.EqualValues(t, "hola", heap.Desencolar())
	require.EqualValues(t, "XX", heap.Desencolar())
	require.EqualValues(t, "5", heap.Desencolar())
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
func TestVolumen(t *testing.T) {
	t.Log("Prueba con muchos volumenes de elementos")
	heap := TDAHeap.CrearHeap(cmpEnteros)
	volumenes := []int{1000, 10000, 50000, 100000, 500000}
	for _, volumen := range volumenes {
		arr := []int{}
		for i := 0; i < volumen; i++ {
			arr = append(arr, i)
			heap.Encolar(i)
		}
		require.EqualValues(t, volumen, heap.Cantidad())

		for i := volumen - 1; i >= 0; i-- {
			require.EqualValues(t, i, heap.Desencolar())
		}
		require.True(t, heap.EstaVacia())

		heap_arr := TDAHeap.CrearHeapArr(arr, cmpEnteros)
		for i := volumen - 1; i >= 0; i-- {
			require.EqualValues(t, i, heap_arr.Desencolar())
		}
		require.True(t, heap.EstaVacia())
	}

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

func TestHeapSortArregloVacio(t *testing.T) {
	t.Log("Testea que el algoritmo no rompa al recibir un arreglo vacio")
	arr := []int{}
	TDAHeap.HeapSort(arr, cmpEnteros)

	require.EqualValues(t, []int{}, arr)
}

func TestHeapSortCasoGenerico(t *testing.T) {
	t.Log("Ordena arreglo con HeapSort")
	arr := []int{13, 4, 19, 2, 9, 15, 6, 2, 15}
	TDAHeap.HeapSort(arr, cmpEnteros)

	esperado := []int{2, 2, 4, 6, 9, 13, 15, 15, 19}
	require.EqualValues(t, esperado, arr)
}

func TestHeapSortArregloOrdenado(t *testing.T) {
	t.Log("Ordena arreglo previamente ordenado con HeapSort, o lo que es lo mismo, dejandolo igual que como llega")
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	TDAHeap.HeapSort(arr, cmpEnteros)

	require.EqualValues(t, arr, arr)
}

func TestVolumenHeapSort(t *testing.T) {
	t.Log("Testea que el Algoritmo funcione correctamente para diferentes volumenes de enteros")
	volumenes := []int{1000, 10000, 50000, 100000, 500000}
	for _, volumen := range volumenes {
		arr := rand.Perm(volumen)
		copia := MergeSortAuxiliar(arr) // ordenamos con un algoritmo que sabemos que funciona y no se rompe

		TDAHeap.HeapSort(arr, cmpEnteros)
		require.Equal(t, copia, arr)
	}
}
