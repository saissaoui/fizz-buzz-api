package services

import (
	"fizz-buzz-api/utils"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/pkg/errors"

	"testing"
)

func TestComputeFizzBuzz_HappyPath(t *testing.T) {
	service := new(FizzBuzzService)
	request := utils.FakeFizzBuzzRequest(2, 3, 10, "Fizz", "Buzz")

	response, err := service.ComputeFizzBuzz(request)

	result := []string{"1", "Fizz", "Buzz", "Fizz", "5", "FizzBuzz", "7", "Fizz", "Buzz", "Fizz"}
	require.NoError(t, err)
	assert.Equal(t, result, response)
}

func TestComputeFizzBuzz_InvalidLimit(t *testing.T) {
	e := errors.New("Limit not valid, must be greater than zero")
	service := new(FizzBuzzService)
	request := utils.FakeFizzBuzzRequest(2, 3, 0, "Fizz", "Buzz")

	_, err := service.ComputeFizzBuzz(request)

	require.Error(t, err)
	assert.Equal(t, e.Error(), err.Error())
}

func TestComputeFizzBuzz_InvalidInt1(t *testing.T) {
	e := errors.New("Int1 not valid, must be greater than zero")

	service := new(FizzBuzzService)
	request := utils.FakeFizzBuzzRequest(-1, 3, 110, "Fizz", "Buzz")

	_, err := service.ComputeFizzBuzz(request)

	require.Error(t, err)
	assert.Equal(t, e.Error(), err.Error())
}

func TestComputeFizzBuzz_InvalidInt2(t *testing.T) {
	e := errors.New("Int2 not valid, must be greater than zero")

	service := new(FizzBuzzService)
	request := utils.FakeFizzBuzzRequest(2, -3, 10, "Fizz", "Buzz")

	_, err := service.ComputeFizzBuzz(request)

	require.Error(t, err)
	assert.Equal(t, e.Error(), err.Error())

}

func TestComputeFizzBuzz_InvalidStr1(t *testing.T) {
	e := errors.New("Str1 not valid, must not be empty")

	service := new(FizzBuzzService)
	request := utils.FakeFizzBuzzRequest(2, 3, 10, "", "Buzz")

	_, err := service.ComputeFizzBuzz(request)

	require.Error(t, err)
	assert.Equal(t, e.Error(), err.Error())

}

func TestComputeFizzBuzz_InvalidStr2(t *testing.T) {
	e := errors.New("Str2 not valid, must not be empty")

	service := new(FizzBuzzService)
	request := utils.FakeFizzBuzzRequest(2, 3, 10, "Fizz", "")

	_, err := service.ComputeFizzBuzz(request)

	require.Error(t, err)
	assert.Equal(t, e.Error(), err.Error())
}
