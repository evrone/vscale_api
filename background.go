package vscale

func (c *Client) Locations() (string, error) {
	return c.get("locations", nil)
}

func (c *Client) Images() (string, error) {
	return c.get("images", nil)
}
