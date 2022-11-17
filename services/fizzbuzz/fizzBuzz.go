//go:generate mockgen -destination=../../test/mocks/services/fizzbuzz/fizzbuzz.go -package=fizzbuzz -source=./fizzbuzz.go

package fizzbuzz

import (
	"fmt"
)

type Service interface {
	Compute(request Command) ([]string, error)
}

type FizzBuzzService struct {
}

func InitService() *FizzBuzzService {
	return &FizzBuzzService{}
}

// Compute computes the fizz buzz repsonse for a given command (int1, int3, limit, str1, str2)
func (service FizzBuzzService) Compute(command Command) (result []string, err error) {
	for i := 1; i <= command.Limit; i++ {
		if i%(command.Int1*command.Int2) == 0 {
			result = append(result, fmt.Sprintf("%s%s", command.Str1, command.Str2))
		} else if i%command.Int1 == 0 {
			result = append(result, command.Str1)
		} else if i%command.Int2 == 0 {
			result = append(result, command.Str2)
		} else {
			result = append(result, fmt.Sprint(i))
		}
	}
	return
}
