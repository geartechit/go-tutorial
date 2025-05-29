package handlers_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go-tutorial/internal/config"
	"go-tutorial/internal/domain/employee"
	"go-tutorial/internal/dto"
	"go-tutorial/internal/handlers"
	"go-tutorial/internal/mocks"
	"go-tutorial/pkg/logger"
	"go-tutorial/pkg/mapper"
	"go-tutorial/pkg/timeutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type EmployeeHandlerTestSuite struct {
	suite.Suite
	ctrl          *gomock.Controller
	handler       *handlers.EmployeeHandler
	mockSvc       *mocks.MockEmployeeService
	mockRepo      *mocks.MockEmployeeRepository
	mockValidator *mocks.MockDTOValidator
	logger        logger.Logger
	recorder      *httptest.ResponseRecorder

	existingEmployee      domain.Employee
	createEmployeeRequest *dto.CreateEmployeeRequest
	updateEmployeeRequest *dto.UpdateEmployeeRequest
}

func TestEmployeeHandlerSuite(t *testing.T) {
	suite.Run(t, new(EmployeeHandlerTestSuite))
}

func (s *EmployeeHandlerTestSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *EmployeeHandlerTestSuite) SetupTest() {
	cfg := config.LoadConfig()
	s.ctrl = gomock.NewController(s.T())
	s.mockSvc = mocks.NewMockEmployeeService(s.ctrl)
	s.mockRepo = mocks.NewMockEmployeeRepository(s.ctrl)
	s.mockValidator = mocks.NewMockDTOValidator(s.ctrl)
	s.logger = logger.NewZapLogger(cfg)

	s.recorder = httptest.NewRecorder()
	s.handler = handlers.NewEmployeeHandler(s.mockSvc, s.logger, s.mockValidator)

	s.existingEmployee = domain.Employee{
		ID:         uuid.New(),
		Name:       "Foo",
		DOB:        timeutil.MustParseTime("2000-01-01T00:00:00Z"),
		Department: "Engineering",
		JobTitle:   "Staff",
		Address:    "140 wireless building",
		JoinedAt:   timeutil.MustParseTime("2000-01-01T00:00:00Z"),
		CreatedAt:  timeutil.MustParseTime("2000-01-01T00:00:00Z"),
		UpdatedAt:  timeutil.MustParseTime("2000-01-01T00:00:00Z"),
	}
	s.createEmployeeRequest = &dto.CreateEmployeeRequest{
		Name:       "Foo",
		Dob:        timeutil.MustParseTime("2000-01-01T00:00:00Z"),
		Department: "Engineering",
		JobTitle:   "Developer",
		Address:    "140 wireless building",
		JoinedAt:   timeutil.MustParseTime("2023-01-01T00:00:00Z"),
	}
}

func (s *EmployeeHandlerTestSuite) TestCreateEmployee_Success() {
	reqBody := s.createEmployeeRequest
	body, err := json.Marshal(reqBody)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	expectedModel := mapper.ToCreateEmployeeModel(reqBody)
	expectedResult := &domain.Employee{
		ID:         uuid.New(),
		Name:       reqBody.Name,
		DOB:        expectedModel.DOB,
		Department: reqBody.Department,
		JobTitle:   reqBody.JobTitle,
		Address:    reqBody.Address,
		JoinedAt:   expectedModel.JoinedAt,
	}

	s.mockValidator.EXPECT().
		Validate(reqBody).
		Return(nil)

	s.mockSvc.EXPECT().
		Create(gomock.Any(), reqBody).
		Return(expectedResult, nil)

	s.handler.CreateEmployee(s.recorder, req)

	s.Equal(http.StatusCreated, s.recorder.Code)

	var res dto.BaseResponse[*dto.EmployeeResponse]
	err = json.NewDecoder(s.recorder.Body).Decode(&res)
	s.Require().NoError(err)

	parsedDob, err := time.Parse(time.RFC3339, res.Data.DOB)
	s.Require().NoError(err)
	s.Equal(reqBody.Dob, parsedDob)

	parsedJoinedAt, err := time.Parse(time.RFC3339, res.Data.JoinedAt)
	s.Require().NoError(err)
	s.Equal(reqBody.JoinedAt, parsedJoinedAt)

	s.Equal(reqBody.Name, res.Data.Name)
	s.Equal(reqBody.Department, res.Data.Department)
	s.Equal(reqBody.Address, res.Data.Address)
	s.Equal(reqBody.JobTitle, res.Data.JobTitle)
}

func (s *EmployeeHandlerTestSuite) TestCreateEmployee_ValidationError() {
	reqBody := s.createEmployeeRequest
	reqBody.Name = ""
	body, err := json.Marshal(reqBody)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	s.mockValidator.EXPECT().
		Validate(gomock.Any()).
		Return([]string{"Name must be at least 2 characters in length"})

	s.handler.CreateEmployee(s.recorder, req)

	s.Equal(http.StatusBadRequest, s.recorder.Code)

	var res dto.ErrorResponse
	err = json.NewDecoder(s.recorder.Body).Decode(&res)
	s.NoError(err)
	s.Equal("invalid request body", res.Message)
	s.Contains(res.Details, "Name must be at least 2 characters in length")
}

func (s *EmployeeHandlerTestSuite) TestCreateEmployee_InternalServerError() {
	reqBody := s.createEmployeeRequest
	body, err := json.Marshal(reqBody)
	s.Require().NoError(err)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/employees", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	s.mockValidator.EXPECT().
		Validate(reqBody).
		Return(nil)

	s.mockSvc.EXPECT().
		Create(context.Background(), reqBody).
		Return(nil, errors.New("mock service error"))

	s.handler.CreateEmployee(s.recorder, req)

	s.Equal(http.StatusInternalServerError, s.recorder.Code)

	var res dto.ErrorResponse
	err = json.NewDecoder(s.recorder.Body).Decode(&res)
	s.NoError(err)
	s.Equal("error creating employee", res.Message)
}

func (s *EmployeeHandlerTestSuite) TestGetEmployeeByID_Success() {
	employeeID := uuid.New()
	expectedEmployee := &domain.Employee{
		ID:         employeeID,
		Name:       "Foo Bar Baz",
		DOB:        timeutil.MustParseTime("2000-01-01T00:00:00Z"),
		Department: "Engineering",
		JobTitle:   "Developer",
		Address:    "140 Wireless Building",
		JoinedAt:   timeutil.MustParseTime("2023-01-01T00:00:00Z"),
	}

	s.mockSvc.EXPECT().
		GetByID(gomock.Any(), employeeID).
		Return(expectedEmployee, nil)

	router := chi.NewRouter()
	router.Get("/api/v1/employees/{id}", s.handler.GetEmployeeByID)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/employees/"+employeeID.String(), nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	s.Equal(http.StatusOK, s.recorder.Code)

	var res dto.BaseResponse[dto.EmployeeResponse]
	err := json.NewDecoder(rec.Body).Decode(&res)
	s.Require().NoError(err)
	s.Equal("Foo Bar Baz", res.Data.Name)
}
