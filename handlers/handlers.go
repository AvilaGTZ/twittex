package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/AvilaGTZ/twittex/middlew"

	"github.com/AvilaGTZ/twittex/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores seteo el puerto por el cual escuchara mi API y ademas agrego como middleware a CORS*/
func Manejadores() {
	router := mux.NewRouter()

	/*Primer handler creado*/
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "9090"
	}
	PORT = "9090"
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
