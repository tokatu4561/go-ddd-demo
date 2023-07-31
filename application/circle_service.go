package application

import (
	"fmt"
	"log"

	"github.com/tokatu4561/go-ddd-demo/domain/model/circle"
	"github.com/tokatu4561/go-ddd-demo/domain/model/user"
	"github.com/tokatu4561/go-ddd-demo/domain/service"
	"github.com/tokatu4561/go-ddd-demo/infrastructure/sql/services/circles"
)

type CircleApplicationService struct {
	circleRepository circle.CircleRepositoryInterface
	circleService    service.CircleService
	userRepository  user.UserRepositoryInterface
}

func NewCircleApplicationService(circleRepository circle.CircleRepositoryInterface, circleService service.CircleService) (*CircleApplicationService, error) {
	return &CircleApplicationService{circleRepository: circleRepository, circleService: circleService}, nil
}

type CircleJoinCommand struct {
	CircleId string
	UserId   string
}

func (cas *CircleApplicationService) Join(command CircleJoinCommand) error {
	memberId, err := user.NewUserId(command.UserId)
	member, err := cas.userRepository.Find(memberId)
	if err != nil {
		return err
	}
	if member == nil {
		return fmt.Errorf("user of %s is not found.", command.UserId)
	}

	circleId, err := circle.NewCircleId(command.CircleId)
	findCircle, err := cas.circleRepository.Find(circleId)
	if err != nil {
		return err
	}
	if findCircle == nil {
		return fmt.Errorf("circle of %s is not found.", command.CircleId)
	}

	circleFullSpecification, err:= circle.NewCircleFullSpecification(cas.userRepository)
	if err != nil {
		return err
	}
	if isSatisfied, err := circleFullSpecification.IsSatisfiedBy(findCircle); err != nil {
		return err
	} else if isSatisfied {
		return fmt.Errorf("circle of %s is full.", command.CircleId)
	}

	if err := findCircle.Join(member); err != nil {
		return err
	}
	if err := cas.circleRepository.Save(findCircle); err != nil {
		return err
	}

	// transaction.Complete();
	return nil
}

type CircleUpdateCommand struct {
	Id   string
	Name string
}

func (cas *CircleApplicationService) Update(command CircleUpdateCommand) (err error) {
	circleId, err := circle.NewCircleId(command.Id)
	if err != nil {
		return err
	}
	targetCircle, err := cas.circleRepository.Find(circleId)
	if err != nil {
		return err
	}
	if targetCircle == nil {
		return fmt.Errorf("circle of %s is not found.", command.Id)
	}

	if command.Name != "" {
		circleName, err := circle.NewCircleName(command.Name)
		if err != nil {
			return err
		}
		_ = targetCircle.ChangeName(*circleName)
		if isCircleExists, err := cas.circleService.Exists(targetCircle); err != nil {
			return err
		} else if isCircleExists {
			return fmt.Errorf("circleName of %s is already exists.", command.Name)
		}
	}

	if err := cas.circleRepository.Save(targetCircle); err != nil {
		return err
	}
	return nil
}

func (cas *CircleApplicationService) Create(circleName string) (err error) {
	newCircleId, err := circle.NewCircleId("test-circle-id")
	if err != nil {
		return nil
	}
	newCircleName, err := circle.NewCircleName(circleName)
	if err != nil {
		return nil
	}

	ownerId, err := user.NewUserId("ownerId")
	if err != nil {
		return nil
	}
	ownerName, err := user.NewUserName("ownerName")
	if err != nil {
		return nil
	}
	owner, err := user.NewUser(*ownerId, *ownerName)
	if err != nil {
		return nil
	}

	memberId, err := user.NewUserId("memberId")
	if err != nil {
		return nil
	}
	memberName, err := user.NewUserName("memberName")
	if err != nil {
		return nil
	}
	member, err := user.NewUser(*memberId, *memberName)
	if err != nil {
		return nil
	}

	members := []user.User{*owner, *member}
	newCircle, err := circle.NewCircle(*newCircleId, *newCircleName, *owner, members)
	if err != nil {
		return nil
	}
	isCircleExists, err := cas.circleService.Exists(newCircle)
	if err != nil {
		return err
	}
	if isCircleExists {
		return fmt.Errorf("circleName of %s is already exists.", circleName)
	}

	if err := cas.circleRepository.Save(newCircle); err != nil {
		return err
	}
	log.Println("success fully saved")
	return nil
}

type RegisterError struct {
	Name    string
	Message string
	Err     error
}

func (err *RegisterError) Error() string {
	return err.Message
}

type CircleData struct {
	Id      circle.CircleId
	Name    circle.CircleName
	Owner   user.User
	Members []user.User
}

func (cas *CircleApplicationService) Get(circleName string) (_ *CircleData, err error) {
	targetName, err := circle.NewCircleName(circleName)
	if err != nil {
		return nil, err
	}
	circle, err := cas.circleRepository.FindByCircleName(targetName)
	if err != nil {
		return nil, err
	}
	return &CircleData{Id: circle.Id, Name: circle.Name, Owner: circle.Owner, Members: circle.Members}, nil
}


// ページネーションなど複雑なクエリはクエリサービスを作成する
// なぜならリポジトリを利用する場合は、全件取得してから指定件数絞り込みが必要になってしまうため　全件取得が無駄
func (cas *CircleApplicationService) GetSummaries(command circles.CircleGetSummaryCommand) (_ []*circles.CircleSummaryData, err error) {
	circleQueryService, err := circles.NewCircleQueryService()
	if err != nil {	
		return nil, err
	}
	circles, err := circleQueryService.GetSummaries(&command)

	return circles, nil
}
