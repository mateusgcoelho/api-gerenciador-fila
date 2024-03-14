package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleCreateUser(userRepository IUserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data CreateUserDto
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userRepository.createUser(data)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	}
}
