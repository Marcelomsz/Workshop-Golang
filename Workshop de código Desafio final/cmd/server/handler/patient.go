package handler

import (
	"errors"
	"strconv"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/internal/patient"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

// NewProductHandler create a new controller for patient
func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

// Get  id
func (h *patientHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		patient, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		web.Success(c, 200, patient)
	}
}

func (h *patientHandler) GetByRG() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("RG")
		RG, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid RG"))
			return
		}
		appointment, patient, dentist, err := h.s.GetByRG(RG)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}		

		//result := domain.ResultRG(Appointment,)
		//result.appointment = appointment
		web.Success(c, 200, appointment )
		web.Success(c, 200, patient )
		web.Success(c, 200, dentist )
	}
}

// validateEmptys
func validateEmptysP(patient *domain.Patient) (bool, error) {
	if patient.FirstName == "" || patient.LastName == "" || patient.RG == 0{
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post create new patient
func (h *patientHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var patient domain.Patient
		err := c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysP(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(patient)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// Delete
func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(c, 404, err)
			return
		}
		web.Success(c, 204, nil)
	}
}

// Put
func (h *patientHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var patient domain.Patient
		err = c.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysP(&patient)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch update a patient or some of they fields
func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Id          	 int    
		FirstName        string 
		LastName         string 
		RG			     int   
		RegistrationDate string
	}

	return func(c *gin.Context) {
		var r Request
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("patient not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Patient{
			FirstName:     		r.FirstName,
			LastName:      		r.LastName,
			RG:  		   		r.RG,
			RegistrationDate:   r.RegistrationDate,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
