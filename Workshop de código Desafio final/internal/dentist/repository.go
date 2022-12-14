package dentist

import (
	"errors"
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
)

type Repository interface {
	// Create a new item
	Create(d domain.Dentist) (domain.Dentist, error)
	// GetByID get a item by id
	GetByID(id int) (domain.Dentist, error)
	// Update a item by id
	Update(id int, d domain.Dentist) (domain.Dentist, error)
	// Delete a item by id
	Delete(id int) error
}

type repository struct {
	storage store.DentistInterface
}

// NewRepository
func NewRepository(storage store.DentistInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Dentist, error) {
	dentist, err := r.storage.GetOneDentist(id)
	if err != nil {
		return domain.Dentist{}, errors.New("dentist not found")
	}
	return dentist, nil

}

func (r *repository) Create(d domain.Dentist) (domain.Dentist, error) {
	//if !r.storage.Exists(d.Registration) {
	//	return domain.Dentist{}, errors.New("Dentist with this Registration already exists")}
	//err := r.storage.Create(d domain.Dentist)
	//if err != nil {
	//	return domain.Dentist{}, errors.New("error creating dentist")
	//}
	return r.storage.CreateDentist(d)
}

func (r *repository) Delete(id int) error {
	err := r.storage.DeleteDentist(id) 
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, d domain.Dentist) (domain.Dentist, error) {
	//if !r.storage.Exists(d.Registration) {
	//	return domain.Dentist{}, errors.New("Registration already exists")
	//}
	//err := r.storage.Update(d)
	//if err != nil {
	//	return domain.Dentist{}, errors.New("error updating dentist")
	//}
	return r.storage.UpdateDentist(id,d)
}
