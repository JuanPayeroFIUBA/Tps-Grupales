package lista_test

import (
	"testing"

	TDALista "tdas/lista"

	"github.com/stretchr/testify/require"
)

const (
	msj_lista1_vacia       = "La lista esta vacia"
	msj_iterador_terminado = "El iterador termino de iterar"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, msj_lista1_vacia, func() { lista.VerPrimero() })
	require.PanicsWithValue(t, msj_lista1_vacia, func() { lista.VerUltimo() })
	require.PanicsWithValue(t, msj_lista1_vacia, func() { lista.BorrarPrimero() })
}

func TestInsertarPrimeroGeneraOrdenInverso(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	entrada := []int{1, 2, 3, 4}
	for _, elem := range entrada {
		lista.InsertarPrimero(elem)
		require.Equal(t, elem, lista.VerPrimero())
	}
	require.Equal(t, len(entrada), lista.Largo())
	require.Equal(t, []int{4, 3, 2, 1}, sliceDesdeLista(lista))
	require.Equal(t, 4, lista.VerPrimero())
	require.Equal(t, 1, lista.VerUltimo())
}

func TestInsertarUltimoMantieneOrden(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	entrada := []string{"a", "b", "c"}
	for _, elem := range entrada {
		lista.InsertarUltimo(elem)
		require.Equal(t, elem, lista.VerUltimo())
	}
	require.Equal(t, entrada, sliceDesdeLista(lista))
	require.Equal(t, "a", lista.VerPrimero())
	require.Equal(t, "c", lista.VerUltimo())
}

func TestBorrarPrimeroHastaVaciar(t *testing.T) {
	lista := listaDesdeSlice([]int{1, 2, 3, 4})
	var salida []int
	for !lista.EstaVacia() {
		salida = append(salida, lista.BorrarPrimero())
	}
	require.Equal(t, []int{1, 2, 3, 4}, salida)
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, msj_lista1_vacia, func() { lista.BorrarPrimero() })
}

func TestOperacionesMixtasMantienenExtremos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for _, elem := range []int{0, 1, 2} {
		lista.InsertarUltimo(elem)
	}
	for _, elem := range []int{-1, -2} {
		lista.InsertarPrimero(elem)
	}
	require.Equal(t, []int{-2, -1, 0, 1, 2}, sliceDesdeLista(lista))
	require.Equal(t, -2, lista.BorrarPrimero())
	require.Equal(t, -1, lista.BorrarPrimero())
	require.Equal(t, []int{0, 1, 2}, sliceDesdeLista(lista))
	require.Equal(t, 0, lista.VerPrimero())
	require.Equal(t, 2, lista.VerUltimo())
	require.Equal(t, 3, lista.Largo())
}

func TestListaVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const cantidad = 50000
	for i := 0; i < cantidad; i++ {
		lista.InsertarUltimo(i)
	}
	require.Equal(t, cantidad, lista.Largo())
	for i := 0; i < cantidad; i++ {
		require.Equal(t, i, lista.BorrarPrimero())
	}
	require.True(t, lista.EstaVacia())
}

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.VerActual() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Siguiente() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Borrar() })
	iter.Insertar(5)
	require.Equal(t, []int{5}, sliceDesdeLista(lista))
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, 5, iter.VerActual())
}

func TestIteradorInsercionesDistintasPosiciones(t *testing.T) {
	lista := listaDesdeSlice([]int{10, 20, 30})
	iter := lista.Iterador()
	iter.Insertar(5)
	require.Equal(t, []int{5, 10, 20, 30}, sliceDesdeLista(lista))
	require.Equal(t, 5, iter.VerActual())

	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(15)
	require.Equal(t, []int{5, 10, 15, 20, 30}, sliceDesdeLista(lista))
	require.Equal(t, 15, iter.VerActual())

	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(35)
	require.Equal(t, []int{5, 10, 15, 20, 30, 35}, sliceDesdeLista(lista))
	require.Equal(t, 35, lista.VerUltimo())
	require.Equal(t, 6, lista.Largo())
}

func TestIteradorBorradosDistintasPosiciones(t *testing.T) {
	lista := listaDesdeSlice([]int{10, 20, 30, 40})
	iter := lista.Iterador()
	require.Equal(t, 10, iter.Borrar())
	require.Equal(t, []int{20, 30, 40}, sliceDesdeLista(lista))

	iter.Siguiente()
	require.Equal(t, 30, iter.Borrar())
	require.Equal(t, []int{20, 40}, sliceDesdeLista(lista))

	require.Equal(t, 40, iter.Borrar())
	require.Equal(t, []int{20}, sliceDesdeLista(lista))
	require.False(t, iter.HaySiguiente())
	require.Equal(t, 20, lista.VerPrimero())
	require.Equal(t, 20, lista.VerUltimo())
	require.Equal(t, 1, lista.Largo())
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Borrar() })
}

func TestIteradorRecorridoCompleto(t *testing.T) {
	lista := listaDesdeSlice([]string{"uno", "dos", "tres"})
	iter := lista.Iterador()
	recorrido := []string{}
	for iter.HaySiguiente() {
		recorrido = append(recorrido, iter.VerActual())
		iter.Siguiente()
	}
	require.Equal(t, []string{"uno", "dos", "tres"}, recorrido)
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.VerActual() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Siguiente() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Borrar() })
}

func TestIteradorVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const cantidad = 20000
	for i := 0; i < cantidad; i++ {
		lista.InsertarUltimo(i)
	}
	iter := lista.Iterador()
	for i := 0; i < cantidad; i++ {
		require.True(t, iter.HaySiguiente())
		require.Equal(t, i, iter.VerActual())
		iter.Siguiente()
	}
	require.False(t, iter.HaySiguiente())
}

func TestIteradorInternoSobreListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	llamado := false
	lista.Iterar(func(elem int) bool {
		llamado = true
		return true
	})
	require.False(t, llamado)
}

func TestIteradorInternoRecorridoCompleto(t *testing.T) {
	lista := listaDesdeSlice([]int{2, 4, 6})
	recorrido := []int{}
	lista.Iterar(func(elem int) bool {
		recorrido = append(recorrido, elem)
		return true
	})
	require.Equal(t, []int{2, 4, 6}, recorrido)
}

func TestIteradorInternoConCorte(t *testing.T) {
	lista := listaDesdeSlice([]int{2, 4, 6, 8})
	recorrido := []int{}
	lista.Iterar(func(elem int) bool {
		recorrido = append(recorrido, elem)
		return elem != 6
	})
	require.Equal(t, []int{2, 4, 6}, recorrido)
}

func TestIteradorInternoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	const cantidad = 30000
	for i := 0; i < cantidad; i++ {
		lista.InsertarUltimo(i)
	}
	contador := 0
	lista.Iterar(func(elem int) bool {
		require.Equal(t, contador, elem)
		contador++
		return true
	})
	require.Equal(t, cantidad, contador)
}

func listaDesdeSlice[T any](elementos []T) TDALista.Lista[T] {
	lista := TDALista.CrearListaEnlazada[T]()
	for _, elem := range elementos {
		lista.InsertarUltimo(elem)
	}
	return lista
}

func sliceDesdeLista[T any](lista TDALista.Lista[T]) []T {
	resultado := []T{}
	lista.Iterar(func(elem T) bool {
		resultado = append(resultado, elem)
		return true
	})
	return resultado
}
