package cola_prioridad

const (
	constCapacidadInicial    = 10
	constFactorCrecimiento   = 2
	constFactorDecrecimiento = 4
	msjPaniColaVacia         = "La cola esta vacia"
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := new(heap[T])
	h.datos = make([]T, constCapacidadInicial)
	h.cantidad = 0
	h.cmp = funcion_cmp
	return h
}

func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := new(heap[T])
	h.datos = make([]T, len(arreglo))
	copy(h.datos, arreglo)
	h.cantidad = len(arreglo)
	h.cmp = funcion_cmp
	heapify(h.datos, h.cantidad, h.cmp)
	return h
}

func (h *heap[T]) EstaVacia() bool {
	return h.cantidad == 0
}

func (h *heap[T]) Cantidad() int {
	return h.cantidad
}

func (h *heap[T]) VerMax() T {
	if h.EstaVacia() {
		panic(msjPaniColaVacia)
	}
	return h.datos[0]
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.datos) {
		h.redimensionar(len(h.datos) * constFactorCrecimiento)
	}
	h.datos[h.cantidad] = elem
	h.cantidad++
	upheap(h.datos, h.cantidad-1, h.cmp)
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic(msjPaniColaVacia)
	}
	maximo := h.datos[0]
	h.cantidad--
	h.datos[0] = h.datos[h.cantidad]
	downheap(h.datos, 0, h.cantidad, h.cmp)

	if h.cantidad > 0 && h.cantidad*constFactorDecrecimiento <= len(h.datos) && len(h.datos) > constCapacidadInicial {
		h.redimensionar(len(h.datos) / constFactorCrecimiento)
	}
	return maximo
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	cant_elementos := len(elementos)
	heapify(elementos, cant_elementos, funcion_cmp)

	for ult_relativo := cant_elementos - 1; ult_relativo > 0; ult_relativo-- {
		elementos[0], elementos[ult_relativo] = elementos[ult_relativo], elementos[0]
		downheap(elementos, 0, ult_relativo, funcion_cmp)
	}
}
