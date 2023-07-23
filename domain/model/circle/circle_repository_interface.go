package circle

type CircleRepositoryInterface interface {
	FindByCircleName(name *CircleName) (*Circle, error)
	Save(circle *Circle) error
}