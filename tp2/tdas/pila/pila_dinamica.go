package pila

/* Definición del struct pila proporcionado por la cátedra. */

const pila_vacia = "La pila esta vacia"
const capacidad_inicial = 1
const const_factor_redimension_agrandar = 2
const const_factor_redimension_achicar = 4

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, capacidad_inicial)
	return pila
}

func (p pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic(pila_vacia)
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	redimension := make([]T, tam)
	copy(redimension, p.datos)
	p.datos = redimension
}

func (p *pilaDinamica[T]) Apilar(elemento T) {
	if len(p.datos) <= p.cantidad {
		p.redimensionar(len(p.datos) * const_factor_redimension_agrandar)
	}
	p.datos[p.cantidad] = elemento
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic(pila_vacia)
	}
	p.cantidad--
	if len(p.datos)/const_factor_redimension_achicar > p.cantidad {
		p.redimensionar(len(p.datos) / const_factor_redimension_achicar)
	}
	return p.datos[p.cantidad]
}
