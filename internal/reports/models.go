package reports

import "time"

type Report struct {
	Id              int       `json:"id"`
	Senha           int       `json:"Senha"`
	PessoaId        int       `json:"pessoaId"`
	ResponsavelId   int       `json:"responsavelId"`
	DataCriacao     time.Time `json:"dataCriacao"`
	DataFinalizacao time.Time `json:"dataFinalizacao"`
	DataAtualizacao time.Time `json:"dataAtualizacao"`
}
