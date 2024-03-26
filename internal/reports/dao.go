package reports

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IReportDao interface {
	createReport(data CreateReportDto) (*Report, error)
	getReportById(id int) (*Report, error)
	verifyIfUsersExists(usersId []int) error
	getReports() ([]Report, error)
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

	if err := r.verifyIfUsersExists([]int{data.PessoaId, data.ResponsavelId}); err != nil {
		return nil, err
	}

	query := `
		SELECT senha FROM filas WHERE id = $1
	`

	rows, err := r.dbPool.Query(context.Background(), query, data.FilaId)
	fmt.Println(data.FilaId)
	var senha *int
	for rows.Next() {
		if err := rows.Scan(&senha); err != nil {
			return nil, err
		}
	}

	if senha == nil {
		return nil, errors.New("Não foi possível encontrar a fila para atendimento.")
	}

	query = `
		SELECT id FROM atendimentos
		WHERE data_finalizacao IS NULL AND (pessoa_id = $1 OR responsavel_id = $2) FOR UPDATE
	`
	var reportIdFounded *int
	rows, err = r.dbPool.Query(
		context.Background(),
		query,
		data.PessoaId, data.ResponsavelId,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&reportIdFounded); err != nil {
			return nil, err
		}
	}

	if reportIdFounded != nil {
		return nil, errors.New("Existe um atendimento em andamento.")
	}

	query = `
		INSERT INTO atendimentos (pessoa_id, responsavel_id, senha) VALUES ($1, $2, $3) RETURNING id
	`

	rows, err = r.dbPool.Query(
		context.Background(),
		query,
		data.PessoaId, data.ResponsavelId, (*senha + 1),
	)
	defer rows.Close()

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
		FROM atendimentos WHERE id = $1
	`

	rows, err := r.dbPool.Query(
		context.Background(),
		query,
		id,
	)
	defer rows.Close()

	if err != nil {
		return nil, errors.New("Não foi possível encontrar atendimento.")
	}

	var report *Report = nil
	for rows.Next() {
		report = &Report{}
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

	return report, nil
}

func (r *reportDao) verifyIfUsersExists(usersId []int) error {
	query := `
		SELECT id FROM usuarios WHERE id IN ($1)
	`

	var userIdsStr []string
	for _, id := range usersId {
		userIdsStr = append(userIdsStr, strconv.Itoa(id))
	}

	whereIds := strings.Join(userIdsStr, ",")

	rows, err := r.dbPool.Query(
		context.Background(),
		query,
		whereIds,
	)
	if err != nil {
		return errors.New("Não foi possível validar usuários.")
	}
	defer rows.Close()

	var usersIdFoundeds []int = []int{}
	for rows.Next() {
		var id int
		err := rows.Scan(&id)

		if err != nil {
			return errors.New("Não foi possível validar usuários.")
		}

		usersIdFoundeds = append(usersIdFoundeds, id)
	}

	if !(len(usersIdFoundeds) != len(usersId)) {
		return errors.New("Nem todos os usuários foram encontrados.")
	}

	return nil
}

func (r *reportDao) getReports() ([]Report, error) {
	query := `
		SELECT id, senha, pessoa_id, responsavel_id, data_criacao, data_finalizacao, data_atualizacao
		FROM atendimentos
	`

	rows, err := r.dbPool.Query(context.Background(), query)
	if err != nil {
		return nil, errors.New("Não foi possível realizar busca de atendimentos.")
	}
	defer rows.Close()

	var reports []Report = []Report{}

	for rows.Next() {
		report := Report{}

		err := rows.Scan(
			&report.Id,
			&report.Senha,
			&report.PessoaId,
			&report.ResponsavelId,
			&report.DataCriacao,
			&report.DataFinalizacao,
			&report.DataAtualizacao,
		)
		if err != nil {
			return nil, errors.New("Não foi possível realizar busca de atendimentos.")
		}

		reports = append(reports, report)
	}

	return reports, nil
}
