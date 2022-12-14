package store

import (
	"log"
	"github.com/bootcamp-go/consignas-go-db.git/internal/domain"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type sqlStore struct {
	db *sql.DB
}

//sql conection for Dentist
func NewSQLStore() DentistInterface {
	database, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/my_db")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{
		db: database,
	}
}

//sql conection for patient
func NewSQLStoreP() PatientInterface {
	database, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/my_db")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{
		db: database,
	}
}

//sql conection for Appointment
func NewSQLStoreA() AppointmentInterface {
	database, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/my_db")
	if err != nil {
		log.Fatalln(err)
		panic(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlStore{
		db: database,
	}
}

//Crud for Dentist
func (s *sqlStore) CreateDentist(d domain.Dentist) (domain.Dentist,error) {
	result, err := s.db.Exec("INSERT INTO dentists (FirstName, LastName, Registration) VALUES (?,?,?)",d.FirstName, d.LastName,d.Registration)
	if err != nil {
		log.Fatal(err)
		}	
	if err != nil {
	return domain.Dentist{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Dentist{}, err
	}
	d.Id = int(lastInsertId)
	return d, nil
}

func (s *sqlStore) GetOneDentist(id int) (domain.Dentist ,error) {
	var d domain.Dentist
	rows, err := s.db.Query("SELECT * FROM dentists WHERE id=?", id)
	if err != nil {
		return d, err
	}
	for rows.Next() {
		if err := rows.Scan(&d.Id,
			&d.FirstName,
			&d.LastName,
			&d.Registration);err != nil {
		log.Println(err.Error())
		return d, err
		}
	}
	return d, nil
}


func (s *sqlStore) UpdateDentist(id int, d domain.Dentist) (domain.Dentist,error) {
	stmt, err := s.db.Prepare("UPDATE dentists SET firstname = ?, lastname = ?, registration =? WHERE id=?")
	if err != nil {
	log.Fatal(err)
	}
	defer stmt.Close() 
	_, err = stmt.Exec(d.FirstName,d.LastName,d.Registration,id)
	if err != nil {
	return domain.Dentist{}, err
	}
	return d,nil
	}

func (s *sqlStore) DeleteDentist(id int) error {
	_, err := s.db.Exec("DELETE FROM dentists WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}


//Crud for Patient
func (s *sqlStore) CreatePatient(p domain.Patient) (domain.Patient,error) {
	result, err := s.db.Exec("INSERT INTO patients (firstname, lastname, rg,registrationDate) VALUES (?,?,?,?)",p.FirstName, p.LastName,p.RG,p.RegistrationDate)
	if err != nil {
		log.Fatal(err)
		}	
	if err != nil {
	return domain.Patient{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Patient{}, err
	}
	p.Id = int(lastInsertId)
	return p, nil
}

func (s *sqlStore) GetOnePatient(id int) (domain.Patient ,error) {
	var d domain.Patient
	rows, err := s.db.Query("SELECT * FROM patients WHERE id=?", id)
	if err != nil {
		return d, err
	}
	for rows.Next() {
		if err := rows.Scan(&d.Id,
			&d.FirstName,
			&d.LastName,
			&d.RG,
			&d.RegistrationDate);err != nil {
		log.Println(err.Error())
		return d, err
		}
	}
	return d, nil
}

func (s *sqlStore) GetByRG(RG int) (domain.Appointment , domain.Patient,domain.Dentist,  error) {
	var d domain.Appointment
	var p domain.Patient
	var k domain.Dentist
	rows, err := s.db.Query("SELECT * FROM patients WHERE rg=?", RG)
	if err != nil {
		return d ,p ,k , err
	}
	for rows.Next() {
		if err := rows.Scan(&p.Id,
			&p.FirstName,
			&p.LastName,
			&p.RG,
			&p.RegistrationDate);err != nil {
		log.Println(err.Error())
		return d ,p ,k , err
		}
	}
	rowsA, err := s.db.Query("SELECT * FROM appointments WHERE id_pacient=?", p.Id)
	if err != nil {
		return d ,p ,k , err
	}
	for rowsA.Next() {
		if err := rowsA.Scan(
			&d.Id,
			&d.IdPatient,
			&d.IdDentist,
			&d.DateTime,
			&d.Description);err != nil {
		log.Println(err.Error())
		return d ,p ,k , err
		}
	}
	rowsD, err := s.db.Query("SELECT * FROM dentists WHERE id=?", d.IdDentist)
	if err != nil {
		return d ,p ,k , err
	}
	for rowsD.Next() {
		if err := rowsD.Scan(
			&k.Id,
			&k.FirstName,
			&k.LastName,
			&k.Registration);err != nil {
		log.Println(err.Error())
		return d ,p ,k , err
		}
	}
	return d ,p ,k , nil

}

func (s *sqlStore) UpdatePatient(id int, d domain.Patient) (domain.Patient,error) {
	stmt, err := s.db.Prepare("UPDATE patients SET firstname = ?, lastname = ?, Rg =? ,RegistrationDate =? WHERE id=?")
	if err != nil {
	log.Fatal(err)
	}
	defer stmt.Close() 
	_, err = stmt.Exec(d.FirstName,d.LastName,d.RG,d.RegistrationDate,id)
	if err != nil {
	return domain.Patient{}, err
	}
	return d,nil
	}

func (s *sqlStore) DeletePatient(id int) error {
	_, err := s.db.Exec("DELETE FROM patients WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}


//Crud for Appointment
func (s *sqlStore) CreateAppointment(p domain.Appointment) (domain.Appointment,error) {
	result, err := s.db.Exec("INSERT INTO appointments (id_pacient, id_dentist, datetime,description) VALUES (?,?,?,?)",p.IdPatient, p.IdDentist,p.DateTime,p.Description)
	if err != nil {
		log.Fatal(err)
		}	
	if err != nil {
	return domain.Appointment{}, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println(err.Error())
		return domain.Appointment{}, err
	}
	p.Id = int(lastInsertId)
	return p, nil
}

func (s *sqlStore) GetOneAppointment(id int) (domain.Appointment ,error) {
	var d domain.Appointment
	rows, err := s.db.Query("SELECT * FROM appointments WHERE id=?", id)
	if err != nil {
		return d, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&d.Id,
			&d.IdPatient,
			&d.IdDentist,
			&d.DateTime,
			&d.Description);err != nil {
		log.Println(err.Error())
		return d, err
		}
	}
	return d, nil
}

func (s *sqlStore) UpdateAppointment(id int, d domain.Appointment) (domain.Appointment,error) {
	stmt, err := s.db.Prepare("UPDATE appointments SET id_pacient= ?, id_dentist= ?, datetime =? ,description =? WHERE id=?")
	if err != nil {
	log.Fatal(err)
	}
	defer stmt.Close() 
	_, err = stmt.Exec(d.IdPatient, d.IdDentist,d.DateTime,d.Description,id)
	if err != nil {
	return domain.Appointment{}, err
	}
	return d,nil
	}

func (s *sqlStore) DeleteAppointment(id int) error {
	_, err := s.db.Exec("DELETE FROM appointments WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}