package user

import (
	"fmt"
)

type UserId struct {
	Id string
}

func NewUserId(id string) (_ *UserId, err error) {
	userId := new(UserId)
	if id == "" {
		return nil, fmt.Errorf("userId is required")
	}
	userId.Id = id
	return userId, nil
}
