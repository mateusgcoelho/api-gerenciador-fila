package persons

import "time"

type Person struct {
	Id              int       `json:"id"`
	Nome            string    `json:"nome"`
	Telefone        *string   `json:"telefone"`
	Cpf             *string   `json:"cpf"`
	DataCriacao     time.Time `json:"dataCriacao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
}
