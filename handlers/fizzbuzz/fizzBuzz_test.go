package fizzbuzz

import (
	"bytes"
	"encoding/json"
	"errors"
	"fizz-buzz-api/services"
	"fizz-buzz-api/services/fizzbuzz"
	fizzbuzzMock "fizz-buzz-api/test/mocks/services/fizzbuzz"
	statisticsMock "fizz-buzz-api/test/mocks/services/statistics"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockFizzBuzzService struct {
	mock.Mock
}

func (m *MockFizzBuzzService) Compute(command fizzbuzz.Command) ([]string, error) {
	args := m.Called(command)
	return args.Get(0).([]string), args.Error(1)
}

type MockStatisticsService struct {
	mock.Mock
}

func (m *MockStatisticsService) CountRequest(command fizzbuzz.Command) error {
	args := m.Called(command)
	return args.Error(0)
}

func TestGetFizzBuzz(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name              string
		requestBody       string
		validRequest      bool
		mockComputeResult []string
		mockComputeError  error
		mockCountError    error
		expectedStatus    int
		expectedResponse  gin.H
	}{
		{
			name:             "Invalid Request Binding",
			requestBody:      `{invalid_json}`, // Invalid JSON format
			validRequest:     false,
			expectedStatus:   http.StatusBadRequest,
			expectedResponse: gin.H{"error": "invalid character 'i' looking for beginning of object key string"},
		},
		{
			name:           "Validation Errors",
			requestBody:    `{"int1": 0, "int2": 0, "limit": -1, "str1": "", "str2": ""}`,
			validRequest:   false,
			expectedStatus: http.StatusBadRequest,
			expectedResponse: gin.H{"errors": []interface{}{
				"int1 is required and must be greater than zero",
				"int2 is required and must be greater than zero",
				"limit must be greater than zero",
				"str1 is required and must not be empty",
				"str2 is required and must not be empty",
			}},
		},
		{
			name:              "Successful FizzBuzz Computation",
			requestBody:       `{"int1": 3, "int2": 5, "limit": 15, "str1": "Fizz", "str2": "Buzz"}`,
			validRequest:      true,
			mockComputeResult: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			expectedStatus:    http.StatusOK,
			expectedResponse:  gin.H{"response": []interface{}{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
		},
		{
			name:             "Compute Service Error",
			requestBody:      `{"int1": 3, "int2": 5, "limit": 15, "str1": "Fizz", "str2": "Buzz"}`,
			validRequest:     true,
			mockComputeError: errors.New("internal error"),
			expectedStatus:   http.StatusInternalServerError,
			expectedResponse: gin.H{"error": "internal error"},
		},
		{
			name:              "Statistics Service Error",
			requestBody:       `{"int1": 3, "int2": 5, "limit": 15, "str1": "Fizz", "str2": "Buzz"}`,
			validRequest:      true,
			mockComputeResult: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
			mockCountError:    errors.New("count error"),
			expectedStatus:    http.StatusOK, // Should still return OK even if counting fails
			expectedResponse:  gin.H{"response": []interface{}{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			// Create mocks
			fizzbuzzMockService := fizzbuzzMock.NewMockService(ctrl)
			statisticsMockService := statisticsMock.NewMockService(ctrl)

			mockServices := services.Services{
				FizzBuzzService:   fizzbuzzMockService,
				StatisticsService: statisticsMockService,
			}

			// Set up mock behaviors
			if tt.validRequest {
				mockCommand := fizzbuzz.Command{Int1: 3, Int2: 5, Limit: 15, Str1: "Fizz", Str2: "Buzz"}

				if tt.mockComputeResult != nil || tt.mockComputeError != nil {
					fizzbuzzMockService.EXPECT().Compute(mockCommand).Return(tt.mockComputeResult, tt.mockComputeError)
				}
				if tt.mockComputeError == nil {
					if tt.mockCountError != nil {
						statisticsMockService.EXPECT().CountRequest(mockCommand).Return(tt.mockCountError)
					} else {
						statisticsMockService.EXPECT().CountRequest(mockCommand).Return(nil)
					}
				}
			}

			// Create the HTTP request and recorder
			req, err := http.NewRequest(http.MethodPost, "/fizzbuzz", bytes.NewBufferString(tt.requestBody))
			require.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)

			// Assign the request to the context
			ctx.Request = req

			// Execute the handler
			handler := GetFizzBuzz(mockServices)
			handler(ctx)

			// Verify status code and response body
			require.Equal(t, tt.expectedStatus, w.Code)
			var response gin.H
			err = json.Unmarshal(w.Body.Bytes(), &response)
			require.NoError(t, err)
			require.Equal(t, tt.expectedResponse, response)
		})
	}
}
