package reports

type CreateReportDto struct {
	PessoaId      int `json:"pessoaId"`
	ResponsavelId int `json:"responsavelId"`
	FilaId        int `json:"filaId"`
}
