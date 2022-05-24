package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/Mynor2397/sqlnulls"
	"github.com/google/uuid"
)

type personService struct {
}

var PersonStorage storage.PersonStorage

// NewPersonService retorna un nuevo servicio para los usuarios
func NewPersonService(personStorage storage.PersonStorage) PersonService {
	PersonStorage = personStorage
	return &personService{}
}

// PersonService implementa el conjunto de metodos de servicio para usuario
type PersonService interface {
	Create(ctx context.Context, person models.Person) (models.Person, error)
	GetOne(ctx context.Context, uuid string) (models.Person, error)
	GetMany(ctx context.Context, filter string, offset, limit int) ([]models.Person, error)
	GetManyWithFullInformation(ctx context.Context, filter string, offset, limit int) ([]models.Person, error)

	Update(ctx context.Context, uuid string, person models.Person) (string, error)

	CreateSubstitute(ctx context.Context, person models.Person) (models.Person, error)
	GetOneSubstitute(ctx context.Context, uuidPerson string) (models.Person, error)
	GetSubstitutes(ctx context.Context) ([]models.Person, error)

	GetNamePerson(ctx context.Context) ([]models.Person, error)
}

func (*personService) Create(ctx context.Context, person models.Person) (models.Person, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &person)
	return PersonStorage.Create(ctx, person)
}

func (*personService) GetOne(ctx context.Context, uuid string) (models.Person, error) {
	return PersonStorage.GetOne(ctx, uuid)
}

func (*personService) GetMany(ctx context.Context, filter string, offset, limit int) ([]models.Person, error) {
	return PersonStorage.PaginationQuery(offset, limit).GetMany(ctx, filter)
}

func (*personService) Update(ctx context.Context, uuid string, person models.Person) (string, error) {
	return PersonStorage.Update(ctx, uuid, person)
}

func (*personService) CreateSubstitute(ctx context.Context, person models.Person) (models.Person, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &person)
	person.IsSubstitute = true
	return PersonStorage.CreateSubstitute(ctx, person)
}

func (*personService) GetOneSubstitute(ctx context.Context, uuidPerson string) (models.Person, error) {
	return PersonStorage.GetOneSubstitute(ctx, uuidPerson)
}

func (*personService) GetSubstitutes(ctx context.Context) ([]models.Person, error) {
	return PersonStorage.GetSubstitutes(ctx)
}

func (*personService) GetNamePerson(ctx context.Context) ([]models.Person, error) {
	return PersonStorage.GetNamePerson(ctx)
}

func ValidateExistePerson(ctx context.Context, uuid string) (sqlnulls.NullString, error) {
	return storage.ValidateExistePerson(ctx, uuid)
}

func (*personService) GetManyWithFullInformation(ctx context.Context, filter string, offset, limit int) ([]models.Person, error) {
	return PersonStorage.PaginationQuery(offset, limit).GetManyWithFullInformation(ctx, filter)
}
