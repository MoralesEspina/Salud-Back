package service

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/DasJalapa/reportes-salud/internal/storage/cross"
	"github.com/google/uuid"
)

type requestVacationService struct{}

var IRequestVacation storage.IRequestVacationStorage

func NewRequestVacationService(IRequestVacationStorage storage.IRequestVacationStorage) IRequestVacationService {
	IRequestVacation = IRequestVacationStorage
	return &requestVacationService{}
}

type IRequestVacationService interface {
	Create(ctx context.Context, request models.RequestVacation, uuidUser string) (models.RequestVacation, error)
	GetRequestsVacations(ctx context.Context, uuidUser, role string) ([]models.RequestVacation, error)
	GetOneRequestVacation(ctx context.Context, uuid string) (models.RequestVacation, error)
	UpdateRequestVacation(ctx context.Context, request models.RequestVacation, uuid string) (string, error)
	DeleteRequestVacation(ctx context.Context, uuid string) (string, error)
}

func (r *requestVacationService) Create(ctx context.Context, request models.RequestVacation, uuidUser string) (models.RequestVacation, error) {
	register, err := cross.GenerateDynamicNumberRegister("vacationrequest")
	if err != nil {
		return request, err
	}

	time := lib.TimeZone("America/Guatemala")
	request.UUIDUser = uuidUser
	request.UUIDRequestVacation = uuid.New().String()
	request.SubmittedAt = time.DateTime
	request.ModifiedAt = time.DateTime
	request.Register = register

	return IRequestVacation.Create(ctx, request)
}

func (r *requestVacationService) GetRequestsVacations(ctx context.Context, uuidUser, role string) ([]models.RequestVacation, error) {
	args := []interface{}{}
	var query string

	if role == "admin" {
		query = `
		SELECT r.uuid, p.fullname, p.cui, r.submittedAt FROM vacationrequest r
		INNER JOIN person p ON r.person_uuid = p.uuid
		ORDER BY r.submittedAt DESC`
	} else {
		query = `
		SELECT r.uuid, p.fullname, p.cui, r.submittedAt FROM vacationrequest r
		INNER JOIN person p ON r.person_uuid = p.uuid
		WHERE user_uuid = ?
		ORDER BY r.submittedAt DESC`
		args = append(args, uuidUser)
	}

	return IRequestVacation.GetRequestsVacations(ctx, query, args)
}

func (r *requestVacationService) GetOneRequestVacation(ctx context.Context, uuid string) (models.RequestVacation, error) {
	return IRequestVacation.GetOneRequestVacation(ctx, uuid)
}

func (r *requestVacationService) UpdateRequestVacation(ctx context.Context, request models.RequestVacation, uuid string) (string, error) {
	request.ModifiedAt = lib.TimeZone("America/Guatemala").DateTime
	return IRequestVacation.UpdateRequestVacation(ctx, request, uuid)
}
func (*requestVacationService) DeleteRequestVacation(ctx context.Context, uuid string) (string, error) {
	return IRequestVacation.DeleteRequestVacation(ctx, uuid)
}
