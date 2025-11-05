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
	id      int
	creador *usuario
	mensaje string
	likes   dicci.DiccionarioOrdenado[string, int]
}

// Funcion de igualdad usada por el diccionario de usuarios para comparar claves
func igualdadStrings(a, b string) bool {
	return a == b
}

// Funcion de igualdad usada por el diccionario de posteos para comparar claves
func igualdadInts(a, b int) bool {
	return a == b
}

// Funcion de comparacion usada por la cola de prioridad que representa el feed para comparar y dar orden a la cola
func comparacionStrings(a, b string) int {
	return strings.Compare(a, b)
}

// Funcion encargada de leer el archivo proporcionado al llamar al programa y procesar la lista de usuarios que venga en este, retorna una lista de strings que representa a los usuarios disponibles
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

// Funcion que se encarga de inicializar los TDAS a utilizar por el programa, posteriormente llama a la lectura de los comandos por entrada estandar
func InicializarYProcesarEntradas(usuarios []string) {
	var cuenta_loggeada *usuario
	diccionario_usuarios := dicci.CrearHash[string, *usuario](igualdadStrings)
	diccionario_posts := dicci.CrearHash[int, *post](igualdadInts)

	//podriamos hacer esto al leerlo mismo, pero quedaria todo junto en una funcion y se veria todo muy amontonado
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
		diccionario_usuarios.Guardar(nombre_usuario, u)
	}

	LeerComandosYProcesar(diccionario_usuarios, diccionario_posts, cuenta_loggeada)
}

// Funcion que se encarga de permitir y leer la entrada de comandos por entrada estandar, asimismo procesa estos datos y llama a la funcion de comando correspondiente en cada caso
func LeerComandosYProcesar(diccionario_usuarios dicci.Diccionario[string, *usuario], diccionario_posts dicci.Diccionario[int, *post], loggeado *usuario) {
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
			loggeado = ComandoLogin(diccionario_usuarios, loggeado, parametro)
		case constLogout:
			loggeado = ComandoLogout(loggeado)
		case constPublicar:
			ComandoPublicar(diccionario_usuarios, diccionario_posts, loggeado, diccionario_posts.Cantidad(), parametro)
		case constConstVerFeed:
			ComandoVerProximoPostEnFeed(loggeado)
		default:
			ComandosDeLikes(comando, parametro, diccionario_posts, loggeado)
		}
	}
}

// Funcion encargada de procesar los comandos de likes, los cuales requieren una conversion de tipos primero
func ComandosDeLikes(comando, parametro string, diccionario_posts dicci.Diccionario[int, *post], loggeado *usuario) {
	id, err := strconv.Atoi(parametro)
	if err != nil {
		fmt.Println("Error: Id no numerico")
	}
	switch comando {
	case constDarLike:
		ComandoLikear(diccionario_posts, loggeado, id)
	case constVerLikes:
		ComandoVerLikes(diccionario_posts, id)

	default:
		fmt.Println("Error: Comando Desconocido")
	}
}

// Funcion que ejecuta el comando "Login" del programa, retornando al usuario para ser guardado como loggeado, en caso de ya existir un usuario loggeado o de que este no pertenezca a la lista de usuarios, se imprimira un Error
func ComandoLogin(diccionario_usuarios dicci.Diccionario[string, *usuario], loggeado *usuario, usuario_a_loguear string) *usuario {
	if loggeado != nil {
		fmt.Println("Error: Ya habia un usuario loggeado")
		return loggeado
	}
	if !diccionario_usuarios.Pertenece(usuario_a_loguear) {
		fmt.Println("Error: usuario no existente")
		return nil
	}
	loggeado = diccionario_usuarios.Obtener(usuario_a_loguear)
	fmt.Println("Hola", usuario_a_loguear)
	return loggeado
}

// Funcion que ejecuta el comando "Logout" del programa, retornando nil para guardar al usuario loggeado como inexistente, en caso de no haber un usuario loggeado, se imprimira un Error
func ComandoLogout(loggeado *usuario) *usuario {
	if loggeado == nil {
		fmt.Println("Error: no habia usuario loggeado")
		return nil
	}
	fmt.Println("Adios")
	return nil
}

// Funcion que ejecuta el comando "Publicar" del programa, guardando al posteo en el feed de cada usuario excepto el loggeado y el el diccionario de posteos, en caso de no haber un usuario loggeado, se imprimira un Error
func ComandoPublicar(diccionario_usuarios dicci.Diccionario[string, *usuario], diccionario_posts dicci.Diccionario[int, *post], loggeado *usuario, ultimo_post_id int, mensaje string) {
	if loggeado == nil {
		fmt.Println("Error: no habia usuario loggeado")
		return
	}

	nuevo_post := &post{
		id:      ultimo_post_id,
		creador: loggeado,
		mensaje: mensaje,
		likes:   dicci.CrearABB[string, int](comparacionStrings),
	}
	diccionario_posts.Guardar(ultimo_post_id, nuevo_post)

	// Agregar el post al feed de todos los usuarios excepto el creador
	for iter := diccionario_usuarios.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		nombre_usuario, usuario := iter.VerActual()
		if nombre_usuario != loggeado.nombre {
			usuario.feed.Encolar(nuevo_post)
		}
	}
	fmt.Println("Post publicado")
}

// Funcion que ejecuta el comando "Ver el Proximo Post del Feed" del programa, mostrando el proximo post en el feed del usuario loggeado asi como informacion adicional, en caso de no haber un usuario loggeado o que ya no queden posts en el feed, se imprimira un Error
func ComandoVerProximoPostEnFeed(loggeado *usuario) {
	if loggeado == nil || loggeado.feed.EstaVacia() {
		fmt.Println("Usuario no loggeado o no hay mas posts para ver")
		return
	}

	p := loggeado.feed.Desencolar()
	fmt.Printf("Post ID %d\n", p.id)
	fmt.Printf("%s dijo: %s\n", p.creador.nombre, p.mensaje)
	fmt.Printf("Likes: %d\n", p.likes.Cantidad())
}

// Funcion que ejecuta el comando "Dar Like" del programa, guardando un "like" en ese posteo asi como el nombre del usuario que dio el like (el usuario loggeado), en caso de no haber un usuario loggeado o que el indice del post no exista, se imprimira un Error
func ComandoLikear(diccionario_posts dicci.Diccionario[int, *post], loggeado *usuario, id int) {
	if loggeado == nil || !diccionario_posts.Pertenece(id) {
		fmt.Println("Error: Usuario no loggeado o Post inexistente")
		return
	}
	p := diccionario_posts.Obtener(id)
	p.likes.Guardar(loggeado.nombre, 0)
	fmt.Println("Post likeado")
}

// Funcion que ejecuta el comando "Ver Likes" del programa, permitiendo ver los likes que hayan dado los usuarios a ese posteo asi como sus nombres, en caso de que el post no tenga likes o que el indice del post no exista, se imprimira un Error
func ComandoVerLikes(diccionario_posts dicci.Diccionario[int, *post], id int) {
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
