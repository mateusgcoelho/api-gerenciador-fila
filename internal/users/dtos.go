package users

import "errors"

type CreateUserDto struct {
	Nome           string `json:"nome"`
	Email          string `json:"email"`
	Telefone       string `json:"telefone"`
	Cpf            string `json:"cpf"`
	Senha          string `json:"senha"`
	CodigoRegistro string `json:"codigoRegistro"`
}

type LoginDto struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

func (u *CreateUserDto) Validate() error {
	if u.Email == "" ||
		u.CodigoRegistro == "" ||
		u.Senha == "" ||
		u.Telefone == "" ||
		u.Nome == "" {
		return errors.New("Verifique os campos de envio.")
	}
	return nil
}
