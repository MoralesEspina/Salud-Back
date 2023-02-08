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

type curriculumService struct {
}

var CurriculumStorage storage.CurriculumStorage

// NewCurriculumService retorna un nuevo servicio para los usuarios
func NewCurriculumService(curriculumStorage storage.CurriculumStorage) CurriculumService {
	CurriculumStorage = curriculumStorage
	return &curriculumService{}
}

// CurriculumService implementa el conjunto de metodos de servicio para usuario
type CurriculumService interface {
	Create(ctx context.Context, curriculum models.Curriculum) (models.Curriculum, error)
	GetOne(ctx context.Context, uuid string) (models.Curriculum, error)
	Update(ctx context.Context, uuid string, curriculum models.Curriculum) (string, error)
	DeleteCurriculum(ctx context.Context, uuid string) (string, error)
}

func (*curriculumService) Create(ctx context.Context, curriculum models.Curriculum) (models.Curriculum, error) {
	uuidString := fmt.Sprintf(`{"uuid": "%s"}`, uuid.New().String())
	json.Unmarshal([]byte(uuidString), &curriculum)
	return CurriculumStorage.Create(ctx, curriculum)
}

func (*curriculumService) GetOne(ctx context.Context, uuid string) (models.Curriculum, error) {
	return CurriculumStorage.GetOne(ctx, uuid)
}

func (*curriculumService) Update(ctx context.Context, uuid string, curriculum models.Curriculum) (string, error) {
	return CurriculumStorage.Update(ctx, uuid, curriculum)
}

func (*curriculumService) DeleteCurriculum(ctx context.Context, uuid string) (string, error) {
	return CurriculumStorage.DeleteCurriculum(ctx, uuid)
}
