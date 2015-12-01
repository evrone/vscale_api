package vscale

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestSSHKeysList(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/sshkeys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		response := `
[
  {
    "id":15,
    "key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFPEvMAMf4DZN34lReVkVklkAET+H57pSK1aBJ4NX2Nd+Asv5iJWBKBELYmDaRxfx6Y8nS6uYPU3EJ+qBI91NOjJSTPchjWSGGWv4SpkxXBEMjWyUob8BFn5rEjEtDMBsR8xPurcs1vkaoet6A9eXw67pVVwsdh48hKQc0DSaYVtmb08ex4uWoadzixM3GUfMnW/2AQK75dyJIhvOTHwxZEeynDFgI9fNgWLbre62uCMmlCyvMYnpG8apz9igscZycimTOWqPvhMskdHBtYBFHFAg/50NH38L52cMjP3/j1CG+crC7l6ij4e7DAVq43jyL6sWlbvpLWPPj/4MIf2W7 user@host.local",
    "name":"publickeyname"
 },
 {
    "id":16,
    "key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABCCCBAQDMt+AlVdjaH1mTPbUM5bI6wYxrOriVTZh2DrA9XhSP1ndTtVsDT0QJanQSw8BvVovu5n4Ves0FqbBR+nmuHKmta+YevTlvfrbjerSQj01mTLI4VgM6RL/XAygroNIKVEmBSKJj9eitl1G1UP5f3vHiPunt2jCNlNK0rexX1klnhwTroHGB9kMjjfPJJMBiV/M0klcAQtBYaAw6VUvA5mjfKnTRJ8HtD4bEJ0O6p85h9UhxelAfi7Giu5MAvzIfwAQRrvQabdgD44OklM+oHhz/tbZo1dv2znUDgFsfKu0MXtMQPFklwlk+BZQNRsS5/1WGn5knyCqDgWUF/LD4QDIJ user@host.local",
    "name":"somekeyname"
  }
]`

		fmt.Fprint(w, response)
	})

	sshkeys, _, err := client.SSHKey.List()
	if err != nil {
		t.Errorf("SSHKey.List returned error: %v", err)
	}

	expected := &[]SSHKey{
		SSHKey{
			ID:   15,
			Key:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDFPEvMAMf4DZN34lReVkVklkAET+H57pSK1aBJ4NX2Nd+Asv5iJWBKBELYmDaRxfx6Y8nS6uYPU3EJ+qBI91NOjJSTPchjWSGGWv4SpkxXBEMjWyUob8BFn5rEjEtDMBsR8xPurcs1vkaoet6A9eXw67pVVwsdh48hKQc0DSaYVtmb08ex4uWoadzixM3GUfMnW/2AQK75dyJIhvOTHwxZEeynDFgI9fNgWLbre62uCMmlCyvMYnpG8apz9igscZycimTOWqPvhMskdHBtYBFHFAg/50NH38L52cMjP3/j1CG+crC7l6ij4e7DAVq43jyL6sWlbvpLWPPj/4MIf2W7 user@host.local",
			Name: "publickeyname",
		},
		SSHKey{
			ID:   16,
			Key:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABCCCBAQDMt+AlVdjaH1mTPbUM5bI6wYxrOriVTZh2DrA9XhSP1ndTtVsDT0QJanQSw8BvVovu5n4Ves0FqbBR+nmuHKmta+YevTlvfrbjerSQj01mTLI4VgM6RL/XAygroNIKVEmBSKJj9eitl1G1UP5f3vHiPunt2jCNlNK0rexX1klnhwTroHGB9kMjjfPJJMBiV/M0klcAQtBYaAw6VUvA5mjfKnTRJ8HtD4bEJ0O6p85h9UhxelAfi7Giu5MAvzIfwAQRrvQabdgD44OklM+oHhz/tbZo1dv2znUDgFsfKu0MXtMQPFklwlk+BZQNRsS5/1WGn5knyCqDgWUF/LD4QDIJ user@host.local",
			Name: "somekeyname",
		},
	}

	if !reflect.DeepEqual(sshkeys, expected) {
		t.Errorf("SSHKey.List returned %+v, expected %+v", sshkeys, expected)
	}
}

func TestSSHKeysCreate(t *testing.T) {
	setup()
	defer teardown()

	createRequest := &SSHKeyCreateRequest{
		Name: "newkey",
		Key:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQClleNF5kP55FOHSdU+1bPRJ7Q4o+jOeYuM+XpasTOhNYVaZRTQYmas/7FF7YImu34kbF2jQpX2GezxafG8E+BqQyiDa0Cb18jmkHDlZNo62W16tuFMc5rsB6yRJPc9WUMC84xxgxGIVSZZAbv9wFTLyK3k6zRdnNXsfefzh6XL4jEh/I0/gnw9phs3MCSvAjHw6futhybaukEwQI5oq8WNB1JRQoNN95Dt+sAwM4Ur6CdbgLtn5jJdRyOHMM/fNfSwLxbr+Lm4xLpP+Fyuyd6gvUebR7fdCSD+2iBBpaLz5LTAX7XXOB/aizTXIIJbSbZ1PjBUmX/uS1cFLYGVfRYT user@host.local",
	}

	mux.HandleFunc("/v1/sshkeys", func(w http.ResponseWriter, r *http.Request) {
		v := new(SSHKeyCreateRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatalf("decode json: %v", err)
		}

		testMethod(t, r, "POST")

		if !reflect.DeepEqual(v, createRequest) {
			t.Errorf("Request body: %+v, expected: %+v", v, createRequest)
		}

		response := `
{
	"key":"ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQClleNF5kP55FOHSdU+1bPRJ7Q4o+jOeYuM+XpasTOhNYVaZRTQYmas/7FF7YImu34kbF2jQpX2GezxafG8E+BqQyiDa0Cb18jmkHDlZNo62W16tuFMc5rsB6yRJPc9WUMC84xxgxGIVSZZAbv9wFTLyK3k6zRdnNXsfefzh6XL4jEh/I0/gnw9phs3MCSvAjHw6futhybaukEwQI5oq8WNB1JRQoNN95Dt+sAwM4Ur6CdbgLtn5jJdRyOHMM/fNfSwLxbr+Lm4xLpP+Fyuyd6gvUebR7fdCSD+2iBBpaLz5LTAX7XXOB/aizTXIIJbSbZ1PjBUmX/uS1cFLYGVfRYT user@host.local",
	"id":16,
	"name":"newkey"
}`

		fmt.Fprint(w, response)
	})

	sshkey, _, err := client.SSHKey.Create(createRequest)
	if err != nil {
		t.Errorf("SSHKeys.Create returned error: %v", err)
	}

	expected := &SSHKey{
		ID:   16,
		Name: "newkey",
		Key:  "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQClleNF5kP55FOHSdU+1bPRJ7Q4o+jOeYuM+XpasTOhNYVaZRTQYmas/7FF7YImu34kbF2jQpX2GezxafG8E+BqQyiDa0Cb18jmkHDlZNo62W16tuFMc5rsB6yRJPc9WUMC84xxgxGIVSZZAbv9wFTLyK3k6zRdnNXsfefzh6XL4jEh/I0/gnw9phs3MCSvAjHw6futhybaukEwQI5oq8WNB1JRQoNN95Dt+sAwM4Ur6CdbgLtn5jJdRyOHMM/fNfSwLxbr+Lm4xLpP+Fyuyd6gvUebR7fdCSD+2iBBpaLz5LTAX7XXOB/aizTXIIJbSbZ1PjBUmX/uS1cFLYGVfRYT user@host.local",
	}

	if !reflect.DeepEqual(sshkey, expected) {
		t.Errorf("SSHKeys.Create returned %+v, expected %+v", sshkey, expected)
	}
}

func TestSSHKeysDelete(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/v1/sshkeys/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		fmt.Fprint(w, `{}`)
	})

	_, err := client.SSHKey.Delete(1)
	if err != nil {
		t.Errorf("SSHKeys.Delete returned error: %v", err)
	}
}
