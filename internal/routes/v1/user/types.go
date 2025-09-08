package user

import "time"
import "github.com/google/uuid"

type user struct {
	Id          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}
