package services

import (
	"chopitto-task/db"
	"golang.org/x/crypto/bcrypt"
)

// Create a new Person Service for interacting with the database
func NewPersonServices(person Person, personStore db.Store) *PersonService {
	return &PersonService{
		Person:      person,
		PersonStore: personStore,
	}
}

// Person Model
type Person struct {
	Id       int
	Email    string
	Password string
	Name     string
}

// Person Service
type PersonService struct {
	Person      Person
	PersonStore db.Store
}

// Create a new Person and INSERT into database
func (personService *PersonService) CreatePerson(person Person) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(person.Password), 8)
	if err != nil {
		return err
	}

	stmt := `INSERT into person(email, password, name) Values($1, $2, $3)`

	_, err = personService.PersonStore.Db.Exec(stmt, person.Email, hashedPassword, person.Name)
	return err
}

func (personService *PersonService) GetPerson(email string) (Person, error) {
	// Query for getting a person from an email
	query := `SELECT id, email, password, name FROM person WHERE email = ?`

	// Prepare the query
	stmt, err := personService.PersonStore.Db.Prepare(query)
	if err != nil {
		// return empty person and error
		return Person{}, err
	}

	// Defer the close of the query (Will close after function is complete)
	defer stmt.Close()

	// Query a single row (email is UNIQUE, so there will only be one row)
	err = stmt.QueryRow(email).Scan(
		// Put data into the person service person
		&personService.Person.Id,
		&personService.Person.Email,
		&personService.Person.Password,
		&personService.Person.Name,
	)
	if err != nil {
		// return empty person and error
		return Person{}, err
	}

	// Return person
	return personService.Person, nil

}

func (personService *PersonService) CheckEmail(email string) (Person, error) {

	query := `SELECT id, email, password, name FROM person
		WHERE email = ?`

	stmt, err := personService.PersonStore.Db.Prepare(query)
	if err != nil {
		return Person{}, err
	}

	defer stmt.Close()

	personService.Person.Email = email
	err = stmt.QueryRow(
		personService.Person.Email,
	).Scan(
		&personService.Person.Id,
		&personService.Person.Email,
		&personService.Person.Password,
		&personService.Person.Name,
	)
	if err != nil {
		return Person{}, err
	}

	return personService.Person, nil
}

func (personService *PersonService) GetAllPersons() ([]Person, error) {
	// Query for getting a person from an email
	query := `SELECT id, email, name FROM person`

	// Prepare the query
	rows, err := personService.PersonStore.Db.Query(query)
	if err != nil {
		// return empty person and error
		println("Get all persons Query error")
		return []Person{}, err
	}

	// Defer the close of the query (Will close after function is complete)
	defer rows.Close()

	persons := []Person{}

	for rows.Next() {
		err = rows.Scan(
			// Put data into the person service person
			&personService.Person.Id,
			&personService.Person.Email,
			&personService.Person.Name,
		)

		if err != nil {
			// return empty person and error
			println("Get All persons scan error")
			return []Person{}, err
		}

		persons = append(persons, personService.Person)

	}
	// Return person list
	return persons, nil

}
