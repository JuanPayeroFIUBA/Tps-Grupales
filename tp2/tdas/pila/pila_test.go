package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

const pila_vacia = "La pila esta vacia"
const test_volumen_1 = 100000
const test_volumen_2 = 1000000
const test_volumen_3 = 5000000

func TestPilaEstaVaciaCasoVerdadero(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
}

func TestPilaEstaVaciaCasoFalso(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.False(t, pila.EstaVacia())
}

func TestPilaVaciadaVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	pila.Desapilar()
	require.PanicsWithValue(t, pila_vacia, func() { pila.VerTope() })
}
func TestPilaVaciadaDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	pila.Desapilar()
	require.PanicsWithValue(t, pila_vacia, func() { pila.Desapilar() })
}
func TestPilaRecienCreadaVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, pila_vacia, func() { pila.VerTope() })

}
func TestPilaRecienCreadaDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, pila_vacia, func() { pila.Desapilar() })
}

func TestApilarUnElemento(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())

}

func TestApilarMuchosElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(99)
	require.EqualValues(t, 99, pila.VerTope())
}
func TestApilarMuchosElementosStrings(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("hola")
	pila.Apilar("ola")
	pila.Apilar("la")
	pila.Apilar("al")
	pila.Apilar("alo")
	pila.Apilar("un numero")
	require.EqualValues(t, "un numero", pila.VerTope())
}

func TestApilarMuchosElementosFloat(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[float64]()
	pila.Apilar(1440.544)
	pila.Apilar(10.4524)
	pila.Apilar(121241.54124)
	pila.Apilar(142.555)
	pila.Apilar(1452.5456)
	pila.Apilar(10.56)
	require.EqualValues(t, 10.56, pila.VerTope())
}

func TestLIFOPila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(99)
	pila.Desapilar()
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(6)
	pila.Desapilar()
	pila.Desapilar()
	require.EqualValues(t, 1, pila.VerTope())
}

func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for contador := 0; contador <= test_volumen_1; contador++ {
		pila.Apilar(contador)
		require.EqualValues(t, contador, pila.VerTope())
	}
	for contador := test_volumen_1; contador >= 0; contador-- {
		require.EqualValues(t, contador, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestVolumenGrande(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for contador := 0; contador <= test_volumen_2; contador++ {
		pila.Apilar(contador)
		require.EqualValues(t, contador, pila.VerTope())
	}
	for contador := test_volumen_2; contador >= 0; contador-- {
		require.EqualValues(t, contador, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}

func TestVolumeMuyGrande(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()

	for contador := 0; contador <= test_volumen_3; contador++ {
		pila.Apilar(contador)
		require.EqualValues(t, contador, pila.VerTope())
	}
	for contador := test_volumen_3; contador >= 0; contador-- {
		require.EqualValues(t, contador, pila.VerTope())
		pila.Desapilar()
	}
	require.True(t, pila.EstaVacia())
}
