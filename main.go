package main

import (
	"log"

	"github.com/kevin1892/twetter/bd"
	"github.com/kevin1892/twetter/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
