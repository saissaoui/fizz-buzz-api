package fizzbuzz

import (
	"reflect"
	"testing"
)

func TestFizzBuzzService_Compute(t *testing.T) {
	service := InitService()
	tests := []struct {
		name     string
		command  Command
		expected []string
	}{
		{
			name: "Basic FizzBuzz",
			command: Command{
				Int1:  3,
				Int2:  5,
				Limit: 15,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: []string{
				"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz",
				"11", "Fizz", "13", "14", "FizzBuzz",
			},
		},
		{
			name: "Only multiples of Int1",
			command: Command{
				Int1:  2,
				Int2:  10,
				Limit: 10,
				Str1:  "Two",
				Str2:  "Ten",
			},
			expected: []string{
				"1", "Two", "3", "Two", "5", "Two", "7", "Two", "9", "Two",
			},
		},
		{
			name: "Only multiples of Int2",
			command: Command{
				Int1:  7,
				Int2:  3,
				Limit: 5,
				Str1:  "Seven",
				Str2:  "Three",
			},
			expected: []string{
				"1", "2", "Three", "4", "5",
			},
		},
		{
			name: "Limit is zero",
			command: Command{
				Int1:  3,
				Int2:  5,
				Limit: 0,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: nil,
		},
		{
			name: "Negative limit",
			command: Command{
				Int1:  3,
				Int2:  5,
				Limit: -5,
				Str1:  "Fizz",
				Str2:  "Buzz",
			},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.Compute(tt.command)
			if err != nil {
				t.Fatalf("Compute() returned an error: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Compute() = %v, expected %v", result, tt.expected)
			}
		})
	}
}
