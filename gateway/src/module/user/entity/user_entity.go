package entity

import (
	"time"
)

type User struct {
	Id        int       `db:"id" json:"id"`
	Name      string    `db:"name" json:"name,omitempty"`
	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}
