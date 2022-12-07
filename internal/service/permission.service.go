package service

import (
	"context"

	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
	"github.com/DasJalapa/reportes-salud/internal/storage/cross"
	"github.com/google/uuid"
)

type permissionService struct{}

var IPermission storage.IPermissionStorage

func NewPermissionService(IPermissionStorage storage.IPermissionStorage) IPermissionService {
	IPermission = IPermissionStorage
	return &permissionService{}
}

type IPermissionService interface {
	Create(ctx context.Context, request models.Permission, uuidUser string) (models.Permission, error)
	GetPermissions(ctx context.Context, uuidUser, role string) ([]models.Permission, error)
	GetOnePermission(ctx context.Context, uuid string) (models.Permission, error)
	UpdatePermission(ctx context.Context, request models.Permission, uuid string) (string, error)
	DeletePermission(ctx context.Context, uuid string) (string, error)
	GetBosssesOne(ctx context.Context) ([]models.Person, error)
	GetBosssesTwo(ctx context.Context) ([]models.Person, error)
}

func (r *permissionService) Create(ctx context.Context, request models.Permission, uuidUser string) (models.Permission, error) {
	register, err := cross.GenerateDynamicNumberRegister("permission")
	if err != nil {
		return request, err
	}

	time := lib.TimeZone("America/Guatemala")
	request.Uuid = uuid.New().String()
	request.SubmittedAt = time.DateTime
	request.ModifiedAt = time.DateTime
	request.Register = register
	request.StatusBossOne = "En Espera"
	request.StatusBossTwo = "En Espera"
	request.Status = "En Espera"

	return IPermission.Create(ctx, request)
}

func (r *permissionService) GetPermissions(ctx context.Context, uuidUser, role string) ([]models.Permission, error) {
	args := []interface{}{}
	var query string

	if role == "admin" {
		query = `
		SELECT r.uuid, r.submittedAt, p.fullname, r.bossOne, r.bossTwo FROM permission r
		INNER JOIN person p ON r.uuidPerson = p.uuid`

	} else {
		query = `
		SELECT r.uuid, p.fullname, p.cui, r.submmittedAt FROM vacationrequest r
		INNER JOIN person p ON r.uuidPerson = p.uuid
		WHERE user_uuid = ?
		ORDER BY r.submittedAt DESC`
		args = append(args, uuidUser)
	}

	return IPermission.GetPermissions(ctx, query, args)
}

func (r *permissionService) GetOnePermission(ctx context.Context, uuid string) (models.Permission, error) {
	return IPermission.GetOnePermission(ctx, uuid)
}

func (r *permissionService) UpdatePermission(ctx context.Context, request models.Permission, uuid string) (string, error) {
	return IPermission.UpdatePermission(ctx, request, uuid)
}
func (*permissionService) DeletePermission(ctx context.Context, uuid string) (string, error) {
	return IPermission.DeletePermission(ctx, uuid)
}

func (*permissionService) GetBosssesOne(ctx context.Context) ([]models.Person, error) {
	return IPermission.GetBosssesOne(ctx)
}

func (*permissionService) GetBosssesTwo(ctx context.Context) ([]models.Person, error) {
	return IPermission.GetBosssesTwo(ctx)
}
