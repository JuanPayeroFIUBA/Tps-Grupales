package cola

const cola_vacia = "La cola esta vacia"

type nodoCola[T any] struct {
	dato      T
	siguiente *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func nodoCrear[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.dato = dato
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	return cola
}

func (c colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic(cola_vacia)
	}
	return c.primero.dato
}

func (c *colaEnlazada[T]) Encolar(elem T) {
	encolado := nodoCrear(elem)

	if c.EstaVacia() {
		c.primero = encolado
	} else {
		c.ultimo.siguiente = encolado
	}
	c.ultimo = encolado
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic(cola_vacia)
	}
	desencolado := c.primero.dato
	c.primero = c.primero.siguiente
	return desencolado
}
