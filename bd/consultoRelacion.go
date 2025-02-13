package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/AvilaGTZ/twittex/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion verifica la relacion entre dos usuarios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twittor")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	fmt.Println(condicion)
	if err != nil {
		//fmt.Println(err.Error())
		return false, err
	}

	return true, nil
}
