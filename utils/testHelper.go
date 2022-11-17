package utils

import "fizz-buzz-api/models"

func FakeFizzBuzzRequest(int1, int2, limit int, str1, str2 string) *models.FizzBuzzRequest {
	return &models.FizzBuzzRequest{
		Int1:  &int1,
		Int2:  &int2,
		Str1:  &str1,
		Str2:  &str2,
		Limit: &limit,
	}
}
