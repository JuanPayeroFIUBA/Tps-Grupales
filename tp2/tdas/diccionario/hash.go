package diccionario

//fuente de la funcion de hashing:
//https://pkg.go.dev/hash/fnv@go1.25.1

type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
)
const (
	capacidadInicial            = 7
	factorRedimAchicar          = 0.3
	divisorCapacidadHash        = 2
	factorRedimAgrandar         = 0.7
	multipicadorCapacidadHash   = 2
	msjPanicIteracionCompletada = "El iterador termino de iterar"
	msjPanicClaveNoEncontrada   = "La clave no pertenece al diccionario"
)

type hashCerrado[K any, V comparable] struct {
	tabla                 []celdaHash[K, V]
	cantidad              int
	cantidad_borrados     int
	capacidad             int
	funcion_igualdad_hash func(K, K) bool
}

type celdaHash[K any, V comparable] struct {
	clave  K
	valor  V
	estado estado
}
type iteradorDiccionario[K any, V comparable] struct {
	tabla           []celdaHash[K, V]
	capacidad       int
	posicion_actual int
}

func CrearHash[K any, V comparable](funcion_igualdad func(K, K) bool) Diccionario[K, V] {
	nuevo := new(hashCerrado[K, V])
	nuevo.funcion_igualdad_hash = funcion_igualdad
	nuevo.tabla = make([]celdaHash[K, V], capacidadInicial)
	nuevo.cantidad = 0
	nuevo.capacidad = capacidadInicial
	return nuevo
}

func (hash hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	factor_carga := float64(hash.cantidad+hash.cantidad_borrados) / float64(hash.capacidad)
	if factor_carga > factorRedimAgrandar {
		hash.redimensionar(multipicadorCapacidadHash * hash.capacidad)
	}

	ubicacion, encontrado := hash_buscar(clave, *hash, true)

	if hash.tabla[ubicacion].estado == BORRADO {
		hash.cantidad_borrados--
	}

	if !encontrado {
		hash.cantidad++
	}

	hash.tabla[ubicacion].clave = clave
	hash.tabla[ubicacion].valor = dato
	hash.tabla[ubicacion].estado = OCUPADO
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	ubicacion, encontrado := hash_buscar(clave, *hash, false)
	if !encontrado {
		panic(msjPanicClaveNoEncontrada)
	}

	valor := hash.tabla[ubicacion].valor
	hash.tabla[ubicacion].estado = BORRADO
	hash.cantidad_borrados++
	hash.cantidad--

	factor_carga := float64(hash.cantidad) / float64(hash.capacidad)
	if factor_carga < factorRedimAchicar && hash.capacidad > capacidadInicial {
		nueva_capacidad := hash.capacidad / divisorCapacidadHash
		nueva_capacidad = max(nueva_capacidad, capacidadInicial)
		hash.redimensionar(nueva_capacidad)
	}

	return valor
}

func (hash hashCerrado[K, V]) Pertenece(clave K) bool {
	_, encontrado := hash_buscar(clave, hash, false)
	return encontrado
}

func (hash hashCerrado[K, V]) Obtener(clave K) V {
	ubicacion, encontrado := hash_buscar(clave, hash, false)
	if !encontrado {
		panic(msjPanicClaveNoEncontrada)
	}
	return hash.tabla[ubicacion].valor
}

func (hash hashCerrado[K, V]) Iterar(visitar func(K, V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO && !visitar(celda.clave, celda.valor) {
			return
		}
	}
}

func (hash hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(iteradorDiccionario[K, V])
	iterador.tabla = hash.tabla
	iterador.capacidad = hash.capacidad
	iterador.posicion_actual = 0
	iterador.buscarProximaPosicionOcupada()
	return iterador
}

func (iter iteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.posicion_actual < iter.capacidad
}

func (iter iteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(msjPanicIteracionCompletada)
	}
	clave := iter.tabla[iter.posicion_actual].clave
	valor := iter.tabla[iter.posicion_actual].valor
	return clave, valor
}

func (iter *iteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(msjPanicIteracionCompletada)
	}
	iter.posicion_actual++
	iter.buscarProximaPosicionOcupada()
}
