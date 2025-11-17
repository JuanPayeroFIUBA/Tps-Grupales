package app_algogram

import (
	"fmt"
	dicc "tdas/diccionario"
)

type algogramImpl struct {
	usuarios dicc.Diccionario[string, Usuario]
	posts    dicc.Diccionario[int, Post]
	actual   Usuario
	proximo  int
}

func crearAlgogram(nombres []string) Algogram {
	usuarios := dicc.CrearHash[string, Usuario](func(a, b string) bool { return a == b })
	posts := dicc.CrearHash[int, Post](func(a, b int) bool { return a == b })

	for i, nombre := range nombres {
		u := crearUsuario(nombre, i)
		usuarios.Guardar(nombre, u)
	}

	return &algogramImpl{usuarios: usuarios, posts: posts, actual: nil, proximo: 0}
}

// NewAlgogram crea la instancia del TDA Algogram a partir de la lista de usuarios
func NewAlgogram(usuarios []string) Algogram { return crearAlgogram(usuarios) }

func (a *algogramImpl) Login(nombre string) string {
	if a.actual != nil {
		return "Error: Ya habia un usuario loggeado"
	}
	if !a.usuarios.Pertenece(nombre) {
		return "Error: usuario no existente"
	}
	a.actual = a.usuarios.Obtener(nombre)
	return fmt.Sprintf("Hola %s", nombre)
}

func (a *algogramImpl) Logout() string {
	if a.actual == nil {
		return "Error: no habia usuario loggeado"
	}
	a.actual = nil
	return "Adios"
}

func (a *algogramImpl) Publicar(mensaje string) string {
	if a.actual == nil {
		return "Error: no habia usuario loggeado"
	}
	p := crearPost(a.proximo, a.actual, mensaje)
	a.posts.Guardar(a.proximo, p)
	a.proximo++

	a.usuarios.Iterar(func(nombre string, u Usuario) bool {
		if a.actual != nil && nombre != a.actual.Nombre() {
			u.EncolarPost(p)
		}
		return true
	})
	return "Post publicado"
}

func (a *algogramImpl) VerSiguienteFeed() string {
	if a.actual == nil || a.actual.FeedVacio() {
		return "Usuario no loggeado o no hay mas posts para ver"
	}
	p, ok := a.actual.SiguientePost()
	if !ok {
		return "Usuario no loggeado o no hay mas posts para ver"
	}
	p.Imprimir()
	return ""
}

func (a *algogramImpl) LikearPost(id int) string {
	if a.actual == nil || !a.posts.Pertenece(id) {
		return "Error: Usuario no loggeado o Post inexistente"
	}
	p := a.posts.Obtener(id)
	p.Likear(a.actual.Nombre())
	return "Post likeado"
}

func (a *algogramImpl) MostrarLikes(id int) string {
	if !a.posts.Pertenece(id) {
		return "Error: Post inexistente o sin likes"
	}
	p := a.posts.Obtener(id)
	if p.CantidadLikes() == 0 {
		return "Error: Post inexistente o sin likes"
	}
	p.MostrarLikes()
	return ""
}
