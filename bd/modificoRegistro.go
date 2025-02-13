package bd

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/AvilaGTZ/twittex/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ModificoRegistro permite modificar el perfil del usuario*/
func ModificoRegistro(u models.Usuario, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}
	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos
	}
	registro["fechaNacimiento"] = u.FechaNacimiento
	if len(u.Avatar) > 0 {
		registro["avatar"] = u.Avatar
	}
	if len(u.Banner) > 0 {
		registro["banner"] = u.Banner
	}
	if len(u.Biografia) > 0 {
		registro["biografia"] = u.Biografia
	}
	//if len(u.Ubicacion) > 0 {
	//	registro["ubicacion"] = u.Ubicacion
	//}
	if len(u.SitioWeb) > 0 {
		registro["sitioWeb"] = u.SitioWeb
	}

	updateString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}

	return true, nil

}
