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
	feed              heap.ColaPrioridad[*post]
}

type post struct {
	id         int
	creador    *usuario
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
	hash_usuarios := dicci.CrearHash[string, *usuario](igualdadStrings)
	hash_posts := dicci.CrearHash[int, *post](igualdadInts)

	for indice, nombre_usuario := range usuarios {
		u := &usuario{
			nombre:            nombre_usuario,
			indice_en_arreglo: indice,
			feed: heap.CrearHeap(func(post1, post2 *post) int {
				indice_usuario1 := int(math.Abs(float64(post1.creador.indice_en_arreglo - indice)))
				indice_usuario2 := int(math.Abs(float64(post2.creador.indice_en_arreglo - indice)))
				if indice_usuario1 == indice_usuario2 {
					return -(post1.id - post2.id)
				}
				return -(indice_usuario1 - indice_usuario2)
			}),
		}
		hash_usuarios.Guardar(nombre_usuario, u)
	}

	var cuenta_logueada *usuario
	LeerEntradaComandos(hash_usuarios, hash_posts, cuenta_logueada)
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

func LeerEntradaComandos(hash_users dicci.Diccionario[string, *usuario], diccionario_posts dicci.Diccionario[int, *post], logueado *usuario) {
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
			logueado = Login(hash_users, logueado, parametro)
		case constLogout:
			logueado = Logout(logueado)
		case constPublicar:
			Publicar(hash_users, diccionario_posts, logueado, diccionario_posts.Cantidad(), parametro)
		case constConstVerFeed:
			VerProximoPostEnFeed(logueado)
		default:
			id, err := strconv.Atoi(parametro)
			if err != nil {
				fmt.Println("Error: Id no numerico")
			}
			switch comando {
			case constDarLike:
				LikearUnPost(diccionario_posts, logueado, id)
			case constVerLikes:
				VerLikesDeUnPost(diccionario_posts, id)

			default:
				fmt.Println("Error: Comando Desconocido")
			}

		}
	}
}

func Login(hash_users dicci.Diccionario[string, *usuario], logueado *usuario, usuario_a_loguear string) *usuario {
	if logueado != nil {
		fmt.Println("Error: Ya habia un usuario loggeado")
		return logueado
	}
	if !hash_users.Pertenece(usuario_a_loguear) {
		fmt.Println("Error: usuario no existente")
		return nil
	}
	logueado = hash_users.Obtener(usuario_a_loguear)
	fmt.Println("Hola", usuario_a_loguear)
	return logueado
}

func Logout(logueado *usuario) *usuario {
	if logueado == nil {
		fmt.Println("Error: no habia usuario loggeado")
		return nil
	}
	fmt.Println("Adios")
	return nil
}

func Publicar(hash_users dicci.Diccionario[string, *usuario], diccionario_posts dicci.Diccionario[int, *post], logueado *usuario, ultimo_post_id int, mensaje string) {
	if logueado == nil {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}

	nuevo_post := &post{
		id:         ultimo_post_id,
		creador: logueado,
		mensaje:    mensaje,
		likes:      dicci.CrearABB[string, int](comparacionStrings),
	}
	diccionario_posts.Guardar(ultimo_post_id, nuevo_post)

	// Agregar el post al feed de todos los usuarios excepto el creador
	for iter := hash_users.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre_usuario, usuario := iter.VerActual()
		if nombre_usuario != logueado.nombre {
			usuario.feed.Encolar(nuevo_post)
		}
	}
	fmt.Println("Post publicado")
}

func VerProximoPostEnFeed(logueado *usuario) {
	if logueado == nil || logueado.feed.EstaVacia() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}

	p := logueado.feed.Desencolar()
	fmt.Printf("Post ID %d\n", p.id)
	fmt.Printf("%s dijo: %s\n", p.creador.nombre, p.mensaje)
	fmt.Printf("Likes: %d\n", p.likes.Cantidad())
}

func LikearUnPost(diccionario_posts dicci.Diccionario[int, *post], logueado *usuario, id int) {
	if logueado == nil || !diccionario_posts.Pertenece(id) {
		fmt.Println("Error: Usuario no loggeado o Post inexistente")
		return
	}
	p := diccionario_posts.Obtener(id)
	p.likes.Guardar(logueado.nombre, 0)
	fmt.Println("Post likeado")
}

func VerLikesDeUnPost(diccionario_posts dicci.Diccionario[int, *post], id int) {
	if !diccionario_posts.Pertenece(id) || diccionario_posts.Obtener(id).likes.Cantidad() == 0 {
		fmt.Println("Error: Post inexistente o sin likes")
		return
	}
	p := diccionario_posts.Obtener(id)
	fmt.Printf("El post tiene %d likes:\n", p.likes.Cantidad())
	for iter := p.likes.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre, _ := iter.VerActual()
		fmt.Printf("\t%s\n", nombre)
	}
}
