package queues

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func handleCreateQueue(queueRepository IQueueRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data CreateQueueDto
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, utils.DefaultResponse{
				Message: "Verifique o corpo da requisição.",
				Data:    nil,
			})
			return
		}

		queue, err := queueRepository.createQueue(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, utils.DefaultResponse{
			Data: queue,
		})
	}
}

func handleGetQueues(queueRepository IQueueRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		queues, err := queueRepository.getQueues()
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, utils.DefaultResponse{
			Data: queues,
		})
	}
}