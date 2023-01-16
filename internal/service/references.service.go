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

type referencesService struct {
}

var ReferencesStorage storage.ReferencesStorage

// NewReferencesService retorna un nuevo servicio para los usuarios
func NewReferencesService(referencesStorage storage.ReferencesStorage) ReferencesService {
	ReferencesStorage = referencesStorage
	return &referencesService{}
}

// ReferencesService implementa el conjunto de metodos de servicio para usuario
type ReferencesService interface {
	CreateRefFamiliar(ctx context.Context, references models.References) (models.References, error)
	GetRefPer(ctx context.Context, uuid string) ([]models.References, error)
	GetRefFam(ctx context.Context, uuid string) ([]models.References, error)
	DeleteRefPer(ctx context.Context, uuid string) (string, error)
	DeleteRefFam(ctx context.Context, uuid string) (string, error)
}

func (*referencesService) CreateRefFamiliar(ctx context.Context, references models.References) (models.References, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &references)
	return ReferencesStorage.CreateRefFamiliar(ctx, references)
}

func (*referencesService) GetRefPer(ctx context.Context, uuid string) ([]models.References, error) {
	return ReferencesStorage.GetRefPer(ctx, uuid)
}

func (*referencesService) GetRefFam(ctx context.Context, uuid string) ([]models.References, error) {
	return ReferencesStorage.GetRefFam(ctx, uuid)
}

func (*referencesService) DeleteRefPer(ctx context.Context, uuid string) (string, error) {
	return ReferencesStorage.DeleteRefPer(ctx, uuid)
}

func (*referencesService) DeleteRefFam(ctx context.Context, uuid string) (string, error) {
	return ReferencesStorage.DeleteRefFam(ctx, uuid)
}
