package diccionario

//fuente de la funcion de hashing:
//https://pkg.go.dev/hash/fnv@go1.25.1

import (
	"fmt"
	"hash/fnv"
)

type Estado int

const (
	VACIO Estado = iota
	OCUPADO
	BORRADO
)
const factor_redimension_achicar = 0.3
const factor_redimension_agrandar = 0.7
const iteracion_completada = "El iterador termino de iterar"
const panic_clave_no_encontrada = "La clave no pertenece al diccionario"

type HashCerrado[K any, V comparable] struct {
	tabla                 []celdaHash[K, V]
	cantidad              int
	cantidad_borrados     int
	capacidad             int
	funcion_igualdad_hash func(K, K) bool
}

type celdaHash[K any, V comparable] struct {
	clave  K
	valor  V
	estado Estado
}
type IteradorDiccionario[K any, V comparable] struct {
	tabla           []celdaHash[K, V]
	capacidad       int
	posicion_actual int
}

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
func fnvHash[K any](clave K, m uint64) uint64 {
	h := fnv.New64a()
	h.Write(convertirABytes(clave))
	return h.Sum64() % m
}

func CrearHash[K any, V comparable](funcion_igualdad func(K, K) bool) Diccionario[K, V] {
	nuevo := new(HashCerrado[K, V])
	nuevo.funcion_igualdad_hash = funcion_igualdad
	nuevo.tabla = make([]celdaHash[K, V], 3)
	nuevo.cantidad = 0
	nuevo.capacidad = 3
	return nuevo
}

func hash_obtener[K any, V comparable](clave K, hash HashCerrado[K, V]) (int, bool) {
	ubicacion := int(fnvHash(clave, uint64(hash.capacidad)))
	celda_actual := hash.tabla[ubicacion]
	contador := 0

	for {
		celda_actual = hash.tabla[ubicacion]
		if celda_actual.estado == VACIO {
			return ubicacion, false
		}
		if celda_actual.estado == OCUPADO && hash.funcion_igualdad_hash(celda_actual.clave, clave) {
			return ubicacion, true
		}
		contador++
		if contador == hash.capacidad {
			return ubicacion, false
		}
		ubicacion = (ubicacion + 1) % hash.capacidad
	}
}

func (hash *HashCerrado[K, V]) redimensionar(nueva_capacidad int) {
	tabla_a_redimensionar := hash.tabla
	tabla_redimensionada := make([]celdaHash[K, V], nueva_capacidad)
	hash.tabla = tabla_redimensionada
	hash.cantidad = 0
	hash.capacidad = nueva_capacidad
	hash.cantidad_borrados = 0

	for _, celda := range tabla_a_redimensionar {
		if celda.estado == OCUPADO {
			hash.Guardar(celda.clave, celda.valor)
		}
	}
}

func (hash HashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *HashCerrado[K, V]) Guardar(clave K, dato V) {
	if factor_redimension_agrandar < (float64(hash.cantidad+hash.cantidad_borrados+1) / float64(hash.capacidad)) {
		hash.redimensionar(2 * hash.capacidad)
	}
	ubicacion, encontrado := hash_obtener(clave, *hash)
	if !encontrado {
		hash.cantidad++
	}
	if hash.tabla[ubicacion].estado == BORRADO {
		hash.cantidad_borrados--
	}
	hash.tabla[ubicacion].clave = clave
	hash.tabla[ubicacion].valor = dato
	hash.tabla[ubicacion].estado = OCUPADO

}

func (hash *HashCerrado[K, V]) Borrar(clave K) V {
	if factor_redimension_achicar > (float64(hash.cantidad-1) / float64(hash.capacidad)) {
		hash.redimensionar(hash.capacidad / 2)
	}
	ubicacion, encontrado := hash_obtener(clave, *hash)
	if !encontrado {
		panic(panic_clave_no_encontrada)
	}
	hash.tabla[ubicacion].estado = BORRADO
	hash.cantidad_borrados++
	hash.cantidad--
	return hash.tabla[ubicacion].valor
}

func (hash HashCerrado[K, V]) Pertenece(clave K) bool {
	_, encontrado := hash_obtener(clave, hash)
	return encontrado
}

func (hash HashCerrado[K, V]) Obtener(clave K) V {
	ubicacion, encontrado := hash_obtener(clave, hash)
	if !encontrado {
		panic(panic_clave_no_encontrada)
	}
	return hash.tabla[ubicacion].valor
}

func (hash HashCerrado[K, V]) Iterar(visitar func(K, V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO && !visitar(celda.clave, celda.valor) {
			return
		}
	}
}

func (iter *IteradorDiccionario[K, V]) buscarProximaPosicionOcupada() {
	for iter.posicion_actual < iter.capacidad {
		if iter.tabla[iter.posicion_actual].estado == OCUPADO {
			return
		}
		iter.posicion_actual++
	}
	iter.posicion_actual = iter.capacidad
}

func (hash HashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(IteradorDiccionario[K, V])
	iterador.tabla = hash.tabla
	iterador.capacidad = hash.capacidad
	iterador.posicion_actual = 0
	iterador.buscarProximaPosicionOcupada()
	return iterador
}

func (iter IteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.posicion_actual < iter.capacidad && iter.tabla[iter.posicion_actual].estado == OCUPADO
}

func (iter IteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(iteracion_completada)
	}
	clave := iter.tabla[iter.posicion_actual].clave //prolijidad che
	valor := iter.tabla[iter.posicion_actual].valor
	return clave, valor
}

func (iter *IteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(iteracion_completada)
	}
	iter.posicion_actual++
	iter.buscarProximaPosicionOcupada()
}
