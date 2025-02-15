package tests

import (
	"context"
	"errors"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/stretchr/testify/mock"
)

func (suite *UserTestSuite) TestUserService_GetUserByID() {
	expectedUser := &models.User{ID: uuid1, Name: "John Doe", Username: "johndoe"}
	suite.mockRepo.On("FindByID", mock.Anything, uuid1).Return(expectedUser, nil)
	suite.mockRepo.On("FindByID", mock.Anything, uuid2).Return(nil, errors.New("user not found"))

	user, err := suite.userService.GetUserByID(context.Background(), uuid1)
	suite.NoError(err)
	suite.Equal(expectedUser, user)

	nilUser, err := suite.userService.GetUserByID(context.Background(), uuid2)
	suite.Error(err)
	suite.Nil(nilUser)
}

func (suite *UserTestSuite) TestUserService_GetUsers() {
	expectedUsers := []models.User{
		{ID: uuid1, Name: "John Doe", Username: "johndoe"},
		{ID: uuid2, Name: "Jane Doe", Username: "janedoe"},
	}
	totalCount := int64(2)

	suite.mockRepo.On("FindAll", mock.Anything, 0, 10).Return(expectedUsers, totalCount, nil)

	users, count, err := suite.userService.GetUsers(context.Background(), 0, 10)
	suite.NoError(err)
	suite.Equal(expectedUsers, users)
	suite.Equal(totalCount, count)
}

func (suite *UserTestSuite) TestUserService_GetUserCount() {
	expectedCount := int64(2)

	suite.mockRepo.On("Count", mock.Anything).Return(expectedCount, nil)

	count, err := suite.userService.GetUserCount(context.Background())
	suite.NoError(err)
	suite.Equal(expectedCount, count)
}
