package cola_prioridad

const (
	CAPACIDAD_INICIAL    = 10
	FACTOR_CRECIMIENTO   = 2
	FACTOR_DECRECIMIENTO = 4
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	h := new(heap[T])
	h.datos = make([]T, CAPACIDAD_INICIAL)
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
		panic("La cola esta vacia")
	}
	return h.datos[0]
}

func (h *heap[T]) Encolar(elem T) {
	if h.cantidad == len(h.datos) {
		h.redimensionar(len(h.datos) * FACTOR_CRECIMIENTO)
	}
	h.datos[h.cantidad] = elem
	h.cantidad++
	upheap(h.datos, h.cantidad-1, h.cmp)
}

func (h *heap[T]) Desencolar() T {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
	maximo := h.datos[0]
	h.cantidad--
	h.datos[0] = h.datos[h.cantidad]
	downheap(h.datos, 0, h.cantidad, h.cmp)

	if h.cantidad > 0 && h.cantidad*FACTOR_DECRECIMIENTO <= len(h.datos) && len(h.datos) > CAPACIDAD_INICIAL {
		h.redimensionar(len(h.datos) / FACTOR_CRECIMIENTO)
	}
	return maximo
}

func (h *heap[T]) redimensionar(nueva_capacidad int) {
	nuevo_arreglo := make([]T, nueva_capacidad)
	copy(nuevo_arreglo, h.datos[:h.cantidad])
	h.datos = nuevo_arreglo
}

func upheap[T any](arr []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	pos_padre := (pos - 1) / 2
	if cmp(arr[pos], arr[pos_padre]) > 0 {
		arr[pos], arr[pos_padre] = arr[pos_padre], arr[pos]
		upheap(arr, pos_padre, cmp)
	}
}

func downheap[T any](arr []T, pos int, tam int, cmp func(T, T) int) {
	hijo_izq := 2*pos + 1
	hijo_der := 2*pos + 2
	mayor := pos

	if hijo_izq < tam && cmp(arr[hijo_izq], arr[mayor]) > 0 {
		mayor = hijo_izq
	}
	if hijo_der < tam && cmp(arr[hijo_der], arr[mayor]) > 0 {
		mayor = hijo_der
	}
	if mayor != pos {
		arr[pos], arr[mayor] = arr[mayor], arr[pos]
		downheap(arr, mayor, tam, cmp)
	}
}

func heapify[T any](arr []T, tam int, cmp func(T, T) int) {
	for i := tam/2 - 1; i >= 0; i-- {
		downheap(arr, i, tam, cmp)
	}
}

func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	cant_elementos := len(elementos)
	heapify(elementos, cant_elementos, funcion_cmp)

	for i := cant_elementos - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downheap(elementos, 0, i, funcion_cmp)
	}
}
