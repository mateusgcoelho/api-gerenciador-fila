package reports

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusgcoelho/api-gerenciador-fila/internal/utils"
)

func handleCreateReport(reportDao IReportDao) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data CreateReportDto
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, utils.DefaultResponse{
				Message: "Verifique o corpo da requisição.",
				Data:    nil,
			})
			return
		}

		report, err := reportDao.createReport(data)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusCreated, utils.DefaultResponse{
			Message: "Atendimento iniciado com sucesso.",
			Data:    report,
		})
	}
}

func handleGetReports(reportDao IReportDao) gin.HandlerFunc {
	return func(c *gin.Context) {
		reports, err := reportDao.getReports()
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.DefaultResponse{
				Message: err.Error(),
				Data:    nil,
			})
			return
		}

		c.JSON(http.StatusOK, utils.DefaultResponse{
			Data: reports,
		})
	}
}
