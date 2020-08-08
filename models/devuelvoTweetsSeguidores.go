package models

import "go.mongodb.org/mongo-driver/bson/primitive"

/*DevuelvoTweetsSeguidores es el modelo que ocntendra los tweets de mis seguidores*/
type DevuelvoTweetsSeguidores struct {
	ID                primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UsuarioID         string             `bson:"usuarioid" json:"userId,omitempty"`
	UsuarioRelacionID string             `bson:"usuariorelacionid" json:"userRelationId,omitempty"`
	Tweet             struct {
		Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
		Fecha   string `bson:"fecha" json:"fecha,omitempty"`
		ID      string `bson:"_id" json:"_id,omitempty"`
	}
}
