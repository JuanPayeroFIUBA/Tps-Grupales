package lista_test

//xd

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const lista_vacia = "La lista esta vacia"
const panic_iteracion_completada = "El iterador termino de iterar"
const test_volumen_1 = 100000
const test_volumen_2 = 1000000
const test_mitad_volumen_1 = 50000
const test_mitad_volumen_2 = 500000

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

// estos dos de abajo no soon muy necesarios solo chequean que el verUltimo no rompa el test de volumen
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

// /////////////////////////
// Tests de Iterador Interno
// /////////////////////////

func TestIteradorInternoSoloUnaIteracion(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.Iterar(func(i int) bool {
		return false
	})
	require.True(t, lista.EstaVacia())
}
func TestVolumenConIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.Iterar(func(i int) bool {
		if i < test_volumen_1 {
			lista.InsertarPrimero(i)
			i++
			return true
		}
		return false
	})
	lista.Iterar(func(i int) bool {
		if i > 0 {
			lista.BorrarPrimero()
			i--
			return true
		}
		return false
	})
	require.True(t, lista.EstaVacia())
}

// /////////////////////////
// Tests de Iterador Externo
// /////////////////////////
func TestIteradorExternoEnListaRecienCreada(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	require.PanicsWithValue(t, panic_iteracion_completada, func() { iter.Siguiente() })
}

func TestIteradorExternoRecienCreado(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(5)
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(999)
	iter := lista.Iterador()
	require.EqualValues(t, 999, iter.VerActual())
}

func TestIteradorExternoInsertarEnPosicionEnLaQueSeCreaElIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	iter.Insertar(60)
	require.EqualValues(t, 60, lista.VerPrimero())
	require.EqualValues(t, 60, iter.VerActual())

}
func TestIteradorExternoInsertarEnPosicionEnLaQueSeCreaElIteradorConListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iter := lista.Iterador()
	lista.InsertarPrimero(6)
	lista.InsertarPrimero(7)
	lista.InsertarPrimero(999)
	iter.Insertar(60)
	require.EqualValues(t, 60, lista.VerPrimero())
	require.EqualValues(t, 60, iter.VerActual())

}
