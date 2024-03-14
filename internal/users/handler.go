package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func handleCreateUser(userRepository IUserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data CreateUserDto
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, utils.DefaultResponse{
				Message: "Verifique o corpo da requisição.",
				Data:    nil,
			})
			return
		}

		user, err := userRepository.createUser(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, utils.DefaultResponse{
			Message: "Usuário cadastrado com sucesso.",
			Data:    user,
		})
	}
}

func handleGetUsers(userRepository IUserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := userRepository.getUsers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, utils.DefaultResponse{
			Data: users,
		})
	}
}
