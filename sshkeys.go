package vscale

// import (
// 	"encoding/json"
// 	"fmt"
// )
//
// type SSHKey struct {
// 	ID        int
// 	Name, Key string
// }
//
// func (c *Client) SSHKeys() (*[]SSHKey, error) {
// 	ret, err := c.get("sshkeys", nil)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	keys := make([]SSHKey, 0)
// 	err = json.Unmarshal([]byte(ret), &keys)
// 	return &keys, err
// }
//
// func (c *Client) NewSSHKey(params map[string]interface{}) (*SSHKey, error) {
// 	ret, err := c.post("sshkeys", params)
// 	if err != nil {
// 		return nil, err
// 	}
//
// 	keys := make([]SSHKey, 0)
// 	fmt.Println(ret)
// 	err = json.Unmarshal([]byte(ret), &keys)
//
// 	if len(keys) > 0 {
// 		return &keys[0], nil
// 	} else {
// 		return nil, err
// 	}
// }
//
// func (c *Client) DeleteSSHKey(id int) (string, error) {
// 	return c.delete(fmt.Sprintf("sshkeys/%d", id), nil)
// }
