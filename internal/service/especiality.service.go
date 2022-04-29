package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/google/uuid"
)

type especialityService struct{}

var EspecialityStorage storage.IEspecialityStorage

// NewEspecialityService retorna un nuevo servicio para los usuarios
func NewEspecialityService(jobStorage storage.IEspecialityStorage) IEspecialityService {
	EspecialityStorage = jobStorage
	return &especialityService{}
}

type IEspecialityService interface {
	CreateEspeciality(ctx context.Context, especiality models.Especiality) (string, error)
	Especialities(ctx context.Context) ([]models.Especiality, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, especiality models.Especiality, uuid string) (models.Especiality, error)
	OneEspeciality(ctx context.Context, uuid string) (models.Especiality, error)
}

func (*especialityService) CreateEspeciality(ctx context.Context, especiality models.Especiality) (string, error) {
	uuidString := fmt.Sprintf(`{"uuid_especiality": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &especiality)
	return EspecialityStorage.CreateEspeciality(ctx, especiality)
}

func (*especialityService) Especialities(ctx context.Context) ([]models.Especiality, error) {
	return EspecialityStorage.Especialities(ctx)
}

func (*especialityService) Delete(ctx context.Context, uuid string) error {
	return EspecialityStorage.Delete(ctx, uuid)
}

func (*especialityService) Update(ctx context.Context, especiality models.Especiality, uuid string) (models.Especiality, error) {
	return EspecialityStorage.Update(ctx, especiality, uuid)
}

func (*especialityService) OneEspeciality(ctx context.Context, uuid string) (models.Especiality, error) {
	return EspecialityStorage.OneEspeciality(ctx, uuid)
}
