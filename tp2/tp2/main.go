package main

import (
	algo "app_algogram"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	usuarios := algo.LeerUsuariosDesdeArchivo(os.Args[1])
	ag := algo.NewAlgogram(usuarios)
	algo.ProcesarEntrada(ag)

}
