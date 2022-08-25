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

type workExpService struct {
}

var WorkExpStorage storage.WorkExpStorage

// NewWorkExpService retorna un nuevo servicio para los usuarios
func NewWorkExpService(workExpStorage storage.WorkExpStorage) WorkExpService {
	WorkExpStorage = workExpStorage
	return &workExpService{}
}

// WorkExpService implementa el conjunto de metodos de servicio para usuario
type WorkExpService interface {
	Create(ctx context.Context, references models.WorkExp) (models.WorkExp, error)
	GetWorks(ctx context.Context, uuid string) ([]models.WorkExp, error)
	DeleteWorks(ctx context.Context, uuid string) (string, error)
}

func (*workExpService) Create(ctx context.Context, workExp models.WorkExp) (models.WorkExp, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &workExp)
	return WorkExpStorage.Create(ctx, workExp)
}

func (*workExpService) GetWorks(ctx context.Context, uuid string) ([]models.WorkExp, error) {
	return WorkExpStorage.GetWorks(ctx, uuid)
}

func (*workExpService) DeleteWorks(ctx context.Context, uuid string) (string, error) {
	return WorkExpStorage.DeleteWorks(ctx, uuid)
}
