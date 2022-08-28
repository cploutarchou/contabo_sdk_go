package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

type URL string
type Method string
type Action string

const (
	BaseUrl = "https://api.contabo.com"

	AuthUrl                    = "https://auth.contabo.com/auth/realms/contabo/protocol/openid-connect/token"
	ComputeInstancesUrl URL    = BaseUrl + "/v1/compute/instances"
	GET                 Method = "GET"
	PUT                 Method = "PUT"
	POST                Method = "POST"
	DELETE              Method = "DELETE"
	PATCH               Method = "PATCH"
	START               Action = "start"
	REBOOT              Action = "restart"
	STOP                Action = "stop"
)

type Config struct {
	ClientID     string
	ClientSecret string
	APIUser      string
	Password     string
}

type Client struct {
	bearerToken string
	Method      Method
	URL         URL
	Body        interface{}
	Instances   Instances
}

// GetBearerToken returns the BearerToken
func (c *Config) GetBearerToken() error {
	params := url.Values{}
	params.Add("client_id", c.ClientID)
	params.Add("client_secret", c.ClientSecret)
	params.Add("username", c.APIUser)
	params.Add("password", c.Password)
	params.Add("grant_type", "password")
	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", AuthUrl, body)
	if err != nil {
		return fmt.Errorf("error creating request: %v ", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v ", err)
	}
	var v struct {
		AccessToken string `json:"access_token"`
	}

	if err = json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return fmt.Errorf("error decoding response: %v ", err)
	}
	if v.AccessToken == "" {
		return fmt.Errorf("error decoding response: %v ", err)
	}
	err = os.Setenv("CONTABO_BEARER_TOKEN", v.AccessToken)
	if err != nil {
		return err
	}
	return nil
}

// Init creates a new Client for the specified service account token
func Init(config *Config) error {
	err := config.GetBearerToken()
	if err != nil {
		return err
	}
	return nil
}

// Do The do request method to  make a new Client request.
func Do(method Method, Url URL, data interface{}) ([]byte, error) {
	if os.Getenv("CONTABO_BEARER_TOKEN") == "" {
		return nil, fmt.Errorf("something went wrong! Unable to get bearer token! ")

	}
	b, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	// convert byte slice to io.Reader
	reader := bytes.NewReader(b)
	req, err := http.NewRequest(string(method), string(Url), reader)
	if err != nil {
		return nil, err
	}
	id := uuid.New()
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("CONTABO_BEARER_TOKEN")))
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

// randomString returns a random string from the given specs.
func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
