package fullname

import (
	"fmt"
	"reflect"
	"regexp"
)

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName string, lastName string) (_ *FullName, err error) {
	fullName := new(FullName)

	// firstName
	if firstName == "" {
		return nil, fmt.Errorf("firstName is required")
	}
	if !ValidateName(firstName) {
		return nil, fmt.Errorf("許可されていない文字が使われています %s", firstName)
	}
	fullName.firstName = firstName

	// lastName
	if lastName == "" {
		return nil, fmt.Errorf("lastName is required")
	}
	if !ValidateName(lastName) {
		return nil, fmt.Errorf("許可されていない文字が使われています %s", lastName)
	}
	fullName.lastName = lastName

	return fullName, nil
}

func ValidateName(value string) bool {
	// アルファベットに限定
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(value)
}

func (fullName *FullName) Equals(otherFullName FullName) bool {
	return reflect.DeepEqual(fullName.firstName, otherFullName.firstName) && reflect.DeepEqual(fullName.lastName, otherFullName.lastName)
}
