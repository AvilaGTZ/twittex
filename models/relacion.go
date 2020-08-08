package models

/*Relacion define la relacion entre dos usuarios a los que sigo*/
type Relacion struct {
	/*Mi usuario que se guardara al momento de seguir a otro*/
	UsuarioID string `bson:"usuarioid" json:"usuarioId"`

	/*El usuario al que seguire*/
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacionId"`
}
