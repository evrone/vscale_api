package vscale

func (c *Client) Payments() (string, error) {
	return c.get("billing/payments", nil)
}

func (c *Client) Consumption(params map[string]interface{}) (string, error) {
	return c.get("billing/consumption", params)
}
