package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"

	//"github.com/Mynor2397/sqlnulls"
	"github.com/google/uuid"
)

type personEducationService struct {
}

var PersonEducationStorage storage.PersonEducationStorage

// NewPersonEducationService retorna un nuevo servicio para los usuarios
func NewPersonEducationService(personEducationStorage storage.PersonEducationStorage) PersonEducationService {
	PersonEducationStorage = personEducationStorage
	return &personEducationService{}
}

// PersonEducationService implementa el conjunto de metodos de servicio para usuario
type PersonEducationService interface {
	Create(ctx context.Context, references models.PersonEducation) (models.PersonEducation, error)
	GetEducations(ctx context.Context, uuid string) ([]models.PersonEducation, error)
	DeleteEducations(ctx context.Context, uuid string) (string, error)
}

func (*personEducationService) Create(ctx context.Context, personEducation models.PersonEducation) (models.PersonEducation, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &personEducation)
	return PersonEducationStorage.Create(ctx, personEducation)
}

func (*personEducationService) GetEducations(ctx context.Context, uuid string) ([]models.PersonEducation, error) {
	return PersonEducationStorage.GetEducations(ctx, uuid)
}

func (*personEducationService) DeleteEducations(ctx context.Context, uuid string) (string, error) {
	return PersonEducationStorage.DeleteEducations(ctx, uuid)
}
