package tests

import (
	"context"
	"errors"

	"github.com/TimiBolu/lema-ai-users-service/handlers"
	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/TimiBolu/lema-ai-users-service/services"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserRepository struct {
	mock.Mock
}

type UserTestSuite struct {
	suite.Suite
	mockRepo    *MockUserRepository
	userService services.UserService
	userHandler handlers.UserHandler
}

func (m *MockUserRepository) FindAll(ctx context.Context, page, limit int) ([]models.User, int64, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]models.User), args.Get(1).(int64), args.Error(2)
}

func (m *MockUserRepository) FindByID(ctx context.Context, id string) (*models.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepository) Count(ctx context.Context) (int64, error) {
	args := m.Called(ctx)
	return args.Get(0).(int64), args.Error(1)
}

func (suite *UserTestSuite) TestUserRepository_FindByID() {
	expectedUser := &models.User{ID: uuid1, FirstName: "John", LastName: "Doe"}
	suite.mockRepo.On("FindByID", mock.Anything, uuid1).Return(expectedUser, nil)
	suite.mockRepo.On("FindByID", mock.Anything, uuid2).Return(nil, errors.New("user not found"))

	user, err := suite.mockRepo.FindByID(context.Background(), uuid1)
	suite.NoError(err)
	suite.Equal(expectedUser, user)

	nilUser, err := suite.mockRepo.FindByID(context.Background(), uuid2)
	suite.Error(err)
	suite.Nil(nilUser)
}

func (suite *UserTestSuite) TestUserRepository_FindAll() {
	expectedUsers := []models.User{
		{ID: uuid1, FirstName: "John", LastName: "Doe"},
		{ID: uuid2, FirstName: "Jane", LastName: "Doe"},
	}
	totalCount := int64(2)

	suite.mockRepo.On("FindAll", mock.Anything, 0, 10).Return(expectedUsers, totalCount, nil)

	users, count, err := suite.mockRepo.FindAll(context.Background(), 0, 10)
	suite.NoError(err)
	suite.Equal(expectedUsers, users)
	suite.Equal(totalCount, count)
}

func (suite *UserTestSuite) TestUserRepository_Count() {
	expectedCount := int64(2)

	suite.mockRepo.On("Count", mock.Anything).Return(expectedCount, nil)

	count, err := suite.mockRepo.Count(context.Background())
	suite.NoError(err)
	suite.Equal(expectedCount, count)
}

func (suite *UserTestSuite) SetupTest() {
	logger := logrus.New()
	suite.mockRepo = new(MockUserRepository)
	suite.userService = services.NewUserService(suite.mockRepo)
	suite.userHandler = *handlers.NewUserHandler(suite.userService, logger)
}
