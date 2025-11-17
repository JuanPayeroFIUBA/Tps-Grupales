package diccionario

func abb_buscar[K, V any](clave K, abb abb[K, V]) (*nodoArbol[K, V], *nodoArbol[K, V], bool) {
	var padre *nodoArbol[K, V] = nil
	actual := abb.raiz
	izquierda := false
	for actual != nil {
		comparacion := abb.funcComparacion(clave, actual.clave)
		if comparacion == 0 {
			return padre, actual, izquierda
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

func iterarRango[K, V any](abb abb[K, V], actual *nodoArbol[K, V], desde, hasta *K, visitar func(K, V) bool) bool {
	if actual == nil {
		return true
	}

	if desde != nil && abb.funcComparacion(actual.clave, *desde) < 0 {
		return iterarRango(abb, actual.der, desde, hasta, visitar)
	}

	if hasta != nil && abb.funcComparacion(actual.clave, *hasta) > 0 {
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

func ApilarRamaIzquierda[K, V any](nodo *nodoArbol[K, V], iter *iteradorDiccionarioConRango[K, V]) {
	for nodo != nil {
		if iter.desde != nil && iter.funcComparacion(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
			continue
		}
		if iter.hasta != nil && iter.funcComparacion(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izq
			continue
		}

		iter.pilaNodos.Apilar(nodo)
		nodo = nodo.izq
	}
}
