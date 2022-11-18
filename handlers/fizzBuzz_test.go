package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fizz-buzz-api/mocks"
	"fizz-buzz-api/models"
	"fizz-buzz-api/services"
	"fizz-buzz-api/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/icrowley/fake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/gin-gonic/gin"
)

type Suite struct {
	suite.Suite
	services services.Services
	w        *httptest.ResponseRecorder
	ctx      *gin.Context
}

// SetupTest setup suite
func (s *Suite) SetupTest() {
	s.w = httptest.NewRecorder()
	s.fakeContext()
}

// TeardownTest teardown suite
func (s *Suite) TeardownTest() {
	s.w.Flush()
}

func (s *Suite) fakeContext() {
	s.ctx, _ = gin.CreateTestContext(s.w)
	var err error
	s.ctx.Request, err = http.NewRequest("GET", fake.DomainName(), nil)
	s.Require().NoError(err)
}
func TestHTTPSuite(t *testing.T) {
	s := new(Suite)
	suite.Run(t, s)
}
func (s Suite) TestGetFizzBuzz_HappyPath() {
	fizzBuzzServiceMock := new(mocks.FizzBuzzService)
	statisticsServiceMock := new(mocks.StatisticsService)
	s.services = services.Services{
		FizzBuzzService:   fizzBuzzServiceMock,
		StatisticsService: statisticsServiceMock,
	}
	req := utils.FakeFizzBuzzRequest(2, 3, 10, "Fizz", "Buzz")
	b, err := json.Marshal(req)
	s.Require().NoError(err)
	s.ctx.Request, err = http.NewRequest("POST", fake.DomainName(), bytes.NewBuffer(b))
	s.Require().NoError(err)
	s.ctx.Request.Header.Add("content-type", "application/json")

	fizzBuzzServiceMock.On("ComputeFizzBuzz", req).Return([]string{"1", "Fizz", "Buzz", "Fizz", "5", "FizzBuzz", "7", "Fizz", "Buzz", "Fizz"}, nil)
	statisticsServiceMock.On("CountRequest", req).Return(nil)
	GetFizzBuzz(s.services)(s.ctx)
	s.Assert().Equal(http.StatusOK, s.w.Code)
}

func (s *Suite) TestGetFizzBuzz_Invalid() {
	invalidErrTest := func(tt *testing.T, name string, body interface{}, svc services.Services) {
		tt.Run(name, func(t *testing.T) {
			s.fakeContext()
			b, err := json.Marshal(body)
			s.Require().NoError(err)
			s.ctx.Request, err = http.NewRequest("POST", fake.DomainName(), bytes.NewBuffer(b))
			require.NoError(t, err)
			s.ctx.Request.Header.Add("content-type", "application/json")
			GetFizzBuzz(svc)(s.ctx)
			assert.Equal(t, http.StatusBadRequest, s.w.Code)
			mock.AssertExpectationsForObjects(t, svc.FizzBuzzService)
		})
	}
	fizzBuzzServiceMock := new(mocks.FizzBuzzService)
	statisticsServiceMock := new(mocks.StatisticsService)
	s.services = services.Services{
		FizzBuzzService:   fizzBuzzServiceMock,
		StatisticsService: statisticsServiceMock,
	}

	invalidErrTest(s.T(), "InvalidBody", "invalid", s.services)

	invalidErrTest(s.T(), "InvalidRequest", &models.FizzBuzzRequest{}, s.services)

}

func (s *Suite) TestGetFizzBuzz_Internal() {
	fizzBuzzServiceMock := new(mocks.FizzBuzzService)
	s.services = services.Services{
		FizzBuzzService: fizzBuzzServiceMock,
	}
	req := utils.FakeFizzBuzzRequest(-4, 3, 10, "Fizz", "Buzz")
	b, err := json.Marshal(req)
	s.Require().NoError(err)
	s.ctx.Request, err = http.NewRequest("POST", fake.DomainName(), bytes.NewBuffer(b))
	s.Require().NoError(err)
	s.ctx.Request.Header.Add("content-type", "application/json")

	fizzBuzzServiceMock.On("ComputeFizzBuzz", mock.Anything).Return(nil, errors.New("internal"))

	GetFizzBuzz(s.services)(s.ctx)
	s.Assert().Equal(http.StatusInternalServerError, s.w.Code)
	mock.AssertExpectationsForObjects(s.T(), fizzBuzzServiceMock)
}
