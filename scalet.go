package vscale

import (
	"fmt"
)

const (
	scaletsBasePath = "/v1/scalets"
)

type ScaletService interface {
	List() (*[]Scalet, *Response, error)
	GetByID(int) (*Scalet, *Response, error)
}

type ScaletServiceOp struct {
	client *Client
}

var _ ScaletService = &ScaletServiceOp{}

type Scalet struct {
	Name           string         `json:"name,omitempty"`
	Hostname       string         `json:"hostname,omitempty"`
	Locked         bool           `json:"locked,omitempty"`
	Location       string         `json:"location,omitempty"`
	Rplan          string         `json:"rplan,omitempty"`
	Active         bool           `json:"active,omitempty"`
	Keys           []ScaletKey    `json:"keys,omitempty"`
	PublicAddress  *ScaletAddress `json:"public_address,omitempty"`
	Status         string         `json:"status,omitempty"`
	MadeFrom       string         `json:"made_from,omitempty"`
	CTID           int            `json:"ctid,omitempty"`
	PrivateAddress *ScaletAddress `json:"private_address,omitempty"`
}

type ScaletKey struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type ScaletAddress struct {
	Address string `json:"address,omitempty"`
	Netmask string `json:"netmask,omitempty"`
	Gateway string `json:"gateway,omitempty"`
}

func (s Scalet) String() string {
	return Stringify(s)
}

func (s ScaletServiceOp) List() (*[]Scalet, *Response, error) {
	path := scaletsBasePath
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalets := &[]Scalet{}
	resp, err := s.client.Do(req, scalets)
	if err != nil {
		return nil, nil, err
	}
	return scalets, resp, err
}

func (s ScaletServiceOp) GetByID(ctid int) (*Scalet, *Response, error) {
	if ctid < 1 {
		return nil, nil, NewArgError("scaletID", "cannot be less than 1")
	}
	path := fmt.Sprintf("%s/%d", scaletsBasePath, ctid)
	req, err := s.client.NewRequest("GET", path, nil)
	if err != nil {
		return nil, nil, err
	}

	scalet := &Scalet{}
	resp, err := s.client.Do(req, scalet)
	if err != nil {
		return nil, nil, err
	}
	return scalet, resp, err
}
