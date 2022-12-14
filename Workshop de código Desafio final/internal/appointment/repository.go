package appointment

import (
	"errors"
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
)

type Repository interface {
	// Create a new item
	Create(p domain.Appointment) (domain.Appointment, error)
	// GetByID get a item by id
	GetByID(id int) (domain.Appointment, error)
	// Update a item by id
	Update(id int, p domain.Appointment) (domain.Appointment, error)
	// Delete a item by id
	Delete(id int) error
}

type repository struct {
	storage store.AppointmentInterface
}

// NewRepository
func NewRepository(storage store.AppointmentInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Appointment, error) {
	appointment, err := r.storage.GetOneAppointment(id)
	if err != nil {
		return domain.Appointment{}, errors.New("appointment not found")
	}
	return appointment, nil

}


func (r *repository) Create(d domain.Appointment) (domain.Appointment, error) {
	//if !r.storage.Exists(d.Registration) {
	//	return domain.Appointment{}, errors.New("Appointment with this Registration already exists")}
	//err := r.storage.Create(d domain.Appointment)
	//if err != nil {
	//	return domain.Appointment{}, errors.New("error creating Appointment")
	//}
	return r.storage.CreateAppointment(d)
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeleteAppointment(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, d domain.Appointment) (domain.Appointment, error) {
	//if !r.storage.Exists(d.Registration) {
	//	return domain.Dentist{}, errors.New("Registration already exists")
	//}
	//err := r.storage.Update(d)
	//if err != nil {
	//	return domain.Dentist{}, errors.New("error updating dentist")
	//}
	return r.storage.UpdateAppointment(id,d)
}
