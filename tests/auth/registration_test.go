package auth

import (
	"GoDeveloperVK-testTask/auth"
	"GoDeveloperVK-testTask/auth/jwt"
	"GoDeveloperVK-testTask/auth/repository"
	"GoDeveloperVK-testTask/tests"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegistration(t *testing.T) {
	tests.StartApp()
	repository.DeleteAllUsers()

	server := httptest.NewServer(http.HandlerFunc(auth.RegistrationHandler))
	client := &http.Client{}

	body := []byte(`{
  		"first_name": "Name",
  		"last_name": "Name",
  		"email": "email@email.com",
  		"password": "password",
  		"role": "admin"
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
	err = json.NewDecoder(response.Body).Decode(jwtToken)
	if err != nil {
		t.Error(err)
	}

	claims, err := jwt.VerifyToken(jwtToken.Token)
	if claims.Username != "email@email.com" && claims.Role != "admin" {
		t.Error("invalid token after registration")
	}

	repository.DeleteAllUsers()
}
