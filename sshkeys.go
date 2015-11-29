package vscale

import (
	"fmt"
)

func (c *Client) SSHKeys() (string, error) {
	return c.get("sshkeys", nil)
}

func (c *Client) NewSSHKey(params map[string]interface{}) (string, error) {
	return c.post("sshkeys", params)
}

func (c *Client) DeleteSSHKey(id int) (string, error) {
	return c.delete(fmt.Sprintf("sshkeys/%d", id), nil)
}
