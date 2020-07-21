package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/property"
)

type PropertyClient struct {
	s ScoretrakClient
}

func NewPropertyClient(c ScoretrakClient) property.Serv {
	return &PropertyClient{c}
}

func (s PropertyClient) Delete(id uint64) error {
	return s.s.genericDelete(fmt.Sprintf("/property/%d", id))
}

func (s PropertyClient) GetAll() ([]*property.Property, error) {
	var sg []*property.Property
	err := s.s.genericGet(&sg, "/property")
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) GetByID(id uint64) (*property.Property, error) {
	sg := &property.Property{}
	err := s.s.genericGet(sg, fmt.Sprintf("/property/%d", id))
	if err != nil {
		return nil, err
	}
	return sg, nil
}

func (s PropertyClient) Store(u *property.Property) error {
	return s.s.genericStore(u, fmt.Sprintf("/property"))
}

func (s PropertyClient) Update(u *property.Property) error {
	return s.s.genericUpdate(u, fmt.Sprintf("/property/%d", u.ID))
}
