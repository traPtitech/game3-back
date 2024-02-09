package domain

import "time"

type Image struct {
	Image     []byte    `db:"image"`
	UpdatedAt time.Time `db:"updated_at"`
}

type Icon struct {
	Image     []byte    `db:"icon"`
	UpdatedAt time.Time `db:"updated_at"`
}
