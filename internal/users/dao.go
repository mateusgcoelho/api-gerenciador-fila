package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/persons"
)

type IUserDao interface {
	createUser(data CreateUserDto) (*User, error)
	getUsers() ([]User, error)
	getUserById(id int) (*User, error)
	getUser(code string, email string, phone string, cpf string) (*User, error)
}

type userDao struct {
	dbPool  *pgxpool.Pool
	authDao auth.IAuthDao
}

func NewUserDao(dbPool *pgxpool.Pool, authDao auth.IAuthDao) IUserDao {
	return userDao{
		dbPool:  dbPool,
		authDao: authDao,
	}
}

func (r userDao) createUser(data CreateUserDto) (*User, error) {
	user, err := r.getUser(data.CodigoRegistro, data.Email, data.Telefone, data.Cpf)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Usuário já cadastrado.")
	}

	passwordHashed, err := r.authDao.HashPassword(data.Senha)
	if err != nil {
		return nil, err
	}

	tx, err := r.dbPool.Begin(context.Background())
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(context.Background()); rollbackErr != nil {
				return
			}
		}
	}()

	sql := "INSERT INTO pessoas (nome, telefone, cpf) VALUES ($1, $2, $3) RETURNING id"
	rows, err := tx.Query(
		context.Background(),
		sql,
		data.Nome, data.Telefone, data.Cpf,
	)
	if err != nil {
		return nil, errors.New("Não foi possível criar usuário.")
	}

	var personId int
	for rows.Next() {
		if err := rows.Scan(&personId); err != nil {
			return nil, errors.New("Não foi possível encontrar pessoa criada.")
		}
	}

	sql = "INSERT INTO usuarios (email, senha, codigo_registro, pessoa_id) VALUES ($1, $2, $3, $4) RETURNING id"
	rows, err = tx.Query(
		context.Background(),
		sql,
		data.Email, passwordHashed, data.CodigoRegistro, personId,
	)
	if err != nil {
		return nil, errors.New("Não foi possível criar usuário.")
	}

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, errors.New("Não foi possível encontrar usuário criado.")
		}
	}
	defer rows.Close()

	if err := tx.Commit(context.Background()); err != nil {
		return nil, errors.New("Não foi possível criar usuário.")
	}

	return r.getUserById(id)
}

func (r userDao) getUserById(id int) (*User, error) {
	sql := `
		SELECT
			u.id AS id_usuario,
			u.email,
			u.senha,
			u.codigo_registro,
			u.permissoes,
			u.data_criacao,
			u.data_atualizacao,
			p.id AS id_pessoa,
			p.nome,
			p.telefone,
			p.cpf,
			p.data_criacao AS pessoa_data_criacao,
			p.data_atualizacao AS pessoa_data_atualizacao
		FROM usuarios AS u
		INNER JOIN pessoas AS p ON p.id = u.pessoa_id
		WHERE u.id = $1
	`

	var user *User = nil
	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
		id,
	)
	if err != nil {
		return nil, errors.New("Não foi possível encontrar usuário.")
	}
	defer rows.Close()

	for rows.Next() {
		user = &User{
			Pessoa: persons.Person{},
		}

		if err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
			&user.Permissoes,
			&user.DataCriacao,
			&user.DataAtualizacao,
			&user.Pessoa.Id,
			&user.Pessoa.Nome,
			&user.Pessoa.Telefone,
			&user.Pessoa.Cpf,
			&user.Pessoa.DataCriacao,
			&user.Pessoa.DataAtualizacao,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r userDao) getUser(code string, email string, phone string, cpf string) (*User, error) {
	sql := `
		SELECT
			u.id AS id_usuario,
			u.email,
			u.senha,
			u.codigo_registro,
			u.permissoes,
			u.data_criacao,
			u.data_atualizacao,
			p.id AS id_pessoa,
			p.nome,
			p.telefone,
			p.cpf,
			p.data_criacao AS pessoa_data_criacao,
			p.data_atualizacao AS pessoa_data_atualizacao
		FROM usuarios AS u
		INNER JOIN pessoas AS p ON p.id = u.pessoa_id
		WHERE u.codigo_registro = $1 OR u.email = $2 OR p.telefone = $3 OR p.cpf = $4
	`

	var user *User = nil
	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
		code, email, phone, cpf,
	)
	if err != nil {
		return nil, errors.New("Não foi possível buscar usuário por código de registro ou e-mail.")
	}
	defer rows.Close()

	for rows.Next() {
		user = &User{
			Pessoa: persons.Person{},
		}

		if err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
			&user.Permissoes,
			&user.DataCriacao,
			&user.DataAtualizacao,
			&user.Pessoa.Id,
			&user.Pessoa.Nome,
			&user.Pessoa.Telefone,
			&user.Pessoa.Cpf,
			&user.Pessoa.DataCriacao,
			&user.Pessoa.DataAtualizacao,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r userDao) getUsers() ([]User, error) {
	sql := `
		SELECT
			u.id AS id_usuario,
			u.email,
			u.senha,
			u.codigo_registro,
			u.permissoes,
			u.data_criacao,
			u.data_atualizacao,
			p.id AS id_pessoa,
			p.nome,
			p.telefone,
			p.cpf,
			p.data_criacao AS pessoa_data_criacao,
			p.data_atualizacao AS pessoa_data_atualizacao
		FROM usuarios AS u
		INNER JOIN pessoas AS p ON p.id = u.pessoa_id
	`

	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
	)
	if err != nil {
		return nil, errors.New("Não foi possível encontrar usuários.")
	}
	defer rows.Close()

	users := []User{}

	for rows.Next() {
		user := User{
			Pessoa: persons.Person{},
		}

		if err := rows.Scan(
			&user.Id,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
			&user.Permissoes,
			&user.DataCriacao,
			&user.DataAtualizacao,
			&user.Pessoa.Id,
			&user.Pessoa.Nome,
			&user.Pessoa.Telefone,
			&user.Pessoa.Cpf,
			&user.Pessoa.DataCriacao,
			&user.Pessoa.DataAtualizacao,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
