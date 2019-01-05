package opennode

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
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

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(p))
	if err != nil {
		return nil, errors.WithStack(err)
	}
	req.Header.Set("Authorization", c.APIKey)
	req.Header.Set("Content-Type", "application/json")
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

	resource := struct {
		Data Charge `json:"data"`
	}{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &resource.Data, nil
}

func (c *Client) UpdateCharge(ch *Charge) error {
	if ch == nil {
		return nil
	}

	url := fmt.Sprintf("%s/charge/%s", c.APIEndpoint, ch.ID)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return errors.WithStack(err)
	}
	req.Header.Set("Authorization", c.APIKey)
	resp, err := c.client.Do(req)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.WithStack(err)
	}

	if resp.StatusCode != http.StatusOK {
		return errors.Errorf(
			"error during charge update: status: %d, body: %s",
			resp.StatusCode, string(body),
		)
	}

	resource := struct {
		Data Charge `json:"data"`
	}{}
	err = json.Unmarshal(body, &resource)
	if err != nil {
		return errors.WithStack(err)
	}

	*ch = resource.Data

	return nil
}

func (clt *Client) VerifyCharge(c *Charge) bool {
	mac := hmac.New(sha256.New, []byte(clt.APIKey))
	mac.Write([]byte(c.ID))
	expected := mac.Sum(nil)
	return hmac.Equal(expected, []byte(c.HashedOrder))
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
