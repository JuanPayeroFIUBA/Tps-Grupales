package main

import (
	algo "algogram"
)

func main() {
	usuarios := algo.LeerEntradaArchivoYGuardar()
	algo.Setup(usuarios)

}
