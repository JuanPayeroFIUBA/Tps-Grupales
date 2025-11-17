package app_algogram

type Algogram interface {
	Login(nombre string) string
	Logout() string
	Publicar(mensaje string) string
	VerSiguienteFeed() string
	LikearPost(id int) string
	MostrarLikes(id int) string
}
