package reports

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IReportDao interface {
	createReport(data CreateReportDto) (*Report, error)
	getReportById(id int) (*Report, error)
}

func NewReportDao(dbPool *pgxpool.Pool) IReportDao {
	return &reportDao{
		dbPool: dbPool,
	}
}

type reportDao struct {
	dbPool *pgxpool.Pool
}

func (r *reportDao) createReport(data CreateReportDto) (*Report, error) {
	if data.PessoaId == data.ResponsavelId {
		return nil, errors.New("Não se pode criar atendimento para si mesmo.")
	}

	if err := verifyIfUsersExists(r.dbPool, data.PessoaId, data.ResponsavelId); err != nil {
		return nil, err
	}

	query := `
		INSERT INTO atendimentos (pessoa_id, responsavel_id, senha) VALUES ($1, $2, $3) RETURNING id
	`

	rows, err := r.dbPool.Query(
		context.Background(),
		query,
		data.PessoaId, data.ResponsavelId, data.Senha,
	)
	if err != nil {
		return nil, errors.New("Não foi possível criar atendimento.")
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, errors.New("Não foi possível encontrar atendimento criado.")
		}
	}

	return r.getReportById(id)
}

func (r *reportDao) getReportById(id int) (*Report, error) {
	query := `
		SELECT id, pessoa_id, responsavel_id, senha, data_finalizacao, data_criacao, data_atualizacao
		FROM atendimentos WHERE = $1
	`

	rows, err := r.dbPool.Query(
		context.Background(),
		query,
		id,
	)
	if err != nil {
		return nil, errors.New("Não foi possível encontrar atendimento.")
	}

	var report *Report = nil
	for rows.Next() {
		if err := rows.Scan(
			&report.Id,
			&report.PessoaId,
			&report.ResponsavelId,
			&report.Senha,
			&report.DataFinalizacao,
			&report.DataCriacao,
			&report.DataAtualizacao,
		); err != nil {
			return nil, errors.New("Não foi possível encontrar atendimento criado.")
		}
	}

	return r.getReportById(id)
}

func verifyIfUsersExists(dbPool *pgxpool.Pool, usersId ...int) error {
	query := `
		SELECT id FROM usuarios WHERE id IN ($1)
	`

	var userIdsStr []string
	for _, id := range usersId {
		userIdsStr = append(userIdsStr, strconv.Itoa(id))
	}

	whereIds := strings.Join(userIdsStr, ",")

	rows, err := dbPool.Query(
		context.Background(),
		query,
		whereIds,
	)
	if err != nil {
		return errors.New("Não foi possível validar usuários.")
	}

	values, err := rows.Values()
	if err != nil {
		return errors.New("Não foi possível validar usuários.")
	}

	if len(values) != len(usersId) {
		return errors.New("Não foi possível encontrar usuário, verifique os ids.")
	}

	return nil
}
