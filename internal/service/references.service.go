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
	Create(ctx context.Context, references models.References) (models.References, error)
	GetReferences(ctx context.Context, uuid string) ([]models.References, error)
}

func (*referencesService) Create(ctx context.Context, references models.References) (models.References, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &references)
	return ReferencesStorage.Create(ctx, references)
}

func (*referencesService) GetReferences(ctx context.Context, uuid string) ([]models.References, error) {
	return ReferencesStorage.GetReferences(ctx, uuid)
}
