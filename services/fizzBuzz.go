package services

import (
	"fizz-buzz-api/models"
	"fmt"

	"github.com/pkg/errors"
)

//go:generate mockery --name=FizzBuzzService --output ./../mocks --case=underscore

type FizzBuzzService interface {
	ComputeFizzBuzz(request *models.FizzBuzzRequest) ([]string, error)
}

type FizzBuzzServiceImpl struct {
}

//ComputeFizzBuzz computes the fizz buzz repsonse for a given request (int1, int3, limit, str1, str2)
func (service FizzBuzzServiceImpl) ComputeFizzBuzz(request *models.FizzBuzzRequest) (result []string, err error) {
	if *request.Limit <= 0 {
		err = errors.New("Limit not valid, must be greater than zero")
		return
	}
	if *request.Int1 <= 0 {
		err = errors.New("Int1 not valid, must be greater than zero")
		return
	}
	if *request.Int2 <= 0 {
		err = errors.New("Int2 not valid, must be greater than zero")
		return
	}
	if len(*request.Str1) == 0 {
		err = errors.New("Str1 not valid, must not be empty")
		return
	}
	if len(*request.Str2) == 0 {
		err = errors.New("Str2 not valid, must not be empty")
		return
	}

	for i := 1; i <= *request.Limit; i++ {
		if i%(*request.Int1**request.Int2) == 0 {
			result = append(result, fmt.Sprintf("%s%s", *request.Str1, *request.Str2))
		} else if i%*request.Int1 == 0 {
			result = append(result, *request.Str1)
		} else if i%*request.Int2 == 0 {
			result = append(result, *request.Str2)
		} else {
			result = append(result, fmt.Sprint(i))
		}
	}
	return
}
