package tests

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/TimiBolu/lema-ai-users-service/dtos"
	"github.com/TimiBolu/lema-ai-users-service/handlers"
	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

// Handler Tests
func (suite *PostTestSuite) TestPostHandler_CreatePost() {
	reqBody := fmt.Sprintf(`{"userId": "%s", "title": "Hello, World!", "body": "This is a test post."}`, uuid1)
	req, _ := http.NewRequest("POST", "/api/posts", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	expectedPost := &models.Post{ID: post1, UserID: uuid1, Title: "Hello, World!", Body: "This is a test post."}
	suite.mockRepo.On("Create", mock.Anything, mock.MatchedBy(func(post *models.Post) bool {
		return post.UserID == expectedPost.UserID &&
			post.Title == expectedPost.Title &&
			post.Body == expectedPost.Body
	})).Return(nil)

	suite.postHandler.CreatePost(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	suite.Equal(http.StatusCreated, resp.StatusCode)
}

func (suite *PostTestSuite) TestPostHandler_GetPostsByUserID() {
	req := httptest.NewRequest("GET", "/api/posts?userId="+uuid1, nil) // Using query param
	w := httptest.NewRecorder()

	expectedPosts := []models.Post{
		{ID: post1, UserID: uuid1, Title: "Hello, World!", Body: "This is a test post."},
		{ID: post2, UserID: uuid1, Title: "Another post", Body: "This is another test post."},
	}

	suite.mockRepo.On("FindByUserID", mock.Anything, uuid1).Return(expectedPosts, nil)

	suite.postHandler.GetPostsByUser(w, req) // Call the handler
	resp := w.Result()
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var result handlers.SuccessResponse[dtos.GetPostsByUserResponseDTO]
	json.NewDecoder(resp.Body).Decode(&result)
	suite.Len(result.Data.Posts, 2) // Expecting two posts in the response
}

func (suite *PostTestSuite) TestPostHandler_DeletePost() {
	req, _ := http.NewRequest("DELETE", "/api/posts/"+post1, nil)
	w := httptest.NewRecorder()

	// Manually set the route variables
	req = mux.SetURLVars(req, map[string]string{"id": post1})

	suite.mockRepo.On("Delete", mock.Anything, post1).Return(nil)

	suite.postHandler.DeletePost(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	suite.Equal(http.StatusNoContent, resp.StatusCode)
}
