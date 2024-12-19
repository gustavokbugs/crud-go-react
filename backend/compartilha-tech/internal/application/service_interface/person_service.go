package service_interface

import (
	"compartilhatech/internal/application/dto"
	"compartilhatech/internal/domain/entities"
)

type PersonService interface {
	Insert(data dto.CreatePerson) (*entities.Person, error)
	List() ([]entities.Person, error)
	GetById(ID string) (*entities.Person, error)
	Update(ID string, data dto.UpdatePerson) (*entities.Person, error)
	Delete(ID string) error
}
