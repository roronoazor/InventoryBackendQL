package auth

import "time"

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role"`
	IsAdmin   string    `json:"isAdmin"`
	CreatedAt time.Time `json:"createdAt"`
}
