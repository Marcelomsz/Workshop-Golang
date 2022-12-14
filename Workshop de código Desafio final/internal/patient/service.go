package patient

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
)

type Service interface {
	// GetByID 
	GetByID(id int) (domain.Patient, error)
	GetByRG(RG int) (domain.Appointment,domain.Patient,domain.Dentist, error)
	// Create 
	Create(p domain.Patient) (domain.Patient, error)
	// Delete 
	Delete(id int) error
	// Update 
	Update(id int, p domain.Patient) (domain.Patient, error)
}

type service struct {
	r Repository
}

// NewService 
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Patient, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	return p, nil
}

func (s *service) GetByRG(RG int) (domain.Appointment,domain.Patient,domain.Dentist, error) {
	appointment, patient, dentist, err := s.r.GetByRG(RG)
	if err != nil {
		return domain.Appointment{},domain.Patient{},domain.Dentist{}, err
	}
	return appointment, patient, dentist, nil

}

func (s *service) Create(d domain.Patient) (domain.Patient, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Patient{}, err
	}
	return d, nil
}
func (s *service) Update(id int, u domain.Patient) (domain.Patient, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Patient{}, err
	}
	if u.FirstName != "" {
		d.FirstName = u.FirstName
	}
	if u.LastName != "" {
		d.LastName = u.LastName
	}
	if u.RG != 0 {
		d.RG = u.RG
	}
	if u.RegistrationDate != "" {
		d.RegistrationDate = u.RegistrationDate
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Patient{}, err
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
