package users

type User struct {
	Nome           string `json:"nome"`
	Email          string `json:"email"`
	Senha          string `json:"senha"`
	CodigoRegistro string `json:"codigo_registro"`
}
