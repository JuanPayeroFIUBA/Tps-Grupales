package cola_prioridad

func Swap[T any](x, y *T) {
	*x, *y = *y, *x
}

func upheap[T any](arr []T, pos int, cmp func(T, T) int) {
	if pos == 0 {
		return
	}
	pos_padre := (pos - 1) / 2
	if cmp(arr[pos], arr[pos_padre]) > 0 {
		Swap(&arr[pos], &arr[pos_padre])
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
		Swap(&arr[pos], &arr[mayor])
		downheap(arr, mayor, tam, cmp)
	}
}

func heapify[T any](arr []T, tam int, cmp func(T, T) int) {
	for i := tam/2 - 1; i >= 0; i-- {
		downheap(arr, i, tam, cmp)
	}
}
