package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la lista no tiene elementos, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento al principio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero saca el primer elemento de la lista y lo devuelve. Si está vacía, devuelve un panic "La lista esta vacia".
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si está vacía, devuelve un panic "La lista esta vacia".
	VerPrimero() T

	// VerUltimo obtiene el valor del último elemento de la lista. Si está vacía, devuelve un panic "La lista esta vacia".
	VerUltimo() T

	// Largo devuelve la cantidad de elementos que tiene la lista.
	Largo() int

	// Iterar recorre la lista aplicando la función visitar a cada elemento, de primero a último.
	// La iteración se para si la función visitar devuelve false o al terminar la lista.
	Iterar(visitar func(T) bool)

	// Iterador devuelve un IteradorLista posicionado al inicio de la lista.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual al que apunta el iterador. Si el iterador terminó de iterar devuelve un panic "El iterador termino de iterar".
	VerActual() T

	// HaySiguiente devuelve true si hay un elemento siguiente para iterar, false en caso contrario.
	HaySiguiente() bool

	// Siguiente avanza el iterador al siguiente elemento. Si el iterador terminó de iterar devuelve un panic "El iterador termino de iterar".
	Siguiente()

	// Insertar agrega un elemento en la posición actual del iterador. El elemento se inserta antes del elemento actual y el iterador mantiene su posición relativa.
	Insertar(T)

	// Borrar elimina el elemento actual del iterador. Si el iterador terminó de iterar,
	// devuelve un panic "El iterador termino de iterar". Retorna el elemento eliminado.
	Borrar() T
}
