package check

import "github.com/gofrs/uuid"

type Repo interface {
	GetAllByRoundID(rID uint) ([]*Check, error)
	GetAllByServiceID(sID uuid.UUID) ([]*Check, error)
	GetByRoundServiceID(rID uint, sID uuid.UUID) (*Check, error)
	Delete(rID uint, sID uuid.UUID) error
	GetAll() ([]*Check, error)
	Store(u []*Check) error
	Upsert(u []*Check) error
	TruncateTable() error
}
