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

func hash_buscar[K any, V comparable](clave K, hash hashCerrado[K, V], accion_guardar bool) (int, bool) {
	ubicacion := int(fnvHash(clave, uint64(hash.capacidad)))
	primera_borrada := -1
	iteraciones := 0

	for iteraciones < hash.capacidad {
		celda := hash.tabla[ubicacion]

		if celda.estado == OCUPADO && hash.funcion_igualdad_hash(celda.clave, clave) {
			return ubicacion, true
		}

		if accion_guardar && celda.estado == BORRADO && primera_borrada == -1 {
			primera_borrada = ubicacion
		}

		if celda.estado == VACIO {
			if accion_guardar && primera_borrada != -1 {
				return primera_borrada, false
			}
			return ubicacion, false
		}

		ubicacion = (ubicacion + 1) % hash.capacidad
		iteraciones++
	}

	if accion_guardar && primera_borrada != -1 {
		return primera_borrada, false
	}
	return ubicacion, false
}

func (hash *hashCerrado[K, V]) redimensionar(nueva_capacidad int) {
	tabla_a_redimensionar := hash.tabla
	hash.tabla = make([]celdaHash[K, V], nueva_capacidad)
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

func abb_buscar[K, V any](clave K, abb abb[K, V]) (*nodoArbol[K, V], *nodoArbol[K, V], bool) {
	var padre *nodoArbol[K, V] = nil
	actual := abb.raiz
	izquierda := false
	for actual != nil {
		comparacion := abb.funcion_comparacion(clave, actual.clave)
		if comparacion == 0 {
			return padre, actual, false
		}
		padre = actual
		if comparacion < 0 {
			actual = actual.izq
			izquierda = true
		}
		if comparacion > 0 {
			actual = actual.der
			izquierda = false
		}
	}
	return padre, nil, izquierda

}

func encontrarMenorPorDerecha[K, V any](nodo, padre *nodoArbol[K, V]) (*nodoArbol[K, V], *nodoArbol[K, V]) {
	if nodo.izq == nil {
		return nodo, padre
	}
	return encontrarMenorPorDerecha(nodo.izq, nodo)
}

func iterar[K, V any](actual *nodoArbol[K, V], visitar func(K, V) bool) bool {
	if actual == nil {
		return true
	}
	if !iterar(actual.izq, visitar) {
		return false
	}

	if !visitar(actual.clave, actual.valor) {
		return false
	}

	return iterar(actual.der, visitar)
}

func iterarRango[K, V any](abb abb[K, V], actual *nodoArbol[K, V], desde, hasta *K, visitar func(K, V) bool) bool {
	if actual == nil {
		return true
	}
	//comparacion_con_desde := 0
	//comparacion_con_hasta := 0
	//if desde != nil  {
	//	comparacion_con_desde = abb.funcion_comparacion(actual.clave, *desde)
	//}
	//if hasta != nil {
	//	comparacion_con_hasta = abb.funcion_comparacion(actual.clave, *hasta)
	//}
	//
	//if comparacion_con_hasta > 0 {
	//	return iterarRango(abb, actual.izq, desde, hasta, visitar)
	//}
	//
	//if comparacion_con_desde >= 0 && comparacion_con_hasta <= 0 {
	//	if !iterarRango(abb, actual.izq, desde, hasta, visitar) {
	//		return false
	//	}
	//	if !visitar(actual.clave, actual.valor) {
	//		return false
	//	}
	//	return iterarRango(abb, actual.der, desde, hasta, visitar)
	//}
	//
	//if comparacion_con_desde < 0 {
	//	return iterarRango(abb, actual.der, desde, hasta, visitar)
	//}
	//return iterarRango(abb, actual.der, desde, hasta, visitar)

	if desde != nil && abb.funcion_comparacion(actual.clave, *desde) < 0 {
		return iterarRango(abb, actual.der, desde, hasta, visitar)
	}

	if hasta != nil && abb.funcion_comparacion(actual.clave, *hasta) > 0 {
		return iterarRango(abb, actual.izq, desde, hasta, visitar)
	}
	if !iterarRango(abb, actual.izq, desde, hasta, visitar) {
		return false
	}
	if !visitar(actual.clave, actual.valor) {
		return false
	}
	return iterarRango(abb, actual.der, desde, hasta, visitar)
}
