package vscale

func (c *Client) Rplans() (string, error) {
	return c.get("rplans", nil)
}

func (c *Client) Prices() (string, error) {
	return c.get("billing/prices", nil)
}
