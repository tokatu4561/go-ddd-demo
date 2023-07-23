package service

import "github.com/tokatu4561/go-ddd-demo/domain/model/circle"

type CircleService struct {
	circleRepository circle.CircleRepositoryInterface
}

func NewCircleService(circleRepository circle.CircleRepositoryInterface) (*CircleService, error) {
	return &CircleService{circleRepository: circleRepository}, nil
}

func (circleService *CircleService) Exists(circle *circle.Circle) (bool, error) {
	circle, err := circleService.circleRepository.FindByCircleName(&circle.Name)
	if err != nil {
		return false, err
	}
	if circle == nil {
		return false, nil
	}
	return true, nil
}