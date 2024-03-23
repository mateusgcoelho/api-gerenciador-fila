package reports

type CreateReportDto struct {
	PessoaId      int `json:"pessoaId"`
	ResponsavelId int `json:"responsavelId"`
	Senha         int `json:"senha"`
}
