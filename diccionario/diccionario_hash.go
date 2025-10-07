package diccionario

//fuente de la funcion de hashing:
//https://pkg.go.dev/hash/fnv@go1.25.1

type Estado int

const (
	VACIO Estado = iota
	OCUPADO
	BORRADO
)
const factor_redimension_achicar = 0.3
const factor_redimension_agrandar = 0.7
const iteracion_completada = "El iterador termino de iterar"
const panic_clave_no_encontrada = "La clave no pertenece al diccionario"

type HashCerrado[K any, V comparable] struct {
	tabla                 []celdaHash[K, V]
	cantidad              int
	cantidad_borrados     int
	capacidad             int
	funcion_igualdad_hash func(K, K) bool
}

type celdaHash[K any, V comparable] struct {
	clave  K
	valor  V
	estado Estado
}
type IteradorDiccionario[K any, V comparable] struct {
	tabla           []celdaHash[K, V]
	capacidad       int
	posicion_actual int
}

func CrearHash[K any, V comparable](funcion_igualdad func(K, K) bool) Diccionario[K, V] {
	nuevo := new(HashCerrado[K, V])
	nuevo.funcion_igualdad_hash = funcion_igualdad
	nuevo.tabla = make([]celdaHash[K, V], 3)
	nuevo.cantidad = 0
	nuevo.capacidad = 3
	return nuevo
}

func (hash HashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *HashCerrado[K, V]) Guardar(clave K, dato V) {
	factor_carga := float64(hash.cantidad+hash.cantidad_borrados+1) / float64(hash.capacidad)
	if factor_carga > factor_redimension_agrandar {
		hash.redimensionar(2 * hash.capacidad)
	}

	ubicacion, encontrado := buscarPosicionGuardar(clave, *hash)

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

func (hash *HashCerrado[K, V]) Borrar(clave K) V {
	ubicacion, encontrado := buscarClave(clave, *hash)
	if !encontrado {
		panic(panic_clave_no_encontrada)
	}

	valor := hash.tabla[ubicacion].valor
	hash.tabla[ubicacion].estado = BORRADO
	hash.cantidad_borrados++
	hash.cantidad--

	factor_carga := float64(hash.cantidad) / float64(hash.capacidad)
	if factor_carga < factor_redimension_achicar && hash.capacidad > 3 {
		nueva_capacidad := hash.capacidad / 2
		if nueva_capacidad < 3 {
			nueva_capacidad = 3
		}
		hash.redimensionar(nueva_capacidad)
	}

	return valor
}

func (hash HashCerrado[K, V]) Pertenece(clave K) bool {
	_, encontrado := buscarClave(clave, hash)
	return encontrado
}

func (hash HashCerrado[K, V]) Obtener(clave K) V {
	ubicacion, encontrado := buscarClave(clave, hash)
	if !encontrado {
		panic(panic_clave_no_encontrada)
	}
	return hash.tabla[ubicacion].valor
}

func (hash HashCerrado[K, V]) Iterar(visitar func(K, V) bool) {
	for _, celda := range hash.tabla {
		if celda.estado == OCUPADO && !visitar(celda.clave, celda.valor) {
			return
		}
	}
}

func (hash HashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := new(IteradorDiccionario[K, V])
	iterador.tabla = hash.tabla
	iterador.capacidad = hash.capacidad
	iterador.posicion_actual = 0
	iterador.buscarProximaPosicionOcupada()
	return iterador
}

func (iter IteradorDiccionario[K, V]) HaySiguiente() bool {
	return iter.posicion_actual < iter.capacidad && iter.tabla[iter.posicion_actual].estado == OCUPADO
}

func (iter IteradorDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic(iteracion_completada)
	}
	clave := iter.tabla[iter.posicion_actual].clave
	valor := iter.tabla[iter.posicion_actual].valor
	return clave, valor
}

func (iter *IteradorDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic(iteracion_completada)
	}
	iter.posicion_actual++
	iter.buscarProximaPosicionOcupada()
}
