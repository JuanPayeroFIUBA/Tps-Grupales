package diccionario

import (
	"tdas/pila"
)

type abb[K, V any] struct {
	altura              int
	cantidad            int
	raiz                *nodoArbol[K, V]
	funcion_comparacion func(K, K) int
}

type nodoArbol[K, V any] struct {
	clave K
	valor V
	izq   *nodoArbol[K, V]
	der   *nodoArbol[K, V]
}

type iteradorDiccionarioOrdenado[K, V any] struct {
	pila_de_nodos pila.Pila[*nodoArbol[K, V]]
}
type iteradorDiccionarioOrdenadoConRango[K, V any] struct {
	desde               *K
	hasta               *K
	funcion_comparacion func(K, K) int
	pila_de_nodos       pila.Pila[*nodoArbol[K, V]]
}

func CrearABB[K, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.altura = 0
	abb.funcion_comparacion = funcionCmp
	return abb
}

func (abb abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nuevo := new(nodoArbol[K, V])
	nuevo.clave = clave
	nuevo.valor = dato
	if abb.raiz == nil {
		abb.raiz = nuevo
		abb.cantidad++
		return
	}
	padre_ubicacion, nodo_ubicacion, izquierda := abb_buscar(clave, *abb)
	if nodo_ubicacion == nil {
		if izquierda {
			padre_ubicacion.izq = nuevo
		} else {
			padre_ubicacion.der = nuevo
		}
		abb.cantidad++
	} else {
		nodo_ubicacion.valor = dato
	}
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo_padre, nodo, izquierda := abb_buscar(clave, *abb)
	if nodo == nil || abb.cantidad == 0 {
		panic(panic_clave_no_encontrada)
	}
	valor_eliminado := nodo.valor

	if nodo.izq == nil && nodo.der == nil {
		if nodo_padre == nil {
			abb.raiz = nil
		} else if izquierda {
			nodo_padre.izq = nil
		} else {
			nodo_padre.der = nil
		}
		abb.cantidad--
		return valor_eliminado
	}

	if nodo.izq != nil && nodo.der != nil {
		menor_por_derecha, padre_menor := encontrarMenorPorDerecha(nodo.der, nodo)
		nodo.clave = menor_por_derecha.clave
		nodo.valor = menor_por_derecha.valor

		if padre_menor == nodo {
			nodo.der = menor_por_derecha.der
		} else {
			padre_menor.izq = menor_por_derecha.der
		}
		abb.cantidad--
		return valor_eliminado
	}

	unico_hijo := nodo.izq
	if unico_hijo == nil {
		unico_hijo = nodo.der
	}
	if nodo_padre == nil {
		abb.raiz = unico_hijo
	} else if izquierda {
		nodo_padre.izq = unico_hijo
	} else {
		nodo_padre.der = unico_hijo
	}
	abb.cantidad--
	return valor_eliminado
}

func (abb abb[K, V]) Pertenece(clave K) bool {
	_, nodo, _ := abb_buscar(clave, abb)
	return nodo != nil
}

func (abb abb[K, V]) Obtener(clave K) V {
	_, nodo, _ := abb_buscar(clave, abb)
	if nodo == nil || abb.cantidad == 0 {
		panic(panic_clave_no_encontrada)
	}
	return nodo.valor
}

func (abb Abb[K, V]) Iterar(visitar func(K, V) bool) {
	abb.IterarRango(nil, nil, visitar)
}

func (abb Abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(abb, abb.raiz, desde, hasta, visitar)
}

func (abb abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador_rango := new(iteradorDiccionarioOrdenadoConRango[K, V])
	iterador_rango.desde = desde
	iterador_rango.hasta = hasta
	iterador_rango.funcion_comparacion = abb.funcion_comparacion
	iterador_rango.pila_de_nodos = pila.CrearPilaDinamica[*nodoArbol[K, V]]()

	actual := abb.raiz
	for actual != nil {
		if desde != nil && abb.funcion_comparacion(actual.clave, *desde) < 0 {
			actual = actual.der
			continue
		}
		if hasta != nil && abb.funcion_comparacion(actual.clave, *hasta) > 0 {
			actual = actual.izq
			continue
		}
		iterador_rango.pila_de_nodos.Apilar(actual)
		actual = actual.izq
	}
	return iterador_rango
}

func (iter iteradorDiccionarioOrdenadoConRango[K, V]) HaySiguiente() bool {
	return !iter.pila_de_nodos.EstaVacia()
}

func (iter iteradorDiccionarioOrdenadoConRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(iteracion_completada)
	}
	nodo := iter.pila_de_nodos.VerTope()
	return nodo.clave, nodo.valor
}

func (iter *iteradorDiccionarioOrdenadoConRango[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(iteracion_completada)
	}
	actual := iter.pila_de_nodos.Desapilar()

	nodo := actual.der
	for nodo != nil {
		if iter.desde != nil && iter.funcion_comparacion(nodo.clave, *iter.desde) < 0 {
			nodo = nodo.der
			continue
		}
		if iter.hasta != nil && iter.funcion_comparacion(nodo.clave, *iter.hasta) > 0 {
			nodo = nodo.izq
			continue
		}

		iter.pila_de_nodos.Apilar(nodo)
		nodo = nodo.izq
	}

}
