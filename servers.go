package vscale

import (
	"fmt"
)

func (c *Client) Scalets() (string, error) {
	return c.get("scalets", nil)
}

func (c *Client) NewScalet(params map[string]interface{}) (string, error) {
	return c.post("scalets", params)
}

func (c *Client) FindScalet(id int) (string, error) {
	return c.get(fmt.Sprintf("scalets/%d", id), nil)
}

func (c *Client) RestartScalet(id int) (string, error) {
	return c.patch(fmt.Sprintf("scalets/%d/restart", id), nil)
}

func (c *Client) StopScalet(id int) (string, error) {
	return c.patch(fmt.Sprintf("scalets/%d/start", id), nil)
}

func (c *Client) StartScalet(id int) (string, error) {
	return c.patch(fmt.Sprintf("scalets/%d/start", id), nil)
}

func (c *Client) UpgradeScalet(id int, params map[string]interface{}) (string, error) {
	return c.post(fmt.Sprintf("scalets/%d/upgrade", id), params)
}

func (c *Client) DeleteScalet(id int) (string, error) {
	return c.delete(fmt.Sprintf("scalets/%d", id), nil)
}

func (c *Client) Tasks() (string, error) {
	return c.get("tasks", nil)
}

func (c *Client) ScaletSSHKeys(id int, params map[string]interface{}) (string, error) {
	return c.patch(fmt.Sprintf("sshkeys/scalets/%d", id), params)
}
