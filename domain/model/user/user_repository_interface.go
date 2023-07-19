package user

type UserRepositoryInterface interface {
	FindByUserName(name *UserName) (*User, error)
	Save(user *User) error
}