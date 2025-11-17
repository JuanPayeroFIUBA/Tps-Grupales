package app_algogram

import (
	heap "tdas/cola_prioridad"
)

type usuarioImpl struct {
	nombre string
	indice int
	feed   heap.ColaPrioridad[Post]
}

func crearUsuario(nombre string, indice int) Usuario {
	// se prioriza menor distancia al indice. SI hay empate, menor ID primero
	cmp := func(a, b Post) int {
		indiceUsuario1 := absInt(a.Autor().Indice() - indice)
		indiceUsuario2 := absInt(b.Autor().Indice() - indice)
		if indiceUsuario1 == indiceUsuario2 {
			if a.ID() == b.ID() {
				return 0
			}
			if a.ID() < b.ID() {
				return 1
			}
			return -1
		}
		if indiceUsuario1 < indiceUsuario2 {
			return 1
		}
		return -1
	}
	return &usuarioImpl{
		nombre: nombre,
		indice: indice,
		feed:   heap.CrearHeap(cmp),
	}
}

func (u *usuarioImpl) Nombre() string { return u.nombre }
func (u *usuarioImpl) Indice() int    { return u.indice }

func (u *usuarioImpl) EncolarPost(p Post) {
	u.feed.Encolar(p)
}

func (u *usuarioImpl) SiguientePost() (Post, bool) {
	if u.feed.EstaVacia() {
		return nil, false
	}
	return u.feed.Desencolar(), true
}

func (u *usuarioImpl) FeedVacio() bool { return u.feed.EstaVacia() }
