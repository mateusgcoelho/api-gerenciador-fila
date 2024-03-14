package users

type CreateUserDto struct {
	Nome           string `json:"nome"`
	Email          string `json:"email"`
	Senha          string `json:"senha"`
	CodigoRegistro string `json:"codigoRegistro"`
}

type LoginDto struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}
