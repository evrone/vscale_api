package vscale

import (
	"fmt"
)

func (c *Client) Tokens() (string, error) {
	return c.get("tokens", nil)
}

func (c *Client) CreateToken(params map[string]interface{}) (string, error) {
	return c.post("tokens", params)
}

func (c *Client) GetToken(id int) (string, error) {
	return c.get(fmt.Sprintf("tokens/%d", id), nil)
}

func (c *Client) UpdateToken(id int, params map[string]interface{}) (string, error) {
	return c.post(fmt.Sprintf("tokens/%d", id), params)
}

func (c *Client) DeleteToken(id int) (string, error) {
	return c.delete(fmt.Sprintf("tokens/%d", id), nil)
}
