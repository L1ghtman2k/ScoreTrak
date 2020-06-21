package client

import (
	"ScoreTrak/pkg/team"
	"fmt"
)

type teamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &teamClient{c}
}

func (t teamClient) DeleteByName(name string) error {
	return t.s.genericDelete(fmt.Sprintf("/team/%s", name))
}

func (t teamClient) GetAll() ([]*team.Team, error) {
	var tm []*team.Team
	err := t.s.genericGet(&tm, "/team")
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) GetByName(name string) (*team.Team, error) {
	tm := &team.Team{}
	err := t.s.genericGet(tm, fmt.Sprintf("/team/%s", name))
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t teamClient) Store(u *team.Team) error {
	return t.s.genericStore(u, fmt.Sprintf("/team"))
}

func (t teamClient) UpdateByName(u *team.Team) error {
	return t.s.genericUpdate(u, fmt.Sprintf("/team/%s", u.Name))
}
