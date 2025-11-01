package main

import (
	algo "algogram"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	usuarios := algo.LeerUsuariosDesdeArchivo(os.Args[1])
	algo.Setup(usuarios)

}
