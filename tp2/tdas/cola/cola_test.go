package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

const cola_vacia = "La cola esta vacia"
const test_volumen_1 = 100000
const test_volumen_2 = 1000000
const test_volumen_3 = 5000000

func TestColaRecienCreada(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
}
func TestColaEstaVaciaCasoFalso(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.False(t, cola.EstaVacia())
}

func TestColaVerPrimeroCasoColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, cola_vacia, func() { cola.VerPrimero() })
}

func TestColaEncolarUnElementoYVerPrimero(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	require.EqualValues(t, 5, cola.VerPrimero())
}

func TestColaEncolarMuchosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	cola.Encolar(6)
	cola.Encolar(7)
	cola.Encolar(58)
	cola.Encolar(9)
	cola.Encolar(5)
	require.EqualValues(t, 5, cola.VerPrimero())
}

func TestColaEncolarMuchosElementosDeRune(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[rune]()
	cola.Encolar('i')
	cola.Encolar('s')
	cola.Encolar('+')
	cola.Encolar(')')
	cola.Encolar('"')
	cola.Encolar('|')
	require.EqualValues(t, 'i', cola.VerPrimero())
}

func TestColaEncolarMuchosElementosDeString(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("sadad")
	cola.Encolar("TDA Cola")
	cola.Encolar("hola")
	cola.Encolar("2222")
	cola.Encolar("esto es un string")
	require.EqualValues(t, "sadad", cola.VerPrimero())
}

func TestColaRecienCreadaDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, cola_vacia, func() { cola.Desencolar() })
}
func TestColaVaciadaDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	cola.Encolar(5)
	cola.Desencolar()
	cola.Desencolar()
	require.PanicsWithValue(t, cola_vacia, func() { cola.Desencolar() })
}

func TestColaDesencolarUnElemento(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(5)
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}
func TestColaEncolarDespuesDeDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	cola.Encolar(5)
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	cola.Encolar(58)
	cola.Encolar(9)
	require.EqualValues(t, 58, cola.VerPrimero())
}

func TestColaDesencolarMuchosElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	cola.Encolar(5)
	require.EqualValues(t, 10, cola.VerPrimero())
	cola.Encolar(6)
	cola.Encolar(7)
	cola.Desencolar()
	cola.Desencolar()
	require.EqualValues(t, 6, cola.VerPrimero())
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
	cola.Encolar(58)
	cola.Encolar(9)
	require.EqualValues(t, 58, cola.VerPrimero())
}

func TestColaDesencolarMuchosElementosString(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[string]()
	cola.Encolar("sadad")
	cola.Encolar("TDA Cola")
	require.EqualValues(t, "sadad", cola.VerPrimero())
	cola.Encolar("hola")
	cola.Encolar("2222")
	cola.Desencolar()
	cola.Desencolar()
	require.EqualValues(t, "hola", cola.VerPrimero())
	cola.Desencolar()
	cola.Desencolar()
	require.True(t, cola.EstaVacia())
}

func TestFIFOCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(10)
	require.EqualValues(t, 10, cola.VerPrimero())
	cola.Encolar(5)
	cola.Desencolar()
	require.EqualValues(t, 5, cola.VerPrimero())
	cola.Encolar(6)
	cola.Desencolar()
	require.EqualValues(t, 6, cola.VerPrimero())
	cola.Encolar(7)
	cola.Desencolar()
	require.EqualValues(t, 7, cola.VerPrimero())
}

func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for contador := 0; contador <= test_volumen_1; contador++ {
		cola.Encolar(contador)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for contador := test_volumen_1; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_1-contador, cola.VerPrimero())
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}
func TestVolumenGrande(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for contador := 0; contador <= test_volumen_2; contador++ {
		cola.Encolar(contador)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for contador := test_volumen_2; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_2-contador, cola.VerPrimero())
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}
func TestVolumenMuyGrande(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()

	for contador := 0; contador <= test_volumen_3; contador++ {
		cola.Encolar(contador)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for contador := test_volumen_3; contador >= 0; contador-- {
		require.EqualValues(t, test_volumen_3-contador, cola.VerPrimero())
		cola.Desencolar()
	}
	require.True(t, cola.EstaVacia())
}
