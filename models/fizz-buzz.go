package models

import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type FizzBuzzRequest struct {
	Int1  *int    `json:"int1" valid:"required"`
	Int2  *int    `json:"int2" valid:"required"`
	Limit *int    `json:"limit" valid:"required"`
	Str1  *string `json:"str1" valid:"required"`
	Str2  *string `json:"str2" valid:"required"`
}

// Validate validates a FizzBuzzRequest
func (req *FizzBuzzRequest) Validate() error {
	if _, err := govalidator.ValidateStruct(req); err != nil {
		return err
	}
	return nil
}

func (req FizzBuzzRequest) ToString() string {
	template := `{int1: %d, int2: %d, limit: %d, str1: %s, str2 : %s }`
	return fmt.Sprintf(template, *req.Int1, *req.Int2, *req.Limit, *req.Str1, *req.Str2)
}
