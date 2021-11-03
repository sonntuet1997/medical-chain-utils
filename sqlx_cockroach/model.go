package sqlxcockroach

import (
	"github.com/gofrs/uuid"
)

type ModelBase struct {
	ID        uuid.UUID `db:"id"`
	UpdatedAt int64     `db:"updated_at"`
	CreatedAt int64     `db:"created_at"`
	DeletedAt int64     `db:"deleted_at"`
}
