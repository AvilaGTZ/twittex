package middlew

import (
	"net/http"

	"github.com/AvilaGTZ/twittex/bd"
)

/*ChequeoBD Es el middleware que se encarga de verificar la conexcion con la base de datos*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
