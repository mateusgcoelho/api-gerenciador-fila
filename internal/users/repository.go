package users

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type IUserRepository interface {
	createUser(data CreateUserDto) (*User, error)
	getUserById(id int) (*User, error)
	getUserByCodeOrEmail(code string, email string) (*User, error)
}

type userRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) IUserRepository {
	return userRepositoryImpl{
		db: db,
	}
}

func (r userRepositoryImpl) createUser(data CreateUserDto) (*User, error) {
	user, err := r.getUserByCodeOrEmail(data.CodigoRegistro, data.Email)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return nil, errors.New("Usuário já cadastrado com codigo de registro e/ou e-mail.")
	}

	sql := "INSERT INTO usuarios (nome, email, senha, codigo_registro) VALUES ($1, $2, $3, $4) RETURNING id"

	var id int
	rows, err := r.db.Query(
		context.Background(),
		sql,
		data.Nome, data.Email, data.Senha, data.CodigoRegistro,
	)
	if err != nil {
		return nil, errors.New("Não foi possível criar usuário.")
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return nil, errors.New("Não foi possível encontrar usuário criado.")
		}
	}

	return r.getUserById(id)
}

func (r userRepositoryImpl) getUserById(id int) (*User, error) {
	sql := "SELECT nome, email, senha, codigo_registro FROM usuarios AS u WHERE u.id = $1"

	var user *User = nil
	rows, err := r.db.Query(
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
			&user.Nome,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}

func (r userRepositoryImpl) getUserByCodeOrEmail(code string, email string) (*User, error) {
	sql := "SELECT nome, email, senha, codigo_registro FROM usuarios AS u WHERE u.codigo_registro = $1 OR u.email = $2"

	var user *User = nil
	rows, err := r.db.Query(
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
			&user.Nome,
			&user.Email,
			&user.Senha,
			&user.CodigoRegistro,
		); err != nil {
			return nil, err
		}
	}

	return user, nil
}
