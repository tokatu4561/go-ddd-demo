package user

type UserRepositoryInterface interface {
	FindByUserName(name *UserName) (*User, error)
	Find(id *UserId) (*User, error)
	Save(user *User) error
	FindByUserIds(ids []UserId) ([]User, error)
}