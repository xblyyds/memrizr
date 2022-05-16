package model

import "github.com/google/uuid"

type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"password"`
	Name     string    `db:"name" json:"name"`
	ImageURL string    `db:"image_url" json:"imageURL"`
	Website  string    `db:"website" json:"website"`
}
