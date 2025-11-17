package app_algogram

import "strings"

func compararStrings(a, b string) int { return strings.Compare(a, b) }

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

const (
	cmdLogin    = "login"
	cmdLogout   = "logout"
	cmdPublicar = "publicar"
	cmdVerFeed  = "ver_siguiente_feed"
	cmdLikear   = "likear_post"
	cmdVerLikes = "mostrar_likes"
)
