package pila

func (pila *pilaDinamica[T]) Transformar(aplicar func(T) T) Pila[T] {
	pila_sal := new(pilaDinamica[T])
	pila_aux := new(pilaDinamica[T])
	for !pila.EstaVacia() {
		pila_aux.Apilar(pila.Desapilar())
	}
	for !pila_aux.EstaVacia() {
		valor := pila_aux.Desapilar()
		pila_sal.Apilar(aplicar(valor))
		pila.Apilar(valor)
	}

	return pila_sal
}
