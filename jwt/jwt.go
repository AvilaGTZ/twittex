package jwt

import (
	"time"

	"github.com/AvilaGTZ/twittex/models"
	jwt "github.com/dgrijalva/jwt-go"
)

/*GeneroJWT Genera el encriptado con JWT*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("Grayefx")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		//"ubicacion":        t.Ubicacion,
		"_id": t.ID.Hex(),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(miClave)
	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
