package vscale

import (
	"fmt"
	"net/url"

	"github.com/parnurzeal/gorequest"
)

type Client struct {
	token       string
	ApiEndpoint *url.URL
}

const (
	defaultApiEndpoint = "https://api.vscale.io/v1/"
)

func NewClient(token string, endpoint string) (*Client, error) {
	if token == "" {
		return nil, fmt.Errorf("TOKEN should be requred")
	}

	rawurl, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	return &Client{
		token:       token,
		ApiEndpoint: rawurl,
	}, nil
}

type netmethod func(string) gorequest.SuperAgent

func (c *Client) Request(agent *gorequest.SuperAgent, params map[string]interface{}) (string, error) {
	_, body, errs := agent.
		Set("X-Token", c.token).
		Set("Accept", "application/json, text/plain").
		Set("Content-Type", "application/json").
		Send(params).
		End()
	if len(errs) > 0 {
		return "", errs[0]
	}
	return body, nil
}

func (c *Client) prepareUrl(path string) (string, error) {
	rawurl, err := c.ApiEndpoint.Parse(path)
	return rawurl.String(), err
}

func (c *Client) get(path string, params map[string]interface{}) (string, error) {
	rawurl, err := c.prepareUrl(path)
	if err != nil {
		return "", err
	}
	return c.Request(gorequest.New().Get(rawurl), params)
}

func (c *Client) post(path string, params map[string]interface{}) (string, error) {
	rawurl, err := c.prepareUrl(path)
	if err != nil {
		return "", err
	}
	return c.Request(gorequest.New().Post(rawurl), params)
}

func (c *Client) put(path string, params map[string]interface{}) (string, error) {
	rawurl, err := c.prepareUrl(path)
	if err != nil {
		return "", err
	}
	return c.Request(gorequest.New().Put(rawurl), params)
}

func (c *Client) patch(path string, params map[string]interface{}) (string, error) {
	rawurl, err := c.prepareUrl(path)
	if err != nil {
		return "", err
	}
	return c.Request(gorequest.New().Patch(rawurl), params)
}

func (c *Client) delete(path string, params map[string]interface{}) (string, error) {
	rawurl, err := c.prepareUrl(path)
	if err != nil {
		return "", err
	}
	return c.Request(gorequest.New().Delete(rawurl), params)
}
