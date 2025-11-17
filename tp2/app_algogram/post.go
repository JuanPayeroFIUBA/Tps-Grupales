package app_algogram

// Interfaz TDA Post
type Post interface {
	ID() int
	Autor() Usuario
	Mensaje() string
	CantidadLikes() int
	Likear(nombreUsuario string)
	Imprimir()
	MostrarLikes()
}
