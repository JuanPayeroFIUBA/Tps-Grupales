package diccionario

import (
	"tdas/pila"
)

type abb[K, V any] struct {
	cantidad        int
	raiz            *nodoArbol[K, V]
	funcComparacion func(K, K) int
}

type nodoArbol[K, V any] struct {
	clave K
	valor V
	izq   *nodoArbol[K, V]
	der   *nodoArbol[K, V]
}

type iteradorDiccionarioConRango[K, V any] struct {
	desde           *K
	hasta           *K
	funcComparacion func(K, K) int
	pilaNodos       pila.Pila[*nodoArbol[K, V]]
}

func CrearABB[K, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	abb := new(abb[K, V])
	abb.funcComparacion = funcionCmp
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
	} else {
		padre_ubicacion, nodo_ubicacion, esIzquierdo := abb_buscar(clave, *abb)
		if nodo_ubicacion == nil {
			if esIzquierdo {
				padre_ubicacion.izq = nuevo
			} else {
				padre_ubicacion.der = nuevo
			}
			abb.cantidad++
		} else {
			nodo_ubicacion.valor = dato
		}
	}
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodo_padre, nodo, esIzquierdo := abb_buscar(clave, *abb)
	if nodo == nil {
		panic(msjPanicClaveNoEncontrada)
	}
	valor_eliminado := nodo.valor

	if nodo.izq == nil && nodo.der == nil {
		if nodo_padre == nil {
			abb.raiz = nil
		} else if esIzquierdo {
			nodo_padre.izq = nil
		} else {
			nodo_padre.der = nil
		}
	} else if nodo.izq != nil && nodo.der != nil {
		menor_por_derecha, padre_menor := encontrarMenorPorDerecha(nodo.der, nodo)
		nodo.clave = menor_por_derecha.clave
		nodo.valor = menor_por_derecha.valor

		if padre_menor == nodo {
			nodo.der = menor_por_derecha.der
		} else {
			padre_menor.izq = menor_por_derecha.der
		}
	} else {
		unico_hijo := nodo.izq
		if unico_hijo == nil {
			unico_hijo = nodo.der
		}
		if nodo_padre == nil {
			abb.raiz = unico_hijo
		} else if esIzquierdo {
			nodo_padre.izq = unico_hijo
		} else {
			nodo_padre.der = unico_hijo
		}
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
		panic(msjPanicClaveNoEncontrada)
	}
	return nodo.valor
}

func (abb abb[K, V]) Iterar(visitar func(K, V) bool) {
	iterarRango(abb, abb.raiz, nil, nil, visitar)
}

func (abb abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

func (abb abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(abb, abb.raiz, desde, hasta, visitar)
}

func (abb abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iterador_rango := new(iteradorDiccionarioConRango[K, V])
	iterador_rango.desde = desde
	iterador_rango.hasta = hasta
	iterador_rango.funcComparacion = abb.funcComparacion
	iterador_rango.pilaNodos = pila.CrearPilaDinamica[*nodoArbol[K, V]]()

	actual := abb.raiz
	ApilarRamaIzquierda(actual, iterador_rango)
	return iterador_rango
}

func (iter iteradorDiccionarioConRango[K, V]) HaySiguiente() bool {
	return !iter.pilaNodos.EstaVacia()
}

func (iter iteradorDiccionarioConRango[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(msjPanicIteracionCompletada)
	}
	nodo := iter.pilaNodos.VerTope()
	return nodo.clave, nodo.valor
}

func (iter *iteradorDiccionarioConRango[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(msjPanicIteracionCompletada)
	}
	actual := iter.pilaNodos.Desapilar()

	nodo := actual.der
	ApilarRamaIzquierda(nodo, iter)
}
