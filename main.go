package main

import (
	"log"

	"github.com/AvilaGTZ/twittex/bd"
	"github.com/AvilaGTZ/twittex/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
