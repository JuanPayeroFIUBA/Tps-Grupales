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

func hash_buscar[K any, V comparable](clave K, hash hashCerrado[K, V], queremos_guardar bool) (int, bool) {
	ubicacion := int(fnvHash(clave, uint64(hash.capacidad)))
	primera_borrada := -1
	iteraciones := 0

	for iteraciones < hash.capacidad {
		celda := hash.tabla[ubicacion]

		if celda.estado == OCUPADO && hash.funcion_igualdad_hash(celda.clave, clave) {
			return ubicacion, true
		}

		if queremos_guardar && celda.estado == BORRADO && primera_borrada == -1 {
			primera_borrada = ubicacion
		}

		if celda.estado == VACIO {
			if queremos_guardar && primera_borrada != -1 {
				return primera_borrada, false
			}
			return ubicacion, false
		}

		ubicacion = (ubicacion + 1) % hash.capacidad
		iteraciones++
	}

	if queremos_guardar && primera_borrada != -1 {
		return primera_borrada, false
	}
	return ubicacion, false
}

func (hash *hashCerrado[K, V]) redimensionar(nueva_capacidad int) {
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

func (iter *iteradorDiccionario[K, V]) buscarProximaPosicionOcupada() {
	for iter.posicion_actual < iter.capacidad {
		if iter.tabla[iter.posicion_actual].estado == OCUPADO {
			return
		}
		iter.posicion_actual++
	}
	iter.posicion_actual = iter.capacidad
}
