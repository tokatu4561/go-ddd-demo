package user

import "context"

type UserRepositoryInterface interface {
	FindByUserName(name *UserName) (*User, error)
	Find(id *UserId) (*User, error)
	Save(user *User, ctx context.Context) error
	FindByUserIds(ids []UserId) ([]User, error)
}