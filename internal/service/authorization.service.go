package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/mysql"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/DasJalapa/reportes-salud/internal/storage/cross"
)

type authorizationService struct{}

const region string = "America/Guatemala"

var AuthorizationStorage storage.AuthorizationStorage

// NewAuthorizationService retorna un nuevo servicio para los usuarios
func NewAuthorizationService(authorizationStorage storage.AuthorizationStorage) AuthorizationService {
	AuthorizationStorage = authorizationStorage
	return &authorizationService{}
}

type AuthorizationService interface {
	Create(ctx context.Context, auhtorization models.Authorization) (models.Authorization, error)
	GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error)
	GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error)
	UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error)
	GetOnlyAuthorizationPDF(ctx context.Context, UUIDAuthorization string) (models.Authorization, error)
	GetBosses(ctx context.Context) ([]models.Boss, error)
	UpdateBoss(ctx context.Context, bosses models.Boss, id string) (models.Boss, error)

	VacationsReport(ctx context.Context, startDateReport, endDateReport string) ([]models.Authorization, error)
}

func (*authorizationService) Create(ctx context.Context, request models.Authorization) (models.Authorization, error) {
	register, err := cross.GenerateDynamicNumberRegister("autorization")
	if err != nil {
		return request, err
	}

	time := lib.TimeZone(region)
	if request.SubmittedAt == "" {
		request.SubmittedAt = time.DateTime
	}

	request.UUIDAuthorization = uuid.New().String()
	request.Register = register
	request.ModifiedAt = time.DateTime
	return AuthorizationStorage.Create(ctx, request)
}

func (*authorizationService) GetManyAuthorizations(ctx context.Context) ([]models.Authorization, error) {
	return AuthorizationStorage.GetManyAuthorizations(ctx)
}

func (*authorizationService) GetOnlyAuthorization(ctx context.Context, uuid string) (models.Authorization, error) {
	return AuthorizationStorage.GetOnlyAuthorization(ctx, uuid)
}

func (*authorizationService) UpdateAuthorization(ctx context.Context, authorization models.Authorization, uuid string) (models.Authorization, error) {
	time := lib.TimeZone(region)
	authorization.ModifiedAt = time.DateTime
	return AuthorizationStorage.UpdateAuthorization(ctx, authorization, uuid)
}

func (*authorizationService) GetOnlyAuthorizationPDF(ctx context.Context, UUIDAuthorization string) (models.Authorization, error) {
	return storage.DataPDFAuthorization(ctx, UUIDAuthorization, mysql.Connect())
}

func (*authorizationService) VacationsReport(ctx context.Context, startDateReport, endDateReport string) ([]models.Authorization, error) {
	return AuthorizationStorage.VacationsReport(ctx, startDateReport, endDateReport)
}

func (*authorizationService) GetBosses(ctx context.Context) ([]models.Boss, error) {
	return AuthorizationStorage.GetBosses(ctx)
}

func (*authorizationService) UpdateBoss(ctx context.Context, bosses models.Boss, id string) (models.Boss, error) {
	return AuthorizationStorage.UpdateBoss(ctx, bosses, id)
}
