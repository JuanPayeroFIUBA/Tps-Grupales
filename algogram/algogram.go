package algogram

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	heap "tdas/cola_prioridad"
	dicci "tdas/diccionario"
)

type usuario struct {
	nombre            string
	indice_en_arreglo int
}

type post struct {
	id         int
	publicador usuario
	mensaje    string
	//likes          dicci.DiccionarioOrdenado[string, int]
	//cantidad_likes int //medio innecesario este
}

func igualdadStrings(a, b string) bool {
	return a == b
}
func igualdadInts(a, b int) bool {
	return a == b
}
func comparacionStrings(a, b string) int {
	return strings.Compare(a, b)
}

//	func Setup(usuarios []string) {
//		hash_usuarios := dicci.CrearHash[string, int](igualdadStrings)
//		hash_posts := dicci.CrearHash[int, string](igualdadInts)
//		hash_heap_posts := dicci.CrearHash[string, heap.ColaPrioridad[post]](igualdadStrings)
//
//		var cuenta_logueada usuario
//
//		for indice, usuario := range usuarios {
//			hash_usuarios.Guardar(usuario, indice)
//			hash_heap_posts.Guardar(usuario, heap.CrearHeap(func(post1, post2 post) int {
//				indice_usuario1 := int(math.Abs(float64(post1.publicador.indice_en_arreglo - indice)))
//				indice_usuario2 := int(math.Abs(float64(post2.publicador.indice_en_arreglo - indice)))
//				if indice_usuario1 == indice_usuario2 {
//					return post1.id - post2.id
//				}
//				return indice_usuario1 - indice_usuario2
//			}))
//		}
//		LeerEntradaComandos(hash_usuarios, hash_heap_posts, cuenta_logueada)
//	}

func Setup(usuarios []string) {
	hash_usuarios := dicci.CrearHash[string, int](igualdadStrings)                           //diccionario usuarios
	hash_posts := dicci.CrearHash[int, dicci.DiccionarioOrdenado[string, int]](igualdadInts) //para likes
	hash_feed := dicci.CrearHash[string, heap.ColaPrioridad[post]](igualdadStrings)          //para el feed

	var cuenta_logueada usuario

	for indice, usuario := range usuarios {
		hash_usuarios.Guardar(usuario, indice)
		hash_feed.Guardar(usuario, heap.CrearHeap(func(post1, post2 post) int {
			indice_usuario1 := int(math.Abs(float64(post1.publicador.indice_en_arreglo - indice)))
			indice_usuario2 := int(math.Abs(float64(post2.publicador.indice_en_arreglo - indice)))
			if indice_usuario1 == indice_usuario2 {
				return post1.id - post2.id
			}
			return indice_usuario1 - indice_usuario2
		}))
	}
	LeerEntradaComandos(hash_usuarios, hash_feed, hash_posts, cuenta_logueada)
}

func LeerEntradaArchivoYGuardar() []string {
	operaciones := []string{}
	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		if entrada.Text() == "" {
			break
		}
		operaciones = append(operaciones, entrada.Text())
	}
	return operaciones
}

func LeerEntradaComandos(hash_users dicci.Diccionario[string, int], hash_feed dicci.Diccionario[string, heap.ColaPrioridad[post]], diccionario_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], logueado usuario) {
	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		if entrada.Text() == "" {
			break
		}
	}
}

//completar esto para que llame a cada funcion dependiendo cual sea el comando

func Login(hash_users dicci.Diccionario[string, int], logueado *usuario, usuario_a_loguear string) {
	if logueado.nombre != "" {
		fmt.Println("Error: Ya habia un usuario loggeado")
		return
	}
	logueado.nombre = usuario_a_loguear
	logueado.indice_en_arreglo = hash_users.Obtener(usuario_a_loguear)
	fmt.Println("Hola", usuario_a_loguear)
}

func Logout(logueado *usuario) {
	if logueado.nombre == "" {
		fmt.Println("Error: no habia usuario loggeados")
		return
	}
	logueado.nombre = ""
	logueado.indice_en_arreglo = -1
	fmt.Println("Adios")
}

func Publicar(hash_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], hash_feed dicci.Diccionario[string, heap.ColaPrioridad[post]], logueado *usuario, ultimo_post_id int, mensaje string) {
	if logueado.nombre == "" {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	hash_posts.Guardar(ultimo_post_id, dicci.CrearABB[string, int](comparacionStrings))
	for iter := hash_feed.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		usuario, heap := iter.VerActual()
		if usuario != logueado.nombre {
			heap.Encolar(post{mensaje: mensaje, id: ultimo_post_id, publicador: *logueado})
		}
	}
	fmt.Println("Post publicado")
}

//func VerProximoPostEnFeed(hash_posts dicci.Diccionario[int, post], hash_heap_posts dicci.Diccionario[string, heap.ColaPrioridad[post]], logueado *usuario) {
//	if logueado.nombre == "" || hash_heap_posts.Obtener(logueado.nombre).EstaVacia() {
//		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
//		return
//	}
//	heap_posts := hash_heap_posts.Obtener(logueado.nombre)
//	//if heap_posts.EstaVacia() {
//	//	fmt.Println("Usuario no loggeado o no hay mas posts para ver")
//	//	return
//	//}
//	post := heap_posts.Desencolar()
//	fmt.Println("Post ID %d \n%s dijo: %s \nLikes: %d", post.id, post.publicador.nombre, post.mensaje, post.cantidad_likes)
//}

//para lo del hash_heap_posts, mejor usar el hash, pero con abbs, misma complejidad mostrar y encontrar es O(log n) asi que no rompe ... el problema es que para encontrar el menor(osea el primer post del feed) abria que hacer un iterador del abb pero eso pediria apilar un monton de nodos, los cuales se desperdiciarian, por lo que no lo veo factible

func VerProximoPostEnFeed(hash_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], hash_feed dicci.Diccionario[string, heap.ColaPrioridad[post]], logueado *usuario) {
	if logueado.nombre == "" || hash_feed.Obtener(logueado.nombre).EstaVacia() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}

	feed := hash_feed.Obtener(logueado.nombre) //O(1)
	post := feed.Desencolar()                  //O(log p)
	cosas := hash_posts.Obtener(post.id)       //O(log p)
	fmt.Printf("Post ID %d \n%s dijo: %s \nLikes: %d", post.id, post.publicador.nombre, post.mensaje, cosas.Cantidad())
}

func LikearUnPost(hash_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], logueado *usuario, id int) {
	if logueado.nombre == "" || !hash_posts.Pertenece(id) {
		fmt.Println("Error: Usuario no loggeado o Post inexistente")
		return
	}
	abb_likes := hash_posts.Obtener(id)
	abb_likes.Guardar(logueado.nombre, 0)
	fmt.Println("Post likeado")
}

func VerLikesDeUnPost(hash_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], id int) {
	if !hash_posts.Pertenece(id) || hash_posts.Obtener(id).Cantidad() == 0 {
		fmt.Println("Error: Post inexistente o sin likes")
		return
	}
	abb_likes := hash_posts.Obtener(id)
	fmt.Printf("El post tiene %d likes: \n", abb_likes.Cantidad())
	for iter := abb_likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre, _ := iter.VerActual()
		fmt.Printf(nombre, "\n")
	}
}
