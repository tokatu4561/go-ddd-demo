package circle

import "github.com/tokatu4561/go-ddd-demo/domain/model/user"

type Circle struct {
	Id      CircleId
	Name    CircleName
	Owner   user.User
	Members []user.UserId
}

type CircleIsFullError struct {
	CircleId CircleId
	Message  string
}

type MemberIsNotFoundError struct {
	MemberId user.UserId
	Message  string
}


func (minfe *MemberIsNotFoundError) Error() string {
	return minfe.Message
}

func (cife *CircleIsFullError) Error() string {
	return cife.Message
}

func NewCircle(id CircleId, name CircleName, owner user.User, members []user.UserId) (*Circle, error) {
	return &Circle{Id: id, Name: name, Owner: owner, Members: members}, nil
}

func (circle *Circle) CountMembers() int {
	return len(circle.Members) + 1 // オーナーも含める
}

func (circle *Circle) IsFull() bool {
	return circle.CountMembers() >= 29
}

func (circle *Circle) Join(newMember *user.User) error {
	if circle.IsFull() {
		return &CircleIsFullError{CircleId: circle.Id, Message: "circle is full"}
	}
	circle.Members = append(circle.Members, *newMember.Id())
	return nil
}

func (circle *Circle) ChangeName(name CircleName) error {
	circle.Name = name
	return nil
}

func (circle *Circle) ChangeMemberName(memberId *user.UserId, changedUserName *user.UserName) error {
	for i, member := range circle.Members {
		if member.Equals(memberId) {
			circle.Members[i].ChangeUserName(*changedUserName)
			return nil
		}
	}
	return &MemberIsNotFoundError{MemberId: *memberId, Message: "member is not found"}
}
