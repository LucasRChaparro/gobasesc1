package productos

type Service interface {
	GetAll() Product
}
type service struct {
	repo Repository
}

func NewService(s Repository) Service {
	return &service{
		repo: s,
	}
}
func (s *service) GetAll() Product {
	ps := s.repo.GetAll()
	return ps
}
