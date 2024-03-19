package users

import (
	"time"

	"github.com/mateusgcoelho/api-gerenciador-fila/internal/persons"
)

type User struct {
	Id              int            `json:"id"`
	Email           string         `json:"email"`
	CodigoRegistro  string         `json:"codigoRegistro"`
	Senha           string         `json:"-"`
	Permissoes      int            `json:"-"`
	Pessoa          persons.Person `json:"pessoa"`
	DataCriacao     time.Time      `json:"dataCriacao"`
	DataAtualizacao time.Time      `json:"dataAtualizacao"`
}
