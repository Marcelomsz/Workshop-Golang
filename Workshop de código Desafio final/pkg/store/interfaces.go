package store

import "github.com/bootcamp-go/consignas-go-db.git/internal/domain"

type DentistInterface interface {
	// Read 
	GetOneDentist(id int) (domain.Dentist, error)
	// Create 
	CreateDentist(dentist domain.Dentist) (domain.Dentist, error)
	// Update 
	UpdateDentist(id int, d domain.Dentist) (domain.Dentist, error)
	// Delete 
	DeleteDentist(id int) error	
}

type PatientInterface interface {
	// Read 
	GetOnePatient(id int) (domain.Patient, error)
	// Create 
	CreatePatient(patient domain.Patient) (domain.Patient, error)
	// Update 
	UpdatePatient(id int, d domain.Patient) (domain.Patient, error)
	// Delete 
	DeletePatient(id int) error	
	GetByRG(RG int) (domain.Appointment, domain.Patient ,domain.Dentist, error)
}

type AppointmentInterface interface {
	// Read 
	GetOneAppointment(id int) (domain.Appointment,error)
	// Create 
	CreateAppointment(appointment domain.Appointment) (domain.Appointment, error)
	// Update 
	UpdateAppointment(id int, d domain.Appointment) (domain.Appointment, error)
	// Delete 
	DeleteAppointment(id int) error	
}


