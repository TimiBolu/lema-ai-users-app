package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/TimiBolu/lema-ai-users-service/handlers"
	"github.com/TimiBolu/lema-ai-users-service/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

func (suite *UserTestSuite) TestUserHandler_GetUsers() {
	req, _ := http.NewRequest("GET", "/api/users?pageNumber=1&pageSize=10", nil)
	w := httptest.NewRecorder()

	expectedUsers := []models.User{
		{ID: uuid1, Name: "John Doe", Username: "johndoe"},
		{ID: uuid2, Name: "Jane Doe", Username: "janedoe"},
	}
	totalCount := int64(2)

	suite.mockRepo.On("FindAll", mock.Anything, 1, 10).Return(expectedUsers, totalCount, nil)

	suite.userHandler.GetUsers(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var response handlers.SuccessResponse[map[string]interface{}]
	json.NewDecoder(resp.Body).Decode(&response)
	suite.Equal(2, int(response.Data["pagination"].(map[string]interface{})["totalItems"].(float64)))
}

func (suite *UserTestSuite) TestUserHandler_GetUserCount() {
	req, _ := http.NewRequest("GET", "/api/users/count", nil)
	w := httptest.NewRecorder()

	expectedCount := int64(2)

	suite.mockRepo.On("Count", mock.Anything).Return(expectedCount, nil)

	suite.userHandler.GetUsersCount(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var res handlers.SuccessResponse[map[string]int64]
	json.NewDecoder(resp.Body).Decode(&res)
	suite.Equal(expectedCount, res.Data["count"])
}

func (suite *UserTestSuite) TestUserHandler_GetUserByID() {
	req, _ := http.NewRequest("GET", "/api/users/"+uuid1, nil)
	w := httptest.NewRecorder()

	req = mux.SetURLVars(req, map[string]string{"id": uuid1})

	expectedUser := &models.User{ID: uuid1, Name: "John Doe", Username: "johndoe"}
	suite.mockRepo.On("FindByID", mock.Anything, uuid1).Return(expectedUser, nil)

	suite.userHandler.GetUserByID(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	suite.Equal(http.StatusOK, resp.StatusCode)

	var res handlers.SuccessResponse[models.User]
	json.NewDecoder(resp.Body).Decode(&res)
	suite.Equal(uuid1, res.Data.ID)
}
