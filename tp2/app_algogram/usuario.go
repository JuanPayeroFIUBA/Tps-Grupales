package app_algogram

// Interfaz TDA Usuario
type Usuario interface {
	Nombre() string
	Indice() int
	EncolarPost(p Post)
	SiguientePost() (Post, bool)
	FeedVacio() bool
}
