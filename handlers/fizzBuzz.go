package handlers

import (
	"net/http"

	"fizz-buzz-api/models"
	"fizz-buzz-api/services"
	"fizz-buzz-api/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetFizzBuzz(s services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(models.FizzBuzzRequest)
		err := ctx.Bind(req)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}

		if err := req.Validate(); err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}
		response, err := s.FizzBuzzService.ComputeFizzBuzz(req)

		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		err = s.StatisticsService.CountRequest(req)
		if err != nil {
			utils.Logger.Error("Error when counting request stats", zap.Error(err))
		}
		ctx.JSON(http.StatusOK, response)
	}
}
