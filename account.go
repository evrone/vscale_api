package vscale

func (c *Client) Account() (string, error) {
	return c.get("account", nil)
}
