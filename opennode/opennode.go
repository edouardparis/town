package opennode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

type Client struct {
	client      *http.Client
	APIKey      string
	APIEndpoint string
}

func (c *Client) Dev() {
	c.APIEndpoint = "https://dev-api.opennode.co/v1"
}

func (c *Client) Prod() {
	c.APIEndpoint = "https://api.opennode.co/v1"
}

func (c *Client) CreateCharge(payload *ChargePayload) (*Charge, error) {
	url := fmt.Sprintf("%s/charges", c.APIEndpoint)

	p, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(p))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.Errorf("error during charge creation")
	}

	var charge Charge
	err = json.Unmarshal(body, &charge)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &charge, nil
}

type Config struct {
	Debug  bool   `json:"debug"`
	APIKey string `json:"api_key"`
}

func NewClient(c *Config) *Client {
	client := &Client{
		APIKey: c.APIKey,
		client: http.DefaultClient,
	}
	client.Prod()
	if c.Debug {
		client.Dev()
	}
	return client
}
