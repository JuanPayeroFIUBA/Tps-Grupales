package diccionario

import (
	"fmt"
	"hash/fnv"
)

func convertirABytes[K any](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}
func fnvHash[K any](clave K, m uint64) uint64 {
	h := fnv.New64a()
	h.Write(convertirABytes(clave))
	return h.Sum64() % m
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

func (iter *IteradorDiccionario[K, V]) buscarProximaPosicionOcupada() {
	for iter.posicion_actual < iter.capacidad {
		if iter.tabla[iter.posicion_actual].estado == OCUPADO {
			return
		}
		iter.posicion_actual++
	}
	iter.posicion_actual = iter.capacidad
}
