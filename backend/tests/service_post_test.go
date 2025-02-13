package tests

import (
	"context"

	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/stretchr/testify/mock"
)

// Service Tests
func (suite *PostTestSuite) TestPostService_CreatePost() {
	expectedPost := models.Post{ID: post1, UserID: uuid1, Title: "Hello, World!", Body: "This is a test post."}
	suite.mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(post *models.Post) bool {
		return post.UserID == expectedPost.UserID &&
			post.Title == expectedPost.Title &&
			post.Body == expectedPost.Body
	})).Return(nil)

	_, err := suite.postService.CreatePost(context.Background(), expectedPost.Title, expectedPost.Body, expectedPost.UserID)
	suite.NoError(err)
}

func (suite *PostTestSuite) TestPostService_GetPostsByUserID() {
	expectedPosts := []models.Post{
		{ID: post1, UserID: uuid1, Title: "Hello, World!", Body: "This is a test post."},
		{ID: post2, UserID: uuid1, Title: "Another post", Body: "This is another test post."},
	}
	suite.mockRepo.On("FindByUserID", mock.Anything, uuid1).Return(expectedPosts, nil)

	posts, err := suite.postService.GetPostsByUser(context.Background(), uuid1)
	suite.NoError(err)
	suite.Equal(expectedPosts, posts)
}

func (suite *PostTestSuite) TestPostService_DeletePost() {
	suite.mockRepo.On("Delete", mock.Anything, post1).Return(nil)

	err := suite.postService.DeletePost(context.Background(), post1)
	suite.NoError(err)
}
