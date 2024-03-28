package users

import (
	"time"

	database "github.com/mateusgcoelho/api-gerenciador-fila/database/sqlc"
)

type UserResponse struct {
	ID        int32     `json:"id"`
	Email     string    `json:"email"`
	Code      string    `json:"code"`
	PersonID  int32     `json:"person_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromUserEntity(u database.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Email:     u.Email,
		Code:      u.Code.String,
		PersonID:  u.PersonID,
		CreatedAt: u.CreatedAt.Time,
		UpdatedAt: u.UpdatedAt.Time,
	}
}

func FromListUserEntity(users []database.User) []UserResponse {
	var usersResponse []UserResponse = []UserResponse{}
	for _, u := range users {
		usersResponse = append(usersResponse, FromUserEntity(u))
	}
	return usersResponse
}
