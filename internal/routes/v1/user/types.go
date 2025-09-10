package user

import "time"
import "github.com/google/uuid"

type User struct {
	Id          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Created     time.Time `json:"created"`
	LastUpdated time.Time `json:"last_updated"`
}

type UserInput struct {
	Email string `json:"email"`
}

type UserPost struct {
	Users []UserInput `json:"users"`
}

type UserPut = UserInput
