package handler

import (
	"errors"
	"strconv"

	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"github.com/bootcamp-go/consignas-go-db.git/internal/dentist"
	"github.com/bootcamp-go/consignas-go-db.git/pkg/web"
	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

// NewProductHandler
func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

// Get by id
func (h *dentistHandler) GetByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		dentist, err := h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		web.Success(c, 200, dentist)
	}
}

// validateEmptys
func validateEmptys(dentist *domain.Dentist) (bool, error) {
	if dentist.FirstName == "" || dentist.LastName == "" || dentist.Registration == ""{
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

// Post
func (h *dentistHandler) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var dentist domain.Dentist
		err := c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(c, 400, err)
			return
		}
		web.Success(c, 201, d)
	}
}

// Delete
func (h *dentistHandler) Delete() gin.HandlerFunc {
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
func (h *dentistHandler) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid id"))
			return
		}
		_, err = h.s.GetByID(id)
		if err != nil {
			web.Failure(c, 404, errors.New("dentist not found"))
			return
		}
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		var dentist domain.Dentist
		err = c.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptys(&dentist)
		if !valid {
			web.Failure(c, 400, err)
			return
		}
		d, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, d)
	}
}

// Patch 
func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		FirstName       string  `json:"firstname,omitempty"`
		LastName		string	`json:"lastname,omitempty"`
		Registration    string  `json:"registration,omitempty"`
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
			web.Failure(c, 404, errors.New("product not found"))
			return
		}
		if err := c.ShouldBindJSON(&r); err != nil {
			web.Failure(c, 400, errors.New("invalid json"))
			return
		}
		update := domain.Dentist{
			FirstName:     r.FirstName,
			LastName:      r.LastName,
			Registration:  r.Registration,
		}
		p, err := h.s.Update(id, update)
		if err != nil {
			web.Failure(c, 409, err)
			return
		}
		web.Success(c, 200, p)
	}

}
