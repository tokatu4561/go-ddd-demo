package circle

import (
	"fmt"
	"reflect"
)

type CircleName struct {
	Value string
}

func NewCircleName(name string) (_ *CircleName, err error) {
	CircleName := new(CircleName)
	if name == "" {
		return nil, fmt.Errorf("name is required")
	}
	if len(name) < 3 {
		return nil, fmt.Errorf("name %v is less than three characters long", name)
	}
	CircleName.Value = name
	return CircleName, nil
}

func (CircleName *CircleName) Equals(other CircleName) bool {
	return reflect.DeepEqual(CircleName.Value, other.Value)
}

func (CircleName *CircleName) String() string {
	return fmt.Sprintf("CircleName: [value: %s]", CircleName.Value)
}