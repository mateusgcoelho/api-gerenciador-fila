package queues

import "time"

type Queue struct {
	Id              int       `json:"id"`
	Nome            string    `json:"nome"`
	SenhaAtual      int       `json:"senhaAtual"`
	DataCriacao     time.Time `json:"dataCriacao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
}
