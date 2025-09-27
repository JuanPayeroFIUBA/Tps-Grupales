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
	actual    *nodoLista[T]
	siguiente *nodoLista[T]
}

const lista_vacia = "La lista esta vacia"

func CrearListaEnlazada[T any]() Lista[T] {
	lista := new(listaEnlazada[T])
	lista.largo = 0
	return lista
}

func (lista listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}
func (lista listaEnlazada[T]) InsertarPrimero(T) {
}
func (lista listaEnlazada[T]) InsertarUltimo(T) {

}
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic(lista_vacia)
	}
	eliminado := lista.primero.dato
	lista.primero = lista.primero.siguiente
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

}
func (lista listaEnlazada[T]) Iterador() IteradorLista[T] {
	iterador := new(IterListaEnlazada[T])
	return iterador
}

func (iter IterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *IterListaEnlazada[T]) Siguiente() {
	iter.actual = iter.siguiente
}

func (iter IterListaEnlazada[T]) VerActual() T {
	return iter.actual.dato
}

func (iter *IterListaEnlazada[T]) Borrar() T {
	eliminado := iter.actual.dato
	iter.actual.dato = nil //aca hay error, vercomo solucionarlo
	return eliminado
}

func (iter *IterListaEnlazada[T]) Insertar(elem T) {
	iter.actual.dato = elem
}
