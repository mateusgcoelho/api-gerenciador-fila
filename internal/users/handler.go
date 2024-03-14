package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/auth"
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

func handleLogin(userRepository IUserRepository, authRepository auth.IAuthRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data LoginDto
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, utils.DefaultResponse{
				Message: "Verifique o corpo da requisição.",
				Data:    nil,
			})
			return
		}

		user, err := userRepository.getUserByCodeOrEmail("", data.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		if user == nil {
			c.JSON(http.StatusBadRequest, utils.DefaultResponse{
				Message: "E-mail e/ou senha inválidos.",
				Data:    nil,
			})
			return
		}

		isValid := authRepository.ComparePasswords(user.Senha, data.Senha)

		if !isValid {
			c.JSON(http.StatusBadRequest, utils.DefaultResponse{
				Message: "E-mail e/ou senha inválidos.",
				Data:    nil,
			})
			return
		}

		tokenJwt, err := authRepository.GenerateJwtToken(auth.JwtPayloadDto{
			Id:         user.Id,
			Permissoes: user.Permissoes,
		})

		c.JSON(http.StatusOK, utils.DefaultResponse{
			Data: tokenJwt,
		})
	}
}
