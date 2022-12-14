package domain

import(
	//"time"
)

type Appointment struct {
	Id          	int 	    `json:"id"`
	IdPatient		int 	    `json:"id_pacient" binding:"required"`
	IdDentist       int			`json:"id_dentist" binding:"required"`
	DateTime        string   `json:"datetime" binding:"required"`
	Description		string      `json:"description" binding:"required"`
}

type AppointmentResponse struct {
	Id          	int 	    
	Patient		    Patient 	
	Dentist         Dentist		
	DateTime        string   	
	Description		string      
}
