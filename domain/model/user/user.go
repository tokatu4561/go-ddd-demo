package user

import (
	"reflect"
)

type User struct {
	userId UserId
	name   UserName
	isPremium bool
}

func NewUser(userId UserId, name UserName, isPremium bool) (*User, error) {
	user := new(User)

	user.userId = userId
	user.name = name
	user.isPremium = isPremium
	return user, nil
}

func (user *User) ChangeUserName(name UserName) (err error) {
	user.name = name
	return nil
}

func (user *User) Id() *UserId {
	return &user.userId
}

func (user *User) Name() *UserName {
	return &user.name
}

func (user *User) Equals(other *UserId) bool {
	return reflect.DeepEqual(user.userId, other)
}

func (user *User) IsPremium() bool {
	return user.isPremium
}
