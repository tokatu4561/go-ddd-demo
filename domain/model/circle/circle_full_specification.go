package circle

import (
	"github.com/tokatu4561/go-ddd-demo/domain/model/user"
)

type CircleFullSpecification struct {
	userRepository user.UserRepositoryInterface
}

func NewCircleFullSpecification(userRepository user.UserRepositoryInterface) (*CircleFullSpecification, error) {
	return &CircleFullSpecification{ userRepository: userRepository}, nil
}

func (cfs *CircleFullSpecification) IsSatisfiedBy(circle *Circle) (bool, error) {
	// サークルに所属するユーザーの最大数はサークルのオーナーとなるユーザーを含めて30人までとする
	// プレミアムユーザーが１０名以上所属しているサークルはメンバーの最大数が５０名になる
	users, err := cfs.userRepository.FindByUserIds(circle.Members)
	if err != nil {
		return false, err
	}
	// プレミアムユーザーの数
	premiumUserCount := 0
	for _, user := range users {
		if user.IsPremium() {
			premiumUserCount++
		}
	}
	
	overtimeUserCount := 30
	if premiumUserCount >= 10 {
		overtimeUserCount = 50
	}

	return circle.CountMembers() >= overtimeUserCount, nil
}