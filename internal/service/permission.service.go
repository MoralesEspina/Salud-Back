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
	GetPermissionss(ctx context.Context, startDate, endDate, status string) ([]models.Permission, error)
	GetPermissions(ctx context.Context, uuidUser, role, startDate, endDate, status string) ([]models.Permission, error)
	GetOnePermission(ctx context.Context, uuid string) (models.Permission, error)
	GetOnePermissionWithName(ctx context.Context, uuid string) (models.Permission, error)
	UpdatePermission(ctx context.Context, request models.Permission, uuid, rol string) (string, error)
	DeletePermission(ctx context.Context, uuid string) (string, error)
	GetBosssesOne(ctx context.Context) ([]models.Person, error)
	GetBosssesTwo(ctx context.Context) ([]models.Person, error)
	GetPermissionsBossOne(ctx context.Context, uuid string) ([]models.Permission, error)
	GetPermissionsBossTwo(ctx context.Context, uuid string) ([]models.Permission, error)
	GetUserPermissionsActives(ctx context.Context, uuid string) ([]models.Permission, error)
	GetUserPermissions(ctx context.Context, uuid string) ([]models.Permission, error)
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
	request.Reason = "-"

	return IPermission.Create(ctx, request)
}

func (r *permissionService) GetPermissions(ctx context.Context, uuidPerson, role, startDate, endDate, status string) ([]models.Permission, error) {
	args := []interface{}{}
	var query string

	if role == "admin" {
		if status == "Todas" {
			query = `SELECT r.uuid, r.submittedAt, r.permissionDate, p.fullname, r.status, r.uuidPerson FROM permission r
					INNER JOIN person p ON r.uuidPerson = p.uuid
					WHERE r.submittedAt >= ? AND r.submittedAt <= ?
					ORDER BY r.submittedAt DESC`
			args = append(args, startDate, endDate)
		} else {
			query = `SELECT r.uuid, r.submittedAt, r.permissionDate, p.fullname, r.status, r.uuidPerson FROM permission r
				INNER JOIN person p ON r.uuidPerson = p.uuid
				WHERE r.submittedAt >= ? AND r.submittedAt <= ? AND r.status = ?
				ORDER BY r.submittedAt DESC`
			args = append(args, startDate, endDate, status)
		}
	} else {
		query = `SELECT r.uuid, r.submittedAt, r.permissionDate, p.fullname, r.status, r.uuidPerson FROM permission r 
		INNER JOIN person p ON r.uuidPerson = p.uuid 
		WHERE r.bossOne = (Select uuid from user u where u.uuidPerson = ?) 
		AND r.submittedAt >= ? AND r.submittedAt <= ? AND r.status NOT LIKE 'En Espera' 
		ORDER BY r.submittedAt ASC;`
		args = append(args, uuidPerson, startDate, endDate)
	}

	return IPermission.GetPermissions(ctx, query, args)
}

func (r *permissionService) GetPermissionss(ctx context.Context, startDate, endDate, status string) ([]models.Permission, error) {
	return IPermission.GetPermissionss(ctx, startDate, endDate, status)
}

func (r *permissionService) GetOnePermission(ctx context.Context, uuid string) (models.Permission, error) {
	return IPermission.GetOnePermission(ctx, uuid)
}

func (r *permissionService) GetOnePermissionWithName(ctx context.Context, uuid string) (models.Permission, error) {
	return IPermission.GetOnePermissionWithName(ctx, uuid)
}

func (r *permissionService) UpdatePermission(ctx context.Context, request models.Permission, uuid, rol string) (string, error) {
	time := lib.TimeZone("America/Guatemala")
	request.ModifiedAt = time.DateTime
	switch {
	case request.BossOne == request.BossTwo:
		request.StatusBossTwo = request.StatusBossOne
		request.Status = request.StatusBossOne
	case request.StatusBossOne == "Denegada" && request.StatusBossTwo == "En Espera":
		request.StatusBossTwo = "En Espera"
		request.Status = "Denegada"
	case request.StatusBossOne == "Aceptada" && request.StatusBossTwo == "En Espera":
		request.StatusBossTwo = "En Espera"
		request.Status = "En Espera"
	case request.StatusBossTwo == "Denegada":
		request.Status = "Denegada"
	case request.StatusBossTwo == "Aceptada":
		request.Status = "Aceptada"
	}

	return IPermission.UpdatePermission(ctx, request, uuid, rol)
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

func (*permissionService) GetPermissionsBossOne(ctx context.Context, uuid string) ([]models.Permission, error) {
	return IPermission.GetPermissionsBossOne(ctx, uuid)
}

func (*permissionService) GetPermissionsBossTwo(ctx context.Context, uuid string) ([]models.Permission, error) {
	return IPermission.GetPermissionsBossTwo(ctx, uuid)
}

func (*permissionService) GetUserPermissionsActives(ctx context.Context, uuid string) ([]models.Permission, error) {
	return IPermission.GetUserPermissionsActives(ctx, uuid)
}

func (*permissionService) GetUserPermissions(ctx context.Context, uuid string) ([]models.Permission, error) {
	return IPermission.GetUserPermissions(ctx, uuid)
}
