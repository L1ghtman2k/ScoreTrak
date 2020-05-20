package host_group

import "ScoreTrak/pkg/host"

// Host Group model represents a set of hosts that have a common purpose, but are in different teams. For instance team 1 web, and team 2 web would bellong to a host group Web
type HostGroup struct {
	ID uint64 `json:"id,omitempty"`

	Name string `json:"name" gorm:"not null"`

	// Enables or disables scoring for a given host group. In case you want to stop scoring a set of simalar hosts, you can set this property to false
	Enabled *bool `json:"enabled,omitempty" gorm:"not null default: false"`

	Hosts []host.Host `gorm:"foreignkey:HostGroupID"`
}