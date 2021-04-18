package entity

// PostData estructura que define los campos que tendrá la solicitud y respuesta
type PostData struct {
	Id       int64  `json:"id"`
	Title    string `json:"title"`
	Messagge string `json:"messagge"`
}
