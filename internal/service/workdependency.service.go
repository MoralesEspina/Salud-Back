package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

type workDependencyService struct{}

var WorkDependencyStorage storage.IWorkDependencyStorage

// NewAuthorizationService retorna un nuevo servicio para los usuarios
func NewWorkDependencyService(workDependencyStorage storage.IWorkDependencyStorage) IWorkDependencyService {
	WorkDependencyStorage = workDependencyStorage
	return &workDependencyService{}
}

type IWorkDependencyService interface {
	GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error)
	CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error)
	OneWorkDependency(ctx context.Context, uuid string) (models.WorkDependency, error)
	Delete(ctx context.Context, uuid string) error
	Update(ctx context.Context, workdependency models.WorkDependency, uuid string) (models.WorkDependency, error)
}

func (*workDependencyService) GetManyWorkDependency(ctx context.Context) ([]models.WorkDependency, error) {
	return WorkDependencyStorage.GetManyWorkDependency(ctx)
}

func (*workDependencyService) CreateWorkDependency(ctx context.Context, dependency models.WorkDependency) (string, error) {
	uuidString := fmt.Sprintf(`{"uuid_work": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &dependency)
	return WorkDependencyStorage.CreateWorkDependency(ctx, dependency)
}

func (*workDependencyService) OneWorkDependency(ctx context.Context, uuid string) (models.WorkDependency, error) {
	return WorkDependencyStorage.OneWorkDependency(ctx, uuid)
}

func (*workDependencyService) Delete(ctx context.Context, uuid string) error {
	return WorkDependencyStorage.Delete(ctx, uuid)
}

func (*workDependencyService) Update(ctx context.Context, workdependency models.WorkDependency, uuid string) (models.WorkDependency, error) {
	return WorkDependencyStorage.Update(ctx, workdependency, uuid)
}
