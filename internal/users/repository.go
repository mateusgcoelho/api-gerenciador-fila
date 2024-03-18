package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
)

type IUserRepository interface {
	createUser(data CreateUserDto) (*User, error)
	getUsers() ([]User, error)
	getUserById(id int) (*User, error)
	getUserByCodeOrEmail(code string, email string) (*User, error)
}

type userRepository struct {
	dbPool         *pgxpool.Pool
	authRepository auth.IAuthRepository
}

func NewUserRepository(dbPool *pgxpool.Pool, authRepository auth.IAuthRepository) IUserRepository {
	return userRepository{
		dbPool:         dbPool,
		authRepository: authRepository,
	}
}

func (r userRepository) createUser(data CreateUserDto) (*User, error) {
	user, err := r.getUserByCodeOrEmail(data.CodigoRegistro, data.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Usuário já cadastrado com codigo de registro e/ou e-mail.")
	}

	passwordHashed, err := r.authRepository.HashPassword(data.Senha)
	if err != nil {
		return nil, err
	}

	sql := "INSERT INTO usuarios (nome, email, senha, codigo_registro) VALUES ($1, $2, $3, $4) RETURNING id"

	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
		data.Nome, data.Email, passwordHashed, data.CodigoRegistro,
	)
	if err != nil {
		return nil, errors.New("Não foi possível criar usuário.")
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, errors.New("Não foi possível encontrar usuário criado.")
		}
	}

	return r.getUserById(id)
}

func (r userRepository) getUserById(id int) (*User, error) {
	sql := `
		SELECT id, nome, email, senha, codigo_registro, permissoes, data_criacao, data_atualizacao
		FROM usuarios AS u WHERE u.id = $1
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
		user = &User{}
		if err := rows.Scan(
			&user.Id,
			&user.Nome,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
			&user.Permissoes,
			&user.DataCriacao,
			&user.DataAtualizacao,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r userRepository) getUserByCodeOrEmail(code string, email string) (*User, error) {
	sql := `
		SELECT id, nome, email, senha, codigo_registro, permissoes, data_criacao, data_atualizacao
		FROM usuarios AS u
		WHERE u.codigo_registro = $1 OR u.email = $2
	`

	var user *User = nil
	rows, err := r.dbPool.Query(
		context.Background(),
		sql,
		code, email,
	)
	if err != nil {
		return nil, errors.New("Não foi possível buscar usuário por código de registro ou e-mail.")
	}
	defer rows.Close()

	for rows.Next() {
		user = &User{}
		if err := rows.Scan(
			&user.Id,
			&user.Nome,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
			&user.Permissoes,
			&user.DataCriacao,
			&user.DataAtualizacao,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r userRepository) getUsers() ([]User, error) {
	sql := `
		SELECT id, nome, email, senha, codigo_registro, permissoes, data_criacao, data_atualizacao
		FROM usuarios
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
		user := User{}
		if err := rows.Scan(
			&user.Id,
			&user.Nome,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
			&user.Permissoes,
			&user.DataCriacao,
			&user.DataAtualizacao,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
