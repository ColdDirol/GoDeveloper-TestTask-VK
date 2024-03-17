package auth

import (
	"GoDeveloperVK-testTask/auth"
	"GoDeveloperVK-testTask/auth/jwt"
	"GoDeveloperVK-testTask/auth/repository"
	"GoDeveloperVK-testTask/models"
	"GoDeveloperVK-testTask/tests"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	tests.StartApp()
	repository.DeleteAllUsers()
	server := httptest.NewServer(http.HandlerFunc(auth.LoginHandler))
	client := &http.Client{}

	user := models.User{
		FirstName: "Name",
		LastName:  "Name",
		Email:     "email@email.com",
		Password:  jwt.Sha256EncodeWithSalt("password"),
		Role:      "admin",
	}
	err := repository.AddUser(user)
	if err != nil {
		t.Error(err)
	}

	body := []byte(`{
  		"email": "email@email.com",
  		"password": "password"
	}`)

	request, err := http.NewRequest("POST", server.URL, bytes.NewBuffer(body))
	if err != nil {
		t.Error(err)
	}
	request.Header.Add("Content-Type", "application/json")

	response, err := client.Do(request)
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", response.StatusCode)
	}

	jwtToken := &jwt.JWT{}
	json.NewDecoder(response.Body).Decode(jwtToken)

	claims, err := jwt.VerifyToken(jwtToken.Token)
	if claims.Username != "email@email.com" && claims.Role != "admin" {
		t.Error("invalid token after registration")
	}

	repository.DeleteAllUsers()
}
