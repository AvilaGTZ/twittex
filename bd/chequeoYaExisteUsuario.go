package bd

import (
	"context"
	"time"

	"github.com/AvilaGTZ/twittex/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ChequeoYaExisteUsuario revisa si el email que tratan de registrar ya existe en la base de datos*/
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {

		return resultado, false, ID
	}

	return resultado, true, ID
}
