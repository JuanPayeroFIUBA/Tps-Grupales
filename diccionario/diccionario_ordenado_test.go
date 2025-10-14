package diccionario_test

import (
	"fmt"
	"strings"
	"testing"

	TDADiccionario "tdas/diccionario"

	"github.com/stretchr/testify/require"
)

const (
	msj_abb_vacio          = "La lista esta vacia"
	msj_iterador_terminado = "El iterador termino de iterar"
	msj_clave_no_pertenece = "La clave no pertenece al diccionario"
)

func igualdadIntsABB(n1, n2 int) int {
	if n1 > n2 {
		return 1
	}
	if n1 == n2 {
		return 0
	}
	return -1
}
func igualdadStringsABB(n1, n2 string) int {
	resul := strings.Compare(n1, n2)
	if resul > 0 {
		return 1
	}
	if resul == 0 {
		return 0
	}
	return -1
}

//////////////////////
//TESTS DE DICCIONARIO
//////////////////////

func TestDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(1))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Obtener(1) })
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Borrar(1) })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Obtener("") })
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Borrar("") })

	dicNum := TDADiccionario.CrearABB[int, string](igualdadIntsABB)
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { dicNum.Borrar(0) })
}

func TestUnElementoDicciOrdenado(t *testing.T) {
	t.Log("Comprueba que el Diccionario ordenado con un elemento tenga esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](igualdadStringsABB)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Obtener("B") })
}

func TestDicciOrdenadoGuardar(t *testing.T) {
	t.Log("Comprueba que el diccionario ordenado pueda guardar varias claves con diferentes tipos")
	abb1 := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	claves_str := []string{"Riquelme", "Palermo", "Bianchi"}
	valores_str := []string{"10", "9", "DT"}

	for i, clave := range claves_str {
		require.False(t, abb1.Pertenece(clave))
		abb1.Guardar(clave, valores_str[i])
		require.EqualValues(t, i+1, abb1.Cantidad())
		require.True(t, abb1.Pertenece(clave))
		require.EqualValues(t, valores_str[i], abb1.Obtener(clave))
	}

	abb2 := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	claves_int := []int{9, 22, 26}
	for i, clave := range claves_int {
		abb2.Guardar(clave, i)
		require.True(t, abb2.Pertenece(clave))
		require.EqualValues(t, i, abb2.Obtener(clave))
	}
	require.EqualValues(t, 3, abb2.Cantidad())
}

func TestReemplazoDatoDiccionarioOrdenado(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestDiccionarioOrdenadoBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Obtener(claves[0]) })
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, msj_clave_no_pertenece, func() { abb.Obtener(claves[1]) })
}

func TestClaveVaciaDiccionarioOrdenado(t *testing.T) {
	t.Log("Guardamos una clave vacía y deberia funcionar sin problemas")
	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestValorNuloDiccionarioOrdenado(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDADiccionario.CrearABB[string, *int](igualdadStringsABB)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

// //////////////////////
// TESTS ITERADOR INTERNO
// //////////////////////
func TestClavesIteradorInternoDeDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, *int](igualdadStringsABB)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscar(cs[0], claves))
	require.NotEqualValues(t, -1, buscar(cs[1], claves))
	require.NotEqualValues(t, -1, buscar(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestValoresIteradorInternoDeDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](igualdadStringsABB)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestValoresConBorradosIteradorInternoDeDiccionarioOrdenado(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno, sin recorrer datos borrados")
	clave0 := "Elefante"
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](igualdadStringsABB)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

// //////////////////////
// TESTS ITERADOR EXTERNO
// //////////////////////

func TestIterarDiccionarioOrdenadoVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	abb := TDADiccionario.CrearABB[string, int](igualdadStringsABB)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.VerActual() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Siguiente() })
}

func TestDiccionarioOrdenadoIterar(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.EqualValues(t, valores[buscar(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscar(tercero, claves))
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.VerActual() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Siguiente() })
}

func TestIteradorDiccionarioOrdenadoNoLlegaAlFinal(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	claves := []string{"A", "B", "C"}
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	abb.Iterador()
	iter2 := abb.Iterador()
	iter2.Siguiente()
	iter3 := abb.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscar(primero, claves))
	require.NotEqualValues(t, -1, buscar(segundo, claves))
	require.NotEqualValues(t, -1, buscar(tercero, claves))
}

func TestIteradorDiccionarioOrdenadoIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	abb := TDADiccionario.CrearABB[string, string](igualdadStringsABB)
	abb.Guardar(clave1, "")
	abb.Guardar(clave2, "")
	abb.Guardar(clave3, "")
	abb.Borrar(clave1)
	abb.Borrar(clave2)
	abb.Borrar(clave3)
	iter := abb.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.VerActual() })
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.Siguiente() })
	abb.Guardar(clave1, "A")
	iter = abb.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestABBIteradoresMantienenOrden(t *testing.T) {
	t.Log("Valida que iteradores interno y externo recorran elementos en orden con inserciones desordenadas")
	abb := TDADiccionario.CrearABB[int, string](igualdadIntsABB)
	claves := []int{12, 6, 26, 3, 9, 22, 31, 11, 77, 28}
	for _, clave := range claves {
		abb.Guardar(clave, fmt.Sprintf("%d", clave))
	}

	clave_anterior := -1
	abb.Iterar(func(clave int, dato string) bool {
		require.True(t, clave > clave_anterior)
		clave_anterior = clave
		return true
	})

	iter := abb.Iterador()
	clave_anterior = -1
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.True(t, clave > clave_anterior)
		clave_anterior = clave
		iter.Siguiente()
	}
}

func TestOrdenConBorradosABB(t *testing.T) {
	t.Log("Valida que despues de borrar elementos el orden se mantenga")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	claves := []int{26, 12, 31, 6, 77, 28, 39, 3, 9, 22, 29}
	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	abb.Borrar(12)
	abb.Borrar(31)
	abb.Borrar(26)

	claves_iteradas := []int{}
	abb.Iterar(func(clave int, dato int) bool {
		claves_iteradas = append(claves_iteradas, clave)
		return true
	})

	for i := 1; i < len(claves_iteradas); i++ {
		require.True(t, claves_iteradas[i] > claves_iteradas[i-1],
			fmt.Sprintf("Clave %d en posicion %d no es mayor que %d en posicion %d",
				claves_iteradas[i], i, claves_iteradas[i-1], i-1))
	}
}

func TestIteradorInternoCorteABB(t *testing.T) {
	t.Log("Valida que el iterador interno se pare correctamente cuando retorna false")
	abb := TDADiccionario.CrearABB[int, string](igualdadIntsABB)
	claves := []int{3, 6, 9, 11, 12, 77, 22, 26, 28, 31}
	for _, clave := range claves {
		abb.Guardar(clave, fmt.Sprintf("%d", clave))
	}

	contador := 0
	abb.Iterar(func(clave int, dato string) bool {
		contador++
		return clave < 12
	})

	require.EqualValues(t, 5, contador)
}

func TestIterarRangoInternoABB(t *testing.T) {
	t.Log("Prueba IterarRango con diferentes combinaciones de limites")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	for i := 1; i <= 20; i++ {
		abb.Guardar(i, i*10)
	}

	contador := 0
	abb.IterarRango(nil, nil, func(clave int, dato int) bool {
		contador++
		return true
	})
	require.EqualValues(t, 20, contador)

	hasta := 5
	contador = 0
	abb.IterarRango(nil, &hasta, func(clave int, dato int) bool {
		require.True(t, clave <= hasta)
		contador++
		return true
	})
	require.EqualValues(t, 5, contador)

	desde := 15
	contador = 0
	abb.IterarRango(&desde, nil, func(clave int, dato int) bool {
		require.True(t, clave >= desde)
		contador++
		return true
	})
	require.EqualValues(t, 6, contador)

	desde, hasta = 8, 12
	contador = 0
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		require.True(t, clave >= desde && clave <= hasta)
		contador++
		return true
	})
	require.EqualValues(t, 5, contador)

	desde, hasta = 25, 30
	contador = 0
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		contador++
		return true
	})
	require.EqualValues(t, 0, contador)

	desde, hasta = 10, 5
	contador = 0
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		contador++
		return true
	})
	require.EqualValues(t, 0, contador)

	desde, hasta = 5, 15
	contador = 0
	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		contador++
		return clave < 10
	})
	require.EqualValues(t, 6, contador)
}

func TestIteradorRangoExternoABB(t *testing.T) {
	t.Log("Prueba IteradorRango externo con diferentes limites")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	for i := 1; i <= 15; i++ {
		abb.Guardar(i, i*10)
	}

	iter := abb.IteradorRango(nil, nil)
	contador := 0
	for iter.HaySiguiente() {
		contador++
		iter.Siguiente()
	}
	require.EqualValues(t, 15, contador)

	hasta := 5
	iter = abb.IteradorRango(nil, &hasta)
	contador = 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.True(t, clave <= hasta)
		contador++
		iter.Siguiente()
	}
	require.EqualValues(t, 5, contador)

	desde := 10
	iter = abb.IteradorRango(&desde, nil)
	contador = 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.True(t, clave >= desde)
		contador++
		iter.Siguiente()
	}
	require.EqualValues(t, 6, contador)

	desde, hasta = 5, 10
	iter = abb.IteradorRango(&desde, &hasta)
	contador = 0
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		require.True(t, clave >= desde && clave <= hasta)
		contador++
		iter.Siguiente()
	}
	require.EqualValues(t, 6, contador)

	desde, hasta = 20, 25
	iter = abb.IteradorRango(&desde, &hasta)
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, msj_iterador_terminado, func() { iter.VerActual() })
}

func TestVolumenABBInserciones(t *testing.T) {
	t.Log("Prueba de volumen con muchas inserciones y verificacion de orden")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	volumen := 5000

	for i := 0; i < volumen; i++ {
		abb.Guardar(i, i*2)
	}

	require.EqualValues(t, volumen, abb.Cantidad())

	for i := 0; i < volumen; i++ {
		require.True(t, abb.Pertenece(i))
		require.EqualValues(t, i*2, abb.Obtener(i))
	}

	clave_anterior := -1
	contador := 0
	abb.Iterar(func(clave int, dato int) bool {
		require.True(t, clave > clave_anterior)
		require.EqualValues(t, clave*2, dato)
		clave_anterior = clave
		contador++
		return true
	})

	require.EqualValues(t, volumen, contador)
}

func TestVolumenABBBorrados(t *testing.T) {
	t.Log("Prueba de volumen con inserciones y borrados")
	abb := TDADiccionario.CrearABB[int, string](igualdadIntsABB)
	volumen := 3000

	for i := 0; i < volumen; i++ {
		abb.Guardar(i, fmt.Sprintf("valor_%d", i))
	}

	require.EqualValues(t, volumen, abb.Cantidad())

	for i := 0; i < volumen; i += 2 {
		valor := abb.Borrar(i)
		require.EqualValues(t, fmt.Sprintf("valor_%d", i), valor)
		require.False(t, abb.Pertenece(i))
	}

	require.EqualValues(t, volumen/2, abb.Cantidad())

	for i := 1; i < volumen; i += 2 {
		require.True(t, abb.Pertenece(i))
		require.EqualValues(t, fmt.Sprintf("valor_%d", i), abb.Obtener(i))
	}
}

func TestVolumenABBIteradorExterno(t *testing.T) {
	t.Log("Prueba de volumen iterando con iterador externo")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	volumen := 4000

	for i := volumen; i > 0; i-- {
		abb.Guardar(i, i*3)
	}

	iter := abb.Iterador()
	clave_anterior := 0
	contador := 0

	for iter.HaySiguiente() {
		clave, dato := iter.VerActual()
		require.True(t, clave > clave_anterior)
		require.EqualValues(t, clave*3, dato)
		clave_anterior = clave
		contador++
		iter.Siguiente()
	}

	require.EqualValues(t, volumen, contador)
}

func TestVolumenABBRangos(t *testing.T) {
	t.Log("Prueba de volumen con iteradores de rango")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	volumen := 10000

	for i := 0; i < volumen; i++ {
		abb.Guardar(i, i)
	}

	desde := 2500
	hasta := 7500
	contador := 0

	abb.IterarRango(&desde, &hasta, func(clave int, dato int) bool {
		require.True(t, clave >= desde && clave <= hasta)
		contador++
		return true
	})

	require.EqualValues(t, 5001, contador)
}

func TestVolumenABBReemplazos(t *testing.T) {
	t.Log("Prueba de volumen con reemplazos de valores")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	volumen := 3000

	for i := 0; i < volumen; i++ {
		abb.Guardar(i, i)
	}

	require.EqualValues(t, volumen, abb.Cantidad())

	for i := 0; i < volumen; i++ {
		abb.Guardar(i, i*10)
	}

	require.EqualValues(t, volumen, abb.Cantidad())

	for i := 0; i < volumen; i++ {
		require.EqualValues(t, i*10, abb.Obtener(i))
	}
}

func TestABBBorrarRaiz(t *testing.T) {
	t.Log("Prueba borrar la raiz en diferentes casos")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	abb.Guardar(26, 26)
	require.EqualValues(t, 26, abb.Borrar(26))
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(26))

	abb.Guardar(26, 26)
	abb.Guardar(12, 12)
	require.EqualValues(t, 26, abb.Borrar(26))
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(12))

	abb = TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	abb.Guardar(26, 26)
	abb.Guardar(31, 31)
	require.EqualValues(t, 26, abb.Borrar(26))
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(31))

	abb = TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	abb.Guardar(26, 26)
	abb.Guardar(12, 12)
	abb.Guardar(31, 31)
	abb.Guardar(6, 6)
	abb.Guardar(22, 22)
	require.EqualValues(t, 26, abb.Borrar(26))
	require.EqualValues(t, 4, abb.Cantidad())

	clave_anterior := 0
	abb.Iterar(func(clave int, dato int) bool {
		require.True(t, clave > clave_anterior)
		clave_anterior = clave
		return true
	})
}

func TestABBDesbalanceadoAscendente(t *testing.T) {
	t.Log("Prueba ABB desbalanceado insertando en orden ascendente")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	claves := []int{3, 6, 9, 11, 12, 22, 26, 28, 31}

	for _, clave := range claves {
		abb.Guardar(clave, clave*2)
	}

	require.EqualValues(t, len(claves), abb.Cantidad())

	for _, clave := range claves {
		require.True(t, abb.Pertenece(clave))
		require.EqualValues(t, clave*2, abb.Obtener(clave))
	}

	clave_anterior := 0
	abb.Iterar(func(clave int, dato int) bool {
		require.True(t, clave > clave_anterior)
		clave_anterior = clave
		return true
	})
}

func TestABBDesbalanceadoDescendente(t *testing.T) {
	t.Log("Prueba ABB desbalanceado insertando en orden descendente")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	claves := []int{31, 28, 26, 22, 12, 11, 9, 6, 3}

	for _, clave := range claves {
		abb.Guardar(clave, clave*2)
	}

	require.EqualValues(t, len(claves), abb.Cantidad())

	for _, clave := range claves {
		require.True(t, abb.Pertenece(clave))
		require.EqualValues(t, clave*2, abb.Obtener(clave))
	}

	clave_anterior := 0
	contador := 0
	abb.Iterar(func(clave int, dato int) bool {
		require.True(t, clave > clave_anterior)
		clave_anterior = clave
		contador++
		return true
	})
	require.EqualValues(t, len(claves), contador)
}

func TestRangoUnSoloElemento(t *testing.T) {
	t.Log("Prueba IterarRango e IteradorRango cuando desde = hasta")
	abb := TDADiccionario.CrearABB[int, string](igualdadIntsABB)
	claves := []int{3, 6, 9, 12, 22, 26, 31}
	for _, clave := range claves {
		abb.Guardar(clave, fmt.Sprintf("%d", clave))
	}

	desde := 12
	hasta := 12
	contador := 0
	abb.IterarRango(&desde, &hasta, func(clave int, dato string) bool {
		require.EqualValues(t, 12, clave)
		require.EqualValues(t, "12", dato)
		contador++
		return true
	})
	require.EqualValues(t, 1, contador)

	iter := abb.IteradorRango(&desde, &hasta)
	require.True(t, iter.HaySiguiente())
	clave, dato := iter.VerActual()
	require.EqualValues(t, 12, clave)
	require.EqualValues(t, "12", dato)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestBorradoSecuencialCompleto(t *testing.T) {
	t.Log("Borra todos los elementos y verifica el estado en cada paso")
	abb := TDADiccionario.CrearABB[int, int](igualdadIntsABB)
	claves := []int{22, 9, 31, 6, 12, 26, 77}

	for _, clave := range claves {
		abb.Guardar(clave, clave)
	}

	require.EqualValues(t, len(claves), abb.Cantidad())

	for i, clave := range claves {
		require.True(t, abb.Pertenece(clave))
		require.EqualValues(t, clave, abb.Borrar(clave))
		require.EqualValues(t, len(claves)-i-1, abb.Cantidad())
		require.False(t, abb.Pertenece(clave))

		clave_anterior := 0
		abb.Iterar(func(c int, d int) bool {
			require.True(t, c > clave_anterior)
			clave_anterior = c
			return true
		})
	}

	require.EqualValues(t, 0, abb.Cantidad())
}
