package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/AvilaGTZ/twittex/bd"
	"github.com/AvilaGTZ/twittex/models"
)

/*GraboTweet permite grabar el tweet en la base de datos*/
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	resgistro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(resgistro)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha logrado insertar el tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
