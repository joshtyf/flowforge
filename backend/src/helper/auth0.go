package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joshtyf/flowforge/src/database/models"
)

type ManagementApiToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
}

type Auth0UserId struct {
	UserId string `json:"user_id"`
}

func CreateUserInAuth0(user *models.UserModel) (*models.UserModel, error) {
	url, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/api/v2/users")
	if err != nil {
		return nil, err
	}

	token, err := getManagementApiToken()
	if err != nil {
		return nil, err
	}

	str := `{
		"email": "%s",
		"name": "%s",
		"connection": "Username-Password-Authentication",
		"password": "%s"
	}`
	jsonStr := []byte(fmt.Sprintf(str, user.Email, user.Name, "17April1998"))
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var userId *Auth0UserId
	err = json.NewDecoder(resp.Body).Decode(&userId)
	if err != nil {
		return nil, err
	}
	user.UserId = userId.UserId
	return user, nil
}

func getManagementApiToken() (*ManagementApiToken, error) {
	issuerURL, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/oauth/token")
	if err != nil {
		return nil, err
	}
	data := url.Values{}
	data.Set("client_secret", os.Getenv("MANAGEMENT_API_SECRET"))
	data.Set("client_id", os.Getenv("MANAGEMENT_API_CLIENT"))
	data.Set("grant_type", "client_credentials")
	data.Set("audience", os.Getenv("MANAGEMENT_API_AUDIENCE"))

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, issuerURL.String(), strings.NewReader(data.Encode())) // URL-encoded payload
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var token *ManagementApiToken
	err = json.NewDecoder(resp.Body).Decode(&token)
	if err != nil {
		return nil, err
	}
	return token, nil
}
