package application

import (
	"github.com/tokatu4561/go-ddd-demo/domain/model/user"
	"github.com/tokatu4561/go-ddd-demo/domain/service"
)

type UserApplicationService struct {
	userRepository user.UserRepositoryInterface
	userService    service.CircleService
}

type GetUserListCommand struct {
	UserIds []string
}

type UserData struct {
	Id      user.UserId
	Name    user.UserName
}

func (uas *CircleApplicationService) GetUsers(command GetUserListCommand) (usersData []UserData, err error) {
	userIds := make([]user.UserId, len(command.UserIds))
	for i, userId := range command.UserIds {
		dId, err := user.NewUserId(userId)
		userIds[i] = *dId
		if err != nil {
			return nil, err
		}
	}

	users, err := uas.userRepository.FindByUserIds(userIds)
	if err != nil {
		return nil, err
	}

	usersData = make([]UserData, len(users))
	for i, u := range users {
		usersData[i] = UserData{
			Id: *u.Id(),
			Name: *u.Name(),
		}
	}

	return usersData, nil
}