package user

type UserService struct {
	userRepository UserRepositoryInterface
}

func NewUserService(userRepository UserRepositoryInterface) (*UserService, error) {
	return &UserService{userRepository: userRepository}, nil
}

func (us *UserService) Exists(user *User) (bool, error) {
	user, err := us.userRepository.FindByUserName(user.Name())
	if err != nil {
		return false, err
	}
	return user != nil, nil
}
