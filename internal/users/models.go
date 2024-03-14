package users

import "time"

type User struct {
	Id              int       `json:"id"`
	Nome            string    `json:"nome"`
	Email           string    `json:"email"`
	Senha           string    `json:"-"`
	CodigoRegistro  string    `json:"codigoRegistro"`
	Permissoes      int       `json:"-"`
	DataCriacao     time.Time `json:"dataCriacao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
}
