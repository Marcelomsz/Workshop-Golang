package handler

import (
	"errors"
	"strconv"
	//"time"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/internal/appointment"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

// NewProductHandler create a new controller for appointment
func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

// Get a appointment by id
func (h *appointmentHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		appointment, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		web.Success(c, 200, appointment)
	}
}


// validateEmptys validate empty fields
func validateEmptysA(appointment *domain.Appointment) (bool, error) {
	if appointment.IdPatient == 0 || appointment.IdDentist == 0{ 
	//|| appointment.DateTime.Format("0000-00-00 00:00:00") == ""
	
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post create new appointment
func (h *appointmentHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var appointment domain.Appointment
		err := c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysA(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(appointment)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// Delete appointment
func (h *appointmentHandler) Delete() gin.HandlerFunc {
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

// Put update appointment
func (h *appointmentHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var appointment domain.Appointment
		err = c.ShouldBindJSON(&appointment)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysA(&appointment)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, appointment)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch update a appointment or some of they fields
func (h *appointmentHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Id             int  
		IdPatient      int 
		IdDentist      int 
		DateTime	   string   
		Description    string
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
			web.Failure(c, 404, errors.New("appointment not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Appointment{
			IdPatient:     r.IdPatient,
			IdDentist:     r.IdDentist,
			DateTime:  	   r.DateTime,
			Description:   r.Description,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}
}
