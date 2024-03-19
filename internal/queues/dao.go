package queues

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IQueueDao interface {
	createQueue(data CreateQueueDto) (*Queue, error)
	getQueueById(id int) (*Queue, error)
	getQueues() ([]Queue, error)
}

type queueDao struct {
	dbPool *pgxpool.Pool
}

func NewQueueDao(dbPool *pgxpool.Pool) IQueueDao {
	return queueDao{
		dbPool: dbPool,
	}
}

func (r queueDao) createQueue(data CreateQueueDto) (*Queue, error) {
	query := `
		INSERT INTO filas (nome, senha_atual) VALUES ($1, $2) RETURNING id
	`

	rows, err := r.dbPool.Query(
		context.Background(),
		query,
		data.Nome, 0,
	)
	if err != nil {
		return nil, errors.New("Não foi possível criar fila.")
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, errors.New("Não foi possível encontrar fila criada.")
		}
	}

	return r.getQueueById(id)
}

func (r queueDao) getQueueById(id int) (*Queue, error) {
	sql := `
		SELECT id, nome, senha_atual, data_criacao, data_atualizacao
		FROM filas AS f WHERE f.id = $1
	`

	var queue *Queue = nil
	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
		id,
	)
	if err != nil {
		return nil, errors.New("Não foi possível encontrar fila.")
	}
	defer rows.Close()

	for rows.Next() {
		queue = &Queue{}
		if err := rows.Scan(
			&queue.Id,
			&queue.Nome,
			&queue.SenhaAtual,
			&queue.DataCriacao,
			&queue.DataAtualizacao,
		); err != nil {
			return nil, err
		}
	}

	return queue, nil
}

func (r queueDao) getQueues() ([]Queue, error) {
	sql := `
		SELECT id, nome, senha_atual, data_criacao, data_atualizacao
		FROM filas
	`

	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
	)
	if err != nil {
		return nil, errors.New("Não foi possível encontrar filas.")
	}
	defer rows.Close()

	queues := []Queue{}

	for rows.Next() {
		queue := Queue{}
		if err := rows.Scan(
			&queue.Id,
			&queue.Nome,
			&queue.SenhaAtual,
			&queue.DataCriacao,
			&queue.DataAtualizacao,
		); err != nil {
			return nil, err
		}

		queues = append(queues, queue)
	}

	return queues, nil
}
