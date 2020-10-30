package report

import (
	"github.com/gofrs/uuid"
	"time"
)

type Report struct {
	ID        uint      `json:"id,omitempty"`
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
	Round uint
	Teams map[uuid.UUID]*SimpleTeam
}

type SimpleTeam struct {
	Hosts   map[uuid.UUID]*SimpleHost
	Name    string
	Enabled bool
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
	Weight             uint
	Points             uint
	PointsBoost        uint
	Properties         map[string]*SimpleProperty
	SimpleServiceGroup *SimpleServiceGroup
}

type SimpleProperty struct {
	Value  string
	Status string
}
