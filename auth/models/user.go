package auth

import "time"

type User struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	IsAdmin   string    `json:"isAdmin"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
