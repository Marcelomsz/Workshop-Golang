package dentist

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type Service interface {
	// GetByID 
	GetByID(id int) (domain.Dentist, error)
	// Create 
	Create(d domain.Dentist) (domain.Dentist, error)
	// Delete 
	Delete(id int) error
	// Update 
	Update(id int, d domain.Dentist) (domain.Dentist, error)
}

type service struct {
	r Repository
}

// NewService 
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Dentist, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	return p, nil
}

func (s *service) Create(d domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}
func (s *service) Update(id int, u domain.Dentist) (domain.Dentist, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentist{}, err
	}
	if u.FirstName != "" {
		d.FirstName = u.FirstName
	}
	if u.LastName != "" {
		d.LastName = u.LastName
	}
	if u.Registration != "" {
		d.Registration = u.Registration
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Dentist{}, err
	}
	return d, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
