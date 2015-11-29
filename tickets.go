package vscale

import (
	"fmt"
)

func (c *Client) Tickets() (string, error) {
	return c.get("tickets", nil)
}

func (c *Client) NewTicket(params map[string]interface{}) (string, error) {
	return c.post("tickets", params)
}

func (c *Client) TicketComments(id int) (string, error) {
	return c.get(fmt.Sprintf("tickets/%d/comments", id), nil)
}

func (c *Client) AddTicketComment(id int, params map[string]interface{}) (string, error) {
	return c.post(fmt.Sprintf("tickets/%d/comments", id), params)
}

func (c *Client) CloseTicket(id int) (string, error) {
	return c.post(fmt.Sprintf("tickets/%d/close", id), nil)
}
