package app_algogram

import (
	"fmt"
	dicc "tdas/diccionario"
)

type postImpl struct {
	id      int
	autor   Usuario
	mensaje string
	likes   dicc.DiccionarioOrdenado[string, int]
}

func crearPost(id int, autor Usuario, mensaje string) Post {
	return &postImpl{
		id:      id,
		autor:   autor,
		mensaje: mensaje,
		likes:   dicc.CrearABB[string, int](compararStrings),
	}
}

func (p *postImpl) ID() int              { return p.id }
func (p *postImpl) Autor() Usuario       { return p.autor }
func (p *postImpl) Mensaje() string      { return p.mensaje }
func (p *postImpl) CantidadLikes() int   { return p.likes.Cantidad() }
func (p *postImpl) Likear(nombre string) { p.likes.Guardar(nombre, 0) }

func (p *postImpl) Imprimir() {
	fmt.Printf("Post ID %d\n", p.id)
	fmt.Printf("%s dijo: %s\n", p.autor.Nombre(), p.mensaje)
	fmt.Printf("Likes: %d\n", p.likes.Cantidad())
}

func (p *postImpl) MostrarLikes() {
	fmt.Printf("El post tiene %d likes:\n", p.likes.Cantidad())
	for it := p.likes.Iterador(); it.HaySiguiente(); it.Siguiente() {
		nombre, _ := it.VerActual()
		fmt.Printf("\t%s\n", nombre)
	}
}
