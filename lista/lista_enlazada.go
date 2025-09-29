package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

type IterListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

const lista_vacia = "La lista esta vacia"
const panic_iteracion_completada = "El iterador termino de iterar"

func nodoCrear[T any](dato T) *nodoLista[T] {
	nodo := new(nodoLista[T])
	nodo.dato = dato
	return nodo
}

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.largo = 0
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(elem T) {
	nuevo := nodoCrear(elem)
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	} else {
		nuevo.siguiente = lista.primero
	}
	lista.primero = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elem T) {
	nuevo := nodoCrear(elem)
	if lista.EstaVacia() {
		lista.primero = nuevo
	} else {
		lista.ultimo.siguiente = nuevo
	}
	lista.ultimo = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(lista_vacia)
	}
	eliminado := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return eliminado
}

func (lista listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic(lista_vacia)
	}
	return lista.primero.dato
}

func (lista listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic(lista_vacia)
	}
	return lista.ultimo.dato
}

func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.dato) {
			return
		}
		actual = actual.siguiente
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(IterListaEnlazada[T])
	iterador.lista = lista
	iterador.actual = lista.primero
	iterador.anterior = nil
	return iterador
}

func (iter IterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *IterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(panic_iteracion_completada)
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.siguiente
}

func (iter IterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic(panic_iteracion_completada)
	}
	return iter.actual.dato
}

func (iter *IterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic(panic_iteracion_completada)
	}
	eliminado := iter.actual.dato
	
	if iter.anterior == nil {
		iter.lista.primero = iter.actual.siguiente
		if iter.lista.primero == nil {
			iter.lista.ultimo = nil
		}
	} else {
		iter.anterior.siguiente = iter.actual.siguiente
		if iter.actual.siguiente == nil {
			iter.lista.ultimo = iter.anterior
		}
	}
	
	iter.actual = iter.actual.siguiente
	iter.lista.largo--
	return eliminado
}

func (iter *IterListaEnlazada[T]) Insertar(elem T) {
	nuevo := nodoCrear(elem)
	nuevo.siguiente = iter.actual
	
	if iter.anterior == nil {
		iter.lista.primero = nuevo
		if iter.actual == nil {
			iter.lista.ultimo = nuevo
		}
	} else {
		iter.anterior.siguiente = nuevo
		if iter.actual == nil {
			iter.lista.ultimo = nuevo
		}
	}
	
	iter.lista.largo++
}
