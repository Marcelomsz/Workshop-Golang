package patient

import (
	"errors"
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
)

type Repository interface {
	// Create a new item
	Create(p domain.Patient) (domain.Patient, error)
	// GetByID get a item by id
	GetByID(id int) (domain.Patient, error)
	GetByRG(RG int) (domain.Appointment,domain.Patient,domain.Dentist, error)	
	// Update a item by id
	Update(id int, p domain.Patient) (domain.Patient, error)
	// Delete a item by id
	Delete(id int) error
}

type repository struct {
	storage store.PatientInterface
}

// NewRepository
func NewRepository(storage store.PatientInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Patient, error) {
	patient, err := r.storage.GetOnePatient(id)
	if err != nil {
		return domain.Patient{}, errors.New("patient not found")
	}
	return patient, nil

}

func (r *repository) GetByRG(RG int) (domain.Appointment,domain.Patient,domain.Dentist, error) {
	appointment, patient, dentist, err := r.storage.GetByRG(RG)
	if err != nil {
		return domain.Appointment{},domain.Patient{},domain.Dentist{}, errors.New("Appointment not found")
	}
	return appointment, patient, dentist, nil

}

func (r *repository) Create(d domain.Patient) (domain.Patient, error) {
	//if !r.storage.Exists(d.Registration) {
	//	return domain.patient{}, errors.New("patient with this Registration already exists")}
	//err := r.storage.Create(d domain.patient)
	//if err != nil {
	//	return domain.patient{}, errors.New("error creating patient")
	//}
	return r.storage.CreatePatient(d)
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeletePatient(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, d domain.Patient) (domain.Patient, error) {
	//if !r.storage.Exists(d.Registration) {
	//	return domain.Dentist{}, errors.New("Registration already exists")
	//}
	//err := r.storage.Update(d)
	//if err != nil {
	//	return domain.Dentist{}, errors.New("error updating dentist")
	//}
	return r.storage.UpdatePatient(id,d)
}
