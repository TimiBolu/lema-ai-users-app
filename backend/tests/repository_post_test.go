package tests

import (
	"context"

	"github.com/TimiBolu/lema-ai-users-service/handlers"
	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/TimiBolu/lema-ai-users-service/services"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockPostRepository struct {
	mock.Mock
}

type PostTestSuite struct {
	suite.Suite
	mockRepo    *MockPostRepository
	postService services.PostService
	postHandler handlers.PostHandler
}

func (m *MockPostRepository) FindByUserID(ctx context.Context, userID string) ([]models.Post, error) {
	args := m.Called(ctx, userID)
	return args.Get(0).([]models.Post), args.Error(1)
}

func (m *MockPostRepository) Create(ctx context.Context, post *models.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *MockPostRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (suite *PostTestSuite) TestPostRepository_FindByUserID() {
	expectedPosts := []models.Post{
		{ID: post1, UserID: uuid1, Title: "Hello, World!", Body: "This is a test post."},
		{ID: post2, UserID: uuid1, Title: "Another post", Body: "This is a another test post."},
	}

	suite.mockRepo.On("FindByUserID", mock.Anything, uuid1).Return(expectedPosts, nil)

	posts, err := suite.mockRepo.FindByUserID(context.Background(), uuid1)
	suite.NoError(err)
	suite.Equal(expectedPosts, posts)
}

// Repository Tests
func (suite *PostTestSuite) TestPostRepository_Create() {
	expectedPost := models.Post{ID: post1, UserID: uuid1, Title: "Hello, World!", Body: "This is a test post."}
	suite.mockRepo.On("Create", mock.Anything, &expectedPost).Return(nil)

	err := suite.mockRepo.Create(context.Background(), &expectedPost)
	suite.NoError(err)
}

func (suite *PostTestSuite) TestPostRepository_Delete() {
	suite.mockRepo.On("Delete", mock.Anything, post1).Return(nil)

	err := suite.mockRepo.Delete(context.Background(), post1)
	suite.NoError(err)
}

func (suite *PostTestSuite) SetupTest() {
	logger := logrus.New()
	suite.mockRepo = new(MockPostRepository)
	suite.postService = services.NewPostService(suite.mockRepo)
	suite.postHandler = *handlers.NewPostHandler(suite.postService, logger)
}
