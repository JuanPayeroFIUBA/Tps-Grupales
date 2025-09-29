package lista_test

//xd

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const lista_vacia = "La lista esta vacia"
const test_volumen_1 = 100000
const test_volumen_2 = 1000000
const test_mitad_volumen_1 = 50000
const test_mitad_volumen_2 = 500000
const test_volumen_iterador = 1000
const test_volumen_iterador_interno = 10000

// ///////////////////////
// Tests Genericos de TDAS
// ///////////////////////
func TestListaRecienCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
}
func TestListaEstaVaciaCasoFalso(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.False(t, lista.EstaVacia())
}

func TestListaVerPrimeroCasolistaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, lista_vacia, func() { lista.VerPrimero() })
}

func TestListaRecienCreadaBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, lista_vacia, func() { lista.BorrarPrimero() })
}

func TestListaInsertarUnElementoAlPrincipioYVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	require.EqualValues(t, 5, lista.VerPrimero())
}
func TestListaInsertarUnElementoAlFinalYVerPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	require.EqualValues(t, 5, lista.VerPrimero())
}

// //////////////////////
// Tests Largo de la Pila
// //////////////////////

func TestListaRecienCreadaLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaLargoInsertarUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	require.EqualValues(t, 1, lista.Largo())
}

func TestListaLargoInsertarVariosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(5)
	require.EqualValues(t, 6, lista.Largo())
}

// ////////////////////////////////////////////////////////////////////
// Tests de Insercion (al principio y final por separado) y Eliminacion
// ////////////////////////////////////////////////////////////////////

func TestListaInsertarPrimeroMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(999)
	require.EqualValues(t, 999, lista.VerPrimero())
}

func TestListaInsertarElementosDeRune(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[rune]()
	lista.InsertarPrimero('i')
	lista.InsertarPrimero('s')
	lista.InsertarPrimero('+')
	lista.InsertarPrimero('|')
	require.EqualValues(t, '|', lista.VerPrimero())
}

func TestListaInsertarElementosDeString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("sadad")
	lista.InsertarPrimero("TDA Lista")
	lista.InsertarPrimero("2222")
	lista.InsertarPrimero("ultimo_primer_valor")
	require.EqualValues(t, "ultimo_primer_valor", lista.VerPrimero())
}

func TestListaBorrarPrimeroUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
}
func TestListaVaciadaBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(5)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.PanicsWithValue(t, lista_vacia, func() { lista.BorrarPrimero() })
}

func TestListaInsertarPrimeroDespuesDeBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(5)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(58)
	lista.InsertarPrimero(9)
	require.EqualValues(t, 9, lista.VerPrimero())
}

func TestListaBorrarPrimeroMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	lista.InsertarPrimero(5)
	require.EqualValues(t, 5, lista.VerPrimero())
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(7)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.EqualValues(t, 5, lista.VerPrimero())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	lista.InsertarPrimero(58)
	lista.InsertarPrimero(9)
	require.EqualValues(t, 9, lista.VerPrimero())
}

func TestListaBorrarPrimeroMuchosElementosString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("sadad")
	lista.InsertarPrimero("TDA lista")
	require.EqualValues(t, "TDA lista", lista.VerPrimero())
	lista.InsertarPrimero("hola")
	lista.InsertarPrimero("2222")
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.EqualValues(t, "TDA lista", lista.VerPrimero())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
}

func TestListaLIFOInsertandoYBorrandoPrimeros(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.EqualValues(t, 10, lista.VerPrimero())
	lista.InsertarPrimero(5)
	lista.BorrarPrimero()
	require.EqualValues(t, 10, lista.VerPrimero())
	lista.InsertarPrimero(6)
	lista.BorrarPrimero()
	require.EqualValues(t, 10, lista.VerPrimero())
	lista.InsertarPrimero(7)
	lista.BorrarPrimero()
	require.EqualValues(t, 10, lista.VerPrimero())
}

func TestVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for contador := 0; contador <= test_volumen_1; contador++ {
		lista.InsertarPrimero(contador)
		require.EqualValues(t, contador, lista.VerPrimero())
	}
	for contador := test_volumen_1; contador >= 0; contador-- {
		require.EqualValues(t, contador, lista.VerPrimero())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}
func TestVolumenMasGrande(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for contador := 0; contador <= test_volumen_2; contador++ {
		lista.InsertarPrimero(contador)
		require.EqualValues(t, contador, lista.VerPrimero())
	}
	for contador := test_volumen_2; contador >= 0; contador-- {
		require.EqualValues(t, contador, lista.VerPrimero())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

func TestListaInsertarUltimoMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	lista.InsertarUltimo(999)
	require.EqualValues(t, 5, lista.VerPrimero())
}

func TestListaInsertarUltimoDespuesDeBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(5)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(58)
	lista.InsertarUltimo(9)
	require.EqualValues(t, 58, lista.VerPrimero())
}

func TestListaBorrarPrimeroMuchosElementosInsertadosUltimos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(5)
	require.EqualValues(t, 10, lista.VerPrimero())
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(7)
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.EqualValues(t, 6, lista.VerPrimero())
	lista.BorrarPrimero()
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	lista.InsertarUltimo(58)
	lista.InsertarUltimo(9)
	require.EqualValues(t, 58, lista.VerPrimero())
}

func TestListaFIFOInsertandoUltimosYBorrandoPrimeros(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(10)
	require.EqualValues(t, 10, lista.VerPrimero())
	lista.InsertarUltimo(5)
	lista.BorrarPrimero()
	require.EqualValues(t, 5, lista.VerPrimero())
	lista.InsertarUltimo(6)
	lista.BorrarPrimero()
	require.EqualValues(t, 6, lista.VerPrimero())
	lista.InsertarUltimo(7)
	lista.BorrarPrimero()
	require.EqualValues(t, 7, lista.VerPrimero())
}

func TestVolumenInsertandoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for contador := 0; contador <= test_volumen_1; contador++ {
		lista.InsertarUltimo(contador)
		require.EqualValues(t, 0, lista.VerPrimero())
	}
	require.EqualValues(t, test_volumen_1+1, lista.Largo())
	for contador := test_volumen_1; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_1-contador, lista.VerPrimero())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}
func TestVolumenMasGrandeInsertandoUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for contador := 0; contador <= test_volumen_2; contador++ {
		lista.InsertarUltimo(contador)
		require.EqualValues(t, 0, lista.VerPrimero())
	}
	require.EqualValues(t, test_volumen_2+1, lista.Largo())
	for contador := test_volumen_2; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_2-contador, lista.VerPrimero())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

// //////////////////////////////////////////////////////////////////
// Tests de Insercion (al principio y final combinados) y Eliminacion
// //////////////////////////////////////////////////////////////////

func TestListaInsertarPrimeroYUltimoMuchosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(10)
	require.EqualValues(t, 10, lista.VerPrimero())
	lista.InsertarUltimo(5)
	lista.BorrarPrimero()
	require.EqualValues(t, 5, lista.VerPrimero())
	lista.InsertarPrimero(6)
	lista.BorrarPrimero()
	require.EqualValues(t, 5, lista.VerPrimero())
	lista.InsertarUltimo(7)
	lista.BorrarPrimero()
	require.EqualValues(t, 7, lista.VerPrimero())
}

func TestVolumenInsertarPrimeroYUltimoYViendoElUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for contador := 0; contador <= test_mitad_volumen_1; contador++ {
		lista.InsertarPrimero(contador)
		require.EqualValues(t, 0, lista.VerUltimo())
	}
	for contador := test_mitad_volumen_1; contador <= test_volumen_1; contador++ {
		lista.InsertarUltimo(contador)
		require.EqualValues(t, contador, lista.VerUltimo())
	}
	require.EqualValues(t, test_volumen_1+2, lista.Largo())
	for contador := test_volumen_1 + 1; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_1, lista.VerUltimo())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}
func TestVolumenMasGrandeInsertarPrimeroYUltimoYViendoElPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()

	for contador := 0; contador <= test_mitad_volumen_2; contador++ {
		lista.InsertarPrimero(contador)
		require.EqualValues(t, contador, lista.VerPrimero())
	}
	require.EqualValues(t, test_mitad_volumen_2+1, lista.Largo())
	for contador := test_mitad_volumen_2; contador >= 0; contador-- {
		require.EqualValues(t, contador, lista.VerPrimero())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	for contador := test_mitad_volumen_2; contador <= test_volumen_2; contador++ {
		lista.InsertarUltimo(contador)
		require.EqualValues(t, test_mitad_volumen_2, lista.VerPrimero())
	}
	require.EqualValues(t, test_mitad_volumen_2+1, lista.Largo())
	for contador := test_mitad_volumen_2; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_2-contador, lista.VerPrimero())
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
}

// /////////////////////////////////////
// Tests del Iterador Externo
// /////////////////////////////////////

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.False(t, iter.HaySiguiente())
}

func TestIteradorVerActualListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorSiguienteListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorBorrarListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

func TestIteradorInsertarEnListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(26)
	require.EqualValues(t, 26, lista.VerPrimero())
	require.EqualValues(t, 26, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
}

func TestIteradorInsertarAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(06)
	lista.InsertarPrimero(11)
	iter := lista.Iterador()
	iter.Insertar(26)
	require.EqualValues(t, 26, lista.VerPrimero())
	require.EqualValues(t, 06, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	lista.InsertarPrimero(06)
	iter := lista.Iterador()
	// Avanzar al final
	iter.Siguiente()
	iter.Siguiente()
	iter.Insertar(5)
	require.EqualValues(t, 06, lista.VerPrimero())
	require.EqualValues(t, 5, lista.VerUltimo())
	require.EqualValues(t, 3, lista.Largo())
}

func TestIteradorInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	lista.InsertarPrimero(11)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Insertar(06)
	require.EqualValues(t, 11, lista.VerPrimero())
	require.EqualValues(t, 26, lista.VerUltimo())
	// Verificar orden 30, 20, 10
	iter2 := lista.Iterador()
	require.EqualValues(t, 11, iter2.VerActual())
	iter2.Siguiente()
	require.EqualValues(t, 06, iter2.VerActual())
	iter2.Siguiente()
	require.EqualValues(t, 26, iter2.VerActual())
}

func TestIteradorBorrarPrimerElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	lista.InsertarPrimero(06)
	lista.InsertarPrimero(11)
	iter := lista.Iterador()
	eliminado := iter.Borrar()
	require.EqualValues(t, 11, eliminado)
	require.EqualValues(t, 06, lista.VerPrimero())
	require.EqualValues(t, 2, lista.Largo())
}

func TestIteradorBorrarUltimoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	lista.InsertarPrimero(06)
	lista.InsertarPrimero(11)
	iter := lista.Iterador()
	iter.Siguiente()
	iter.Siguiente()
	eliminado := iter.Borrar()
	require.EqualValues(t, 26, eliminado)
	require.EqualValues(t, 11, lista.VerPrimero())
	require.EqualValues(t, 06, lista.VerUltimo())
	require.EqualValues(t, 2, lista.Largo())
}

func TestIteradorBorrarElementoDelMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	lista.InsertarPrimero(06)
	lista.InsertarPrimero(11)
	iter := lista.Iterador()
	iter.Siguiente()
	eliminado := iter.Borrar()
	require.EqualValues(t, 06, eliminado)
	require.EqualValues(t, 11, lista.VerPrimero())
	require.EqualValues(t, 26, lista.VerUltimo())
	require.EqualValues(t, 2, lista.Largo())
	// Verificar que el elemento no esta
	iter2 := lista.Iterador()
	require.EqualValues(t, 11, iter2.VerActual())
	iter2.Siguiente()
	require.EqualValues(t, 26, iter2.VerActual())
}

func TestIteradorRecorrerCompleto(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	elementos := []int{26, 06, 11}
	for _, elem := range elementos {
		lista.InsertarPrimero(elem)
	}

	elementosEsperados := []int{11, 06, 26}
	
	iter := lista.Iterador()
	indice := 0
	for iter.HaySiguiente() {
		require.EqualValues(t, elementosEsperados[indice], iter.VerActual())
		iter.Siguiente()
		indice++
	}
	require.EqualValues(t, len(elementosEsperados), indice)
}

func TestIteradorBorrarTodosLosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	lista.InsertarPrimero(06)
	lista.InsertarPrimero(11)
	
	iter := lista.Iterador()
	eliminados := []int{}
	for iter.HaySiguiente() {
		eliminados = append(eliminados, iter.Borrar())
	}
	
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	require.Equal(t, []int{11, 06, 26}, eliminados)
}

func TestIteradorVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < test_volumen_iterador; i++ {
		lista.InsertarUltimo(i)
	}
	
	iter := lista.Iterador()
	contador := 0
	for iter.HaySiguiente() {
		require.EqualValues(t, contador, iter.VerActual())
		iter.Siguiente()
		contador++
	}
	require.EqualValues(t, test_volumen_iterador, contador)
}

// /////////////////////////////////////
// Tests del Iterador Interno
// /////////////////////////////////////

func TestIteradorInternoListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	ejecutado := false
	lista.Iterar(func(elem int) bool {
		ejecutado = true
		return true
	})
	require.False(t, ejecutado)
}

func TestIteradorInternoUnElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(26)
	elementos := []int{}
	lista.Iterar(func(elem int) bool {
		elementos = append(elementos, elem)
		return true
	})
	require.Equal(t, []int{26}, elementos)
}

func TestIteradorInternoVariosElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(26)
	lista.InsertarUltimo(06)
	lista.InsertarUltimo(11)
	
	elementos := []int{}
	lista.Iterar(func(elem int) bool {
		elementos = append(elementos, elem)
		return true
	})
	require.Equal(t, []int{26, 06, 11}, elementos)
}

func TestIteradorInternoCorteAlPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(26)
	lista.InsertarUltimo(06)
	lista.InsertarUltimo(11)
	
	elementos := []int{}
	lista.Iterar(func(elem int) bool {
		elementos = append(elementos, elem)
		return false
	})
	require.Equal(t, []int{26}, elementos)
}

func TestIteradorInternoCorteEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(26)
	lista.InsertarUltimo(06)
	lista.InsertarUltimo(11)
	lista.InsertarUltimo(12)
	
	elementos := []int{}
	lista.Iterar(func(elem int) bool {
		elementos = append(elementos, elem)
		return elem != 06
	})
	require.Equal(t, []int{26, 06}, elementos)
}

func TestIteradorInternoBuscarElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarUltimo("Palermo")
	lista.InsertarUltimo("Bianchi")
	lista.InsertarUltimo("Schelotto")
	
	encontrado := false
	lista.Iterar(func(elem string) bool {
		if elem == "Bianchi" {
			encontrado = true
			return false
		}
		return true
	})
	require.True(t, encontrado)
}

func TestIteradorInternoVolumen(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i < test_volumen_iterador_interno; i++ {
		lista.InsertarUltimo(i)
	}
	
	contador := 0
	lista.Iterar(func(elem int) bool {
		require.EqualValues(t, contador, elem)
		contador++
		return true
	})
	require.EqualValues(t, test_volumen_iterador_interno, contador)
}
