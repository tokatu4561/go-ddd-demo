package circle

type CircleRepositoryInterface interface {
	FindByCircleName(name *CircleName) (*Circle, error)
	Find(id *CircleId) (*Circle, error)
	Save(circle *Circle) error
}