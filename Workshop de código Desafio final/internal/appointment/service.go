package appointment

import (
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	//"github.com/bootcamp-go/consignas-go-db.git/internal/dentist"
	//"github.com/bootcamp-go/consignas-go-db.git/internal/patient"
)

type Service interface {
	// GetByID 
	GetByID(id int) (domain.Appointment, error)
	// Create 
	Create(p domain.Appointment) (domain.Appointment, error)
	// Delete 
	Delete(id int) error
	// Update 
	Update(id int, p domain.Appointment) (domain.Appointment, error)
}

type service struct {
	r Repository
}

// NewService 
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Appointment, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err
	}
	return p, nil
}

func (s *service) Create(d domain.Appointment) (domain.Appointment, error) {
	d, err := s.r.Create(d)
	if err != nil {
		return domain.Appointment{}, err
	}
	return d, nil
}

func (s *service) Update(id int, u domain.Appointment) (domain.Appointment, error) {
	d, err := s.r.GetByID(id)
	if err != nil {
		return domain.Appointment{}, err		
	}
	if u.IdPatient != 0 {	
	d.IdPatient = u.IdPatient
	}
	if u.IdDentist != 0 {
	d.IdDentist = u.IdDentist
	}
	if u.DateTime!= "" {
		d.DateTime = u.DateTime
	}
	if u.Description != "" {
		d.Description = u.Description
	}
	d, err = s.r.Update(id, d)
	if err != nil {
		return domain.Appointment{}, err
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
