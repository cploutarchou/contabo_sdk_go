package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	endpoint     string
	ClientID     string
	ClientSecret string
	APIUser      string
	Password     string
	BearerToken  string
}

//getBearerToken returns the BearerToken
func (c *Client) getBearerToken() error {
	params := url.Values{}
	params.Add("client_id", c.ClientID)
	params.Add("client_secret", c.ClientSecret)
	params.Add("username", c.APIUser)
	params.Add("password", c.Password)
	params.Add("grant_type", "password")
	body := strings.NewReader(params.Encode())
	req, err := http.NewRequest("POST", AuthUrl, body)
	if err != nil {
		return fmt.Errorf("Error creating request: %v ", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("Error sending request: %v ", err)
	}
	var v struct {
		AccessToken string `json:"access_token"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return fmt.Errorf("Error decoding response: %v ", err)
	}
	if v.AccessToken == "" {
		return fmt.Errorf("Error decoding response: %v ", err)
	}
	c.BearerToken = v.AccessToken
	return nil
}
