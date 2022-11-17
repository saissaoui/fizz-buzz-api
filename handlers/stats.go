package handlers

import (
	"fizz-buzz-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetStatistics(s services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {

		resp, err := s.StatisticsService.GetStatistics()

		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, resp)
	}
}
