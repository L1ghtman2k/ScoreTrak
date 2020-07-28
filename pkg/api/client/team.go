package client

import (
	"fmt"
	"github.com/L1ghtman2k/ScoreTrak/pkg/team"
)

type TeamClient struct {
	s ScoretrakClient
}

func NewTeamClient(c ScoretrakClient) team.Serv {
	return &TeamClient{c}
}

func (t TeamClient) Delete(id uint64) error {
	return t.s.genericDelete(fmt.Sprintf("/team/%d", id))
}

func (t TeamClient) GetAll() ([]*team.Team, error) {
	var tm []*team.Team
	err := t.s.genericGet(&tm, "/team")
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t TeamClient) GetByID(id uint64) (*team.Team, error) {
	tm := &team.Team{}
	err := t.s.genericGet(tm, fmt.Sprintf("/team/%d", id))
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (t TeamClient) Store(u *team.Team) error {
	return t.s.genericStore(u, fmt.Sprintf("/team"))
}

func (t TeamClient) Update(u *team.Team) error {
	return t.s.genericUpdate(u, fmt.Sprintf("/team/%d", u.ID))
}
