package repository

import (
	"database/sql"
	"fmt"

	domain_user "github.com/tokatu4561/go-ddd-demo/domain/model/user"
)


type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) (*domain_user.UserRepositoryInterface, error) {
	return &UserRepository{db: db}, nil
}

func (ur *UserRepository) FindByUserName(name *domain_user.UserName) (user *domain_user.User, err error) {
	tx, err := ur.db.Begin()
	if err != nil {
		return
	}
	defer func() {
		switch err {
		case nil:
			err = tx.Commit()
		default:
			tx.Rollback()
		}
	}()

	rows, err := tx.Query("SELECT id, name FROM users WHERE name = $1", name.Value)
	if err != nil {
		return nil, &FindByUserNameQueryError{UserName: *name, Message: fmt.Sprintf("error is occured in userrepository.FindByUserName: %s", err), Err: err}
	}
	defer rows.Close()

	userId := &domain_user.UserId{}
	userName := &domain_user.UserName{}
	for rows.Next() {
		err := rows.Scan(&userId.value, &userName.value)
		if err != nil {
			return nil, err
		}
		user = &domain_user.User{id: *userId, name: *userName}
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return user, nil
}

type FindByUserNameQueryError struct {
	UserName domain_user.UserName
	Message  string
	Err      error
}

func (err *FindByUserNameQueryError) Error() string {
	return err.Message
}

func (ur *UserRepository) Save(user *domain_user.User) (err error) {
	_, err = ur.db.Exec("INSERT INTO users(id, name) VALUES ($1, $2)", user.Id().Id, user.Name().Value)
	if err != nil {
		return &SaveQueryRowError{UserName: *user.Name(), Message: fmt.Sprintf("userrepository.Save err: %s", err), Err: err}
	}
	return nil
}

type SaveQueryRowError struct {
	UserName domain_user.UserName
	Message  string
	Err      error
}

func (err *SaveQueryRowError) Error() string {
	return err.Message
}
