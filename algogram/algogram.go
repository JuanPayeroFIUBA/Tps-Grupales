package algogram

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	heap "tdas/cola_prioridad"
	dicci "tdas/diccionario"
)

const (
	constLogin        = "login"
	constLogout       = "logout"
	constPublicar     = "publicar"
	constConstVerFeed = "ver_siguiente_feed"
	constDarLike      = "likear_post"
	constVerLikes     = "mostrar_likes"
)

type usuario struct {
	nombre            string
	indice_en_arreglo int
	feed              heap.ColaPrioridad[post_en_feed]
}
type post_en_feed struct {
	mensaje                  string
	id_post                  int
	dif_indices_con_loggeado int
}

type post struct {
	id         int
	publicador usuario
	mensaje    string
	likes      dicci.DiccionarioOrdenado[string, int]
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
				return -(post1.id - post2.id)
			}
			return -(indice_usuario1 - indice_usuario2)
		}))
	}
	LeerEntradaComandos(hash_usuarios, hash_feed, hash_posts, cuenta_logueada)
}

func LeerUsuariosDesdeArchivo(ruta string) []string {
	archivo, err := os.Open(ruta)
	if err != nil {
		return []string{}
	}
	defer archivo.Close()
	usuarios := []string{}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		linea := strings.TrimSpace(scanner.Text())
		if linea == "" {
			continue
		}
		usuarios = append(usuarios, linea)
	}
	return usuarios
}

func LeerEntradaComandos(hash_users dicci.Diccionario[string, int], hash_feed dicci.Diccionario[string, heap.ColaPrioridad[post]], diccionario_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], logueado usuario) {
	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		linea := strings.TrimSpace(entrada.Text())
		if linea == "" {
			break
		}
		partes := strings.SplitN(linea, " ", 2)
		comando := partes[0]
		parametro := ""
		if len(partes) > 1 {
			parametro = partes[1]
		}

		switch comando {
		case constLogin:
			Login(hash_users, &logueado, parametro)
		case constLogout:
			Logout(&logueado)
		case constPublicar:
			Publicar(diccionario_posts, hash_feed, &logueado, diccionario_posts.Cantidad(), parametro)
		case constConstVerFeed:
			VerProximoPostEnFeed(diccionario_posts, hash_feed, &logueado)
		default:
			id, err := strconv.Atoi(parametro)
			if err != nil {
				fmt.Println("Error: Id no numerico")
			}
			switch comando {
			case constDarLike:
				LikearUnPost(diccionario_posts, &logueado, id)
			case constVerLikes:
				VerLikesDeUnPost(diccionario_posts, id)

			default:
				fmt.Println("Error: Comando Desconocido")
			}

		}
	}
}

func Login(hash_users dicci.Diccionario[string, int], logueado *usuario, usuario_a_loguear string) {
	if logueado.nombre != "" {
		fmt.Println("Error: Ya habia un usuario loggeado")
		return
	}
	if !hash_users.Pertenece(usuario_a_loguear) {
		fmt.Println("Error: usuario no existente")
		return
	}
	logueado.nombre = usuario_a_loguear
	logueado.indice_en_arreglo = hash_users.Obtener(usuario_a_loguear)
	fmt.Println("Hola", usuario_a_loguear)
}

func Logout(logueado *usuario) {
	if logueado.nombre == "" {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	*logueado = usuario{}
	fmt.Println("Adios")
}

func Publicar(diccionario_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], hash_feed dicci.Diccionario[string, heap.ColaPrioridad[post]], logueado *usuario, ultimo_post_id int, mensaje string) {
	if logueado.nombre == "" {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}
	diccionario_posts.Guardar(ultimo_post_id, dicci.CrearABB[string, int](comparacionStrings))
	for iter := hash_feed.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		usuario, cola := iter.VerActual()
		if usuario != logueado.nombre {
			cola.Encolar(post{mensaje: mensaje, id: ultimo_post_id, publicador: *logueado})
		}
	}
	fmt.Println("Post publicado")
}

func VerProximoPostEnFeed(diccionario_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], hash_feed dicci.Diccionario[string, heap.ColaPrioridad[post]], logueado *usuario) {
	if logueado.nombre == "" || hash_feed.Obtener(logueado.nombre).EstaVacia() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}

	feed := hash_feed.Obtener(logueado.nombre)
	p := feed.Desencolar()
	likes := diccionario_posts.Obtener(p.id)
	fmt.Printf("Post ID %d\n", p.id)
	fmt.Printf("%s dijo: %s\n", p.publicador.nombre, p.mensaje)
	fmt.Printf("Likes: %d\n", likes.Cantidad())
}

func LikearUnPost(diccionario_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], logueado *usuario, id int) {
	if logueado.nombre == "" || !diccionario_posts.Pertenece(id) {
		fmt.Println("Error: Usuario no loggeado o Post inexistente")
		return
	}
	abb_likes := diccionario_posts.Obtener(id)
	abb_likes.Guardar(logueado.nombre, 0)
	fmt.Println("Post likeado")
}

func VerLikesDeUnPost(diccionario_posts dicci.Diccionario[int, dicci.DiccionarioOrdenado[string, int]], id int) {
	if !diccionario_posts.Pertenece(id) || diccionario_posts.Obtener(id).Cantidad() == 0 {
		fmt.Println("Error: Post inexistente o sin likes")
		return
	}
	abb_likes := diccionario_posts.Obtener(id)
	fmt.Printf("El post tiene %d likes:\n", abb_likes.Cantidad())
	for iter := abb_likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre, _ := iter.VerActual()
		fmt.Printf("\t%s\n", nombre)
	}
}
