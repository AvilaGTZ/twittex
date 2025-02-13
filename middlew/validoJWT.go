package middlew

import (
	"net/http"

	"github.com/AvilaGTZ/twittex/routers"
)

/*ValidoJWT Permite validar el JWT Que nos viene en la aplicacion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Error en el Token "+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
