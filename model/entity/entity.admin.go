package entity

import (
	"database/sql"
	"time"
)

type AdminEntity struct {
	UID        string       `db:"uid"`
	Password   string       `db:"password"`
	Email      string       `db:"email"`
	Username   string       `db:"username"`
	FullName   string       `db:"full_name"`
	InsertedAt time.Time    `db:"inserted_at"`
	UpdatedAt  sql.NullTime `db:"updated_at,omitempty"`
}
