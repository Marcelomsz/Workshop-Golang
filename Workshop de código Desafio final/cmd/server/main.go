package main

import (
	"github.com/bootcamp-go/consignas-go-db.git/cmd/server/handler"
	"github.com/bootcamp-go/consignas-go-db.git/internal/dentist"
	"github.com/bootcamp-go/consignas-go-db.git/internal/patient"
	"github.com/bootcamp-go/consignas-go-db.git/internal/appointment"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/store"
	"github.com/gin-gonic/gin"
)

func main() {

	//Criation of storage, repository service and handle for 
	sqlStorage := store.NewSQLStore()
	repo := dentist.NewRepository(sqlStorage)
	service := dentist.NewService(repo)
	dentistHandler := handler.NewDentistHandler(service)

	//Criation of storage, repository service and handle for patient
	sqlStorageP := store.NewSQLStoreP()
	repoP := patient.NewRepository(sqlStorageP)
	serviceP := patient.NewService(repoP)
	patientHandler := handler.NewPatientHandler(serviceP)

	//Criation of storage, repository service and handle for appointment
	sqlStorageA := store.NewSQLStoreA()
	repoA := appointment.NewRepository(sqlStorageA)
	serviceA := appointment.NewService(repoA)
	appointmentHandler := handler.NewAppointmentHandler(serviceA)


	r := gin.Default()

	//creates a endpoint group "localhost:8080/dentists"
	r.GET("/ping", func(c *gin.Context) { c.String(200, "pong") })
	dentists := r.Group("/dentists")
	{
		// CRUD for dentist table
		dentists.GET(":id", dentistHandler.GetByID())
		dentists.POST("", dentistHandler.Post())
		dentists.DELETE(":id", dentistHandler.Delete())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.PUT(":id", dentistHandler.Put())
	}

	//creates a endpoint group "localhost:8080/dentists"
	patients := r.Group("/patients")
	{

		//consult by pacients'RG tha return  appointment, dentist informations and pacient informations.
		patients.GET("/RG/:RG", patientHandler.GetByRG())

		// CRUD for patient table
		patients.GET(":id", patientHandler.GetByID())
		patients.POST("", patientHandler.Post())
		patients.DELETE(":id", patientHandler.Delete())
		patients.PATCH(":id", patientHandler.Patch())
		patients.PUT(":id", patientHandler.Put())
	}

	//creates a endpoint group "localhost:8080/appointments"
	appointments := r.Group("/appointments")
	{
		// CRUD for appointment table
		appointments.GET(":id", appointmentHandler.GetByID())
		appointments.POST("", appointmentHandler.Post())
		appointments.DELETE(":id", appointmentHandler.Delete())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.PUT(":id", appointmentHandler.Put())
	}

	r.Run(":8080")
}

