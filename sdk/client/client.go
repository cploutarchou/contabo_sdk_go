package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Config struct {
	endpoint     string
	ClientID     string
	ClientSecret string
	APIUser      string
	Password     string
}

type client struct {
	bearerToken string
	Method      Method
	URL         URL
	Body        interface{}
}

// GetBearerToken returns the BearerToken
func (c *Config) GetBearerToken() (*string, error) {
	params := url.Values{}
	params.Add("client_id", c.ClientID)
	params.Add("client_secret", c.ClientSecret)
	params.Add("username", c.APIUser)
	params.Add("password", c.Password)
	params.Add("grant_type", "password")
	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", AuthUrl, body)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v ", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v ", err)
	}
	var v struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, fmt.Errorf("error decoding response: %v ", err)
	}
	if v.AccessToken == "" {
		return nil, fmt.Errorf("error decoding response: %v ", err)
	}

	return &v.AccessToken, nil
}

func CreateClient(config *Config) (*client, error) {
	token, err := config.GetBearerToken()
	if err != nil {
		return nil, err
	}
	return &client{bearerToken: *token}, nil
}

func (c *client) Do() ([]byte, error) {
	b, err := json.Marshal(c.Body)
	if err != nil {
		panic(err)
	}
	// convert byte slice to io.Reader
	reader := bytes.NewReader(b)
	req, err := http.NewRequest(string(c.Method), string(c.URL), reader)
	if err != nil {
		// handle err
	}
	id := uuid.New()
	fmt.Println("Generated UUID:")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.bearerToken))
	req.Header.Set("X-Request-Id", id.String())
	req.Header.Set("X-Trace-Id", randomString(8))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
