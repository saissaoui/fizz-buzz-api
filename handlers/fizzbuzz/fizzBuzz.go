package fizzbuzz

import (
	"net/http"

	"fizz-buzz-api/services"
	"fizz-buzz-api/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetFizzBuzz(s services.Services) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := new(Request)
		err := ctx.Bind(req)
		if err != nil {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"error": err.Error()})
			return
		}

		if errors := req.Validate(); len(errors) > 0 {
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{"errors": errors})
			return
		}
		command := req.ToCommand()
		response, err := s.FizzBuzzService.Compute(command)

		if err != nil {
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{"error": err.Error()})
			return
		}
		err = s.StatisticsService.CountRequest(command)
		if err != nil {
			utils.Logger.Error("Error when counting request stats", zap.Error(err))
		}
		ctx.JSON(http.StatusOK, gin.H{"response": response})
	}
}
