package models

/*RespuestaLogin tiene el token que se devuelve con login*/
type RespuestaLogin struct {
	Token string `json:"token, omitempty"`
}
