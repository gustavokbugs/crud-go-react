package services

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/application/service_interface"
	"compartilhatech/internal/domain/entities"
	"compartilhatech/internal/infra/database/sqlc/queries"
	"context"
	"database/sql"
	"fmt"
)

type PersonService struct {
	db               *sql.DB
	PersonRepository []entities.Person
}

func NewPersonService(db *sql.DB) service_interface.PersonService {
	return &PersonService{
		db:               db,
		PersonRepository: []entities.Person{},
	}
}

func (s *PersonService) Insert(data dto.CreatePerson) (*entities.Person, error) {
	fmt.Println("Insert called with data:", data)
	p := entities.NewPerson()
	p.Name = data.Name
	p.Age = data.Age

	if data.Active != nil {
		p.Active = *data.Active
	}

	dbConn := queries.New(s.db)
	err := dbConn.InsertPerson(context.Background(), queries.InsertPersonParams{
		ID:        p.ID,
		Name:      p.Name,
		Age:       sql.NullInt32{Int32: int32(p.Age), Valid: true},
		Active:    p.Active,
		CreatedAt: p.CreatedAt,
		UpdatedAt: p.UpdatedAt,
	})
	if err != nil {
		fmt.Println("Error inserting person into database:", err)
		return nil, err
	}

	fmt.Println("Person inserted successfully:", p)
	return p, nil
}

func (s *PersonService) List() ([]entities.Person, error) {
	fmt.Println("List called")
	dbConn := queries.New(s.db)

	p, err := dbConn.GetPersons(context.Background())
	if err != nil {
		fmt.Println("Error fetching persons from database:", err)
		return nil, err
	}

	persons := []entities.Person{}
	for _, v := range p {
		persons = append(persons, entities.Person{
			ID:        v.ID,
			Name:      v.Name,
			Age:       int(v.Age.Int32),
			Active:    v.Active,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	fmt.Println("Persons fetched successfully:", persons)
	return persons, nil
}

func (s *PersonService) GetById(ID string) (*entities.Person, error) {
	fmt.Println("GetById called with ID:", ID)
	dbConn := queries.New(s.db)

	personData, err := dbConn.GetPersonById(context.Background(), ID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Person not found in database for ID:", ID)
			return nil, nil
		}
		fmt.Println("Error fetching person from database:", err)
		return nil, err
	}

	person := &entities.Person{
		ID:        personData.ID,
		Name:      personData.Name,
		Age:       int(personData.Age.Int32),
		Active:    personData.Active,
		CreatedAt: personData.CreatedAt,
		UpdatedAt: personData.UpdatedAt,
	}

	fmt.Println("Person fetched successfully:", person)
	return person, nil
}

func (s *PersonService) Update(ID string, data dto.UpdatePerson) (*entities.Person, error) {
	fmt.Println("Update called with ID:", ID, "and data:", data)
	dbConn := queries.New(s.db)

	currentPerson, err := s.GetById(ID)
	if err != nil {
		fmt.Println("Error fetching person for update:", err)
		return nil, err
	}
	if currentPerson == nil {
		fmt.Println("Person not found for ID:", ID)
		return nil, fmt.Errorf("person not found")
	}

	if data.Name != nil {
		currentPerson.Name = *data.Name
	}
	if data.Age != nil {
		currentPerson.Age = *data.Age
	}
	if data.Active != nil {
		currentPerson.Active = *data.Active
	}

	err = dbConn.UpdatePerson(context.Background(), queries.UpdatePersonParams{
		ID:        currentPerson.ID,
		Name:      sql.NullString{String: currentPerson.Name, Valid: true},
		Age:       sql.NullInt32{Int32: int32(currentPerson.Age), Valid: true},
		Active:    sql.NullBool{Bool: currentPerson.Active, Valid: true},
		UpdatedAt: currentPerson.UpdatedAt,     
	})
	if err != nil {
		fmt.Println("Error updating person in database:", err)
		return nil, err
	}

	fmt.Println("Person updated successfully:", currentPerson)
	return currentPerson, nil
}


func (s *PersonService) Delete(ID string) error {
	fmt.Println("Delete called with ID:", ID)
	dbConn := queries.New(s.db)

	err := dbConn.DeletePerson(context.Background(), ID)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("Person not found for deletion:", ID)
			return fmt.Errorf("not found")
		}
		fmt.Println("Error deleting person from database:", err)
		return err
	}

	fmt.Println("Person deleted successfully for ID:", ID)
	return nil
}
