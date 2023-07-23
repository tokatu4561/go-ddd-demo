package circle

import "fmt"

type CircleId struct {
	Value string
}

func NewCircleId(id string) (_ *CircleId, err error) {
	circleId := new(CircleId)
	if id == "" {
		return nil, fmt.Errorf("userId is required")
	}
	circleId.Value = id
	return circleId, nil
}
