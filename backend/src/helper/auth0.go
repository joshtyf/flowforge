package helper

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

type Auth0Identity struct {
	Connection string `json:"connection"`
	UserId     string `json:"user_id"`
	Provider   string `json:"provider"`
	IsSocial   bool   `json:"isSocial"`
}

type Auth0UserDetails struct {
	Email      string          `json:"email"`
	Identities []Auth0Identity `json:"identities"`
}

func CreateUsersInAuth0(users []models.UserModel, passwords []string) ([]models.UserModel, error) {
	token, err := getManagementApiToken()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(users); i++ {
		err = createUserInAuth0(&users[i], passwords[i], token)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func GetUserIdForUsers(users []models.UserModel) ([]models.UserModel, error) {
	token, err := getManagementApiToken()
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(users); i++ {
		err = getAuth0UserIdByEmail(&users[i], token)
		if err != nil {
			return nil, err
		}
	}

	return users, nil
}

func GetAuth0UserDetailsForUser(user *models.UserModel) error {
	token, err := getManagementApiToken()
	if err != nil {
		return err
	}

	url, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/api/v2/users/" + user.UserId)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("response code %d encountered while retrieving user details: %s", resp.StatusCode, string(bytes))
	}

	var userDetails *Auth0UserDetails
	err = json.NewDecoder(resp.Body).Decode(&userDetails)
	if err != nil {
		return err
	}

	user.Email = userDetails.Email
	user.IdentityProvider = userDetails.Identities[0].Provider
	return nil
}

func getAuth0UserIdByEmail(user *models.UserModel, token *ManagementApiToken) error {
	url, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/api/v2/users-by-email")
	if err != nil {
		return err
	}

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return err
	}

	q := req.URL.Query()
	q.Add("email", user.Email)
	q.Add("fields", "user_id")
	req.URL.RawQuery = q.Encode()
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("response code %d encountered while retrieving user id: %s", resp.StatusCode, string(bytes))
	}

	var userId []*Auth0UserId
	err = json.NewDecoder(resp.Body).Decode(&userId)
	if err != nil {
		return err
	}
	if len(userId) <= 0 {
		return fmt.Errorf("user with %s has not been created", user.Email)
	}

	user.UserId = userId[0].UserId
	return nil
}

func createUserInAuth0(user *models.UserModel, password string, token *ManagementApiToken) error {
	url, err := url.Parse("https://" + os.Getenv("AUTH0_DOMAIN") + "/api/v2/users")
	if err != nil {
		return err
	}

	str := `{
		"email": "%s",
		"name": "%s",
		"connection": "Username-Password-Authentication",
		"password": "%s"
	}`
	jsonStr := []byte(fmt.Sprintf(str, user.Email, user.Name, password))
	req, err := http.NewRequest("POST", url.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", token.TokenType, token.AccessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	var userId *Auth0UserId
	err = json.NewDecoder(resp.Body).Decode(&userId)
	if err != nil {
		return err
	}
	user.UserId = userId.UserId
	return nil
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
