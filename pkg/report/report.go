package report

import (
	"github.com/gofrs/uuid"
	"time"
)

type Report struct {
	ID        uint64    `json:"id,omitempty"`
	Cache     string    `json:"cache,omitempty" gorm:"not null;default:'{}'"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d Report) TableName() string {
	return "report"
}

func NewReport() *Report {
	c := &Report{}
	c.ID = 1
	c.Cache = "{}"
	return c
}

type SimpleReport struct {
	Round uint64
	Teams map[uuid.UUID]*SimpleTeam
}

// Todo: Look into including Paused service (Not scored for the latest round, but points from rounds where service wasnt paused will be included into the total (Also shows up in the report as yellow))
type SimpleTeam struct {
	Hosts       map[uuid.UUID]*SimpleHost
	Name        string
	Enabled     bool
	Hidden      bool
	TotalPoints uint64
}

type SimpleHostGroup struct {
	ID      uuid.UUID
	Name    string
	Enabled bool
}

type SimpleHost struct {
	HostGroup *SimpleHostGroup
	Address   string
	Services  map[uuid.UUID]*SimpleService
	Enabled   bool
}

type SimpleServiceGroup struct {
	ID      uuid.UUID
	Name    string
	Enabled bool
}

type SimpleService struct {
	Enabled            bool
	Name               string
	DisplayName        string
	Passed             bool
	Log                string
	Err                string
	Weight             uint64
	Points             uint64
	PointsBoost        uint64
	Properties         map[string]*SimpleProperty
	SimpleServiceGroup *SimpleServiceGroup
}

type SimpleProperty struct {
	Value  string
	Status string
}
