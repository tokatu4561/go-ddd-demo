package application

import (
	"context"
	"fmt"

	"github.com/tokatu4561/go-ddd-demo/domain/model/user"
	"github.com/tokatu4561/go-ddd-demo/infrastructure/transaction"
)

type UserApplicationService struct {
	userRepository user.UserRepositoryInterface
	userService    user.UserService
	userFactory   user.UserFactory
	transaction   transaction.Transaction
}

type GetUserListCommand struct {
	UserIds []string
}

type UserData struct {
	Id      user.UserId
	Name    user.UserName
}

func (uas *UserApplicationService) GetUsers(command GetUserListCommand) (usersData []UserData, err error) {
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

type UserRegisterCommand struct {
	Name string
}

// register user
func (uas *UserApplicationService) Register(command UserRegisterCommand, ctx context.Context) error {
	userName, err := user.NewUserName(command.Name)
	if err != nil {
		return nil
	}
	user, err := uas.userFactory.Create(*userName)

	isExists, err := uas.userService.Exists(user)
	if err != nil {
		return err
	}
	if isExists {
		return fmt.Errorf("user of %s is already exists.", command.Name)
	}
	
	// user の登録は重複した名前が存在してはいけないというルールがあった場合、トランザクションレベルで排他制御も必要
	err = uas.transaction.DoInTx(ctx, func(ctx context.Context) error {
		if err = uas.userRepository.Save(user, ctx); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}