package app_algogram

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LeerUsuariosDesdeArchivo(ruta string) []string {
	f, err := os.Open(ruta)
	if err != nil {
		return []string{}
	}
	defer f.Close()
	var usuarios []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		linea := strings.TrimSpace(sc.Text())
		if linea == "" {
			continue
		}
		usuarios = append(usuarios, linea)
	}
	return usuarios
}

func procesarEntrada(a Algogram) {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		linea := strings.TrimSpace(sc.Text())
		if linea == "" {
			break
		}
		comando, parametro := parsearLinea(linea)
		switch comando {
		case cmdLogin:
			out := a.Login(parametro)
			if out != "" {
				fmt.Println(out)
			}
		case cmdLogout:
			out := a.Logout()
			if out != "" {
				fmt.Println(out)
			}
		case cmdPublicar:
			out := a.Publicar(parametro)
			if out != "" {
				fmt.Println(out)
			}
		case cmdVerFeed:
			out := a.VerSiguienteFeed()
			if out != "" {
				fmt.Println(out)
			}
		case cmdLikear:
			id, err := strconv.Atoi(parametro)
			if err != nil {
				fmt.Println("Error: Id no numerico")
				continue
			}
			out := a.LikearPost(id)
			if out != "" {
				fmt.Println(out)
			}
		case cmdVerLikes:
			id, err := strconv.Atoi(parametro)
			if err != nil {
				fmt.Println("Error: Id no numerico")
				continue
			}
			out := a.MostrarLikes(id)
			if out != "" {
				fmt.Println(out)
			}
		default:
			fmt.Println("Error: Comando Desconocido")
		}
	}
}

// Lee comandos por stdin y los ejecuta contra el TDA Algogram
func ProcesarEntrada(a Algogram) { procesarEntrada(a) }

func parsearLinea(linea string) (string, string) {
	idx := strings.Index(linea, " ")
	if idx == -1 {
		return linea, ""
	}
	return linea[:idx], linea[idx+1:]
}
