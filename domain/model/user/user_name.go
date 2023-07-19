package user

import (
	"fmt"
	"reflect"
)

type UserName struct {
	Value string
}

func NewUserName(name string) (_ *UserName, err error) {
	userName := new(UserName)
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if len(name) < 3 {
		return nil, fmt.Errorf("name %v is less than three characters long", name)
	}
	userName.Value = name
	return userName, nil
}

func (userName *UserName) Equals(other UserName) bool {
	return reflect.DeepEqual(userName.Value, other.Value)
}

func (userName *UserName) String() string {
	return fmt.Sprintf("UserName: [value: %s]", userName.Value)
}