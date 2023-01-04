package service

import (
	"context"

	"github.com/google/uuid"

	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/storage"
)

type userService struct{}

var Userstorage storage.UserStorage

// NewUserService retorna un nuevo servicio para los usuarios
func NewUserService(userstorage storage.UserStorage) UserService {
	Userstorage = userstorage
	return &userService{}
}

// UserService implementa el conjunto de metodos de servicio para usuario
type UserService interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Login(ctx context.Context, user *models.User) (models.User, error)
	GetOneUser(ctx context.Context, uuid string) (models.User, error)
	Update(ctx context.Context, uuid string, user *models.User) (string, error)
	ManyUsers(ctx context.Context) ([]models.User, error)
	ManyAdminsAndMembers(ctx context.Context) ([]models.User, error)
	ManyEmployees(ctx context.Context) ([]models.User, error)
	ManyBosses(ctx context.Context) ([]models.User, error)
	Roles(ctx context.Context) ([]models.Rol, error)
	ChangePassword(ctx context.Context, uuidUser, actualPassword, newPassword string) error
	DeleteUser(ctx context.Context, uuid string) (string, error)

	UserInformationByToken(ctx context.Context, uuid string) (models.User, error)
}

// UserCreate es el servicio de conexion al storage de crear usuario
func (*userService) Create(ctx context.Context, user *models.User) (string, error) {
	user.ID = uuid.New().String()

	return Userstorage.Create(ctx, user)
}

func (*userService) Update(ctx context.Context, uuid string, user *models.User) (string, error) {
	return Userstorage.Update(ctx, uuid, user)
}

// UserLogin es el servicio de conexion al storage de login de usuario
func (*userService) Login(ctx context.Context, user *models.User) (models.User, error) {
	return Userstorage.Login(ctx, user)
}

func (*userService) GetOneUser(ctx context.Context, uuid string) (models.User, error) {
	return Userstorage.GetOneUser(ctx, uuid)
}

func (*userService) ManyUsers(ctx context.Context) ([]models.User, error) {
	return Userstorage.GetManyUsers(ctx)
}

func (*userService) ManyAdminsAndMembers(ctx context.Context) ([]models.User, error) {
	return Userstorage.GetManyAdminsAndMembers(ctx)
}

func (*userService) ManyEmployees(ctx context.Context) ([]models.User, error) {
	return Userstorage.GetManyEmployees(ctx)
}

func (*userService) ManyBosses(ctx context.Context) ([]models.User, error) {
	return Userstorage.GetManyBosses(ctx)
}

func (*userService) Roles(ctx context.Context) ([]models.Rol, error) {
	return Userstorage.Roles(ctx)
}

func (*userService) ChangePassword(ctx context.Context, uuidUser, actualPassword, newPassword string) error {
	return Userstorage.ChangePassword(ctx, uuidUser, actualPassword, newPassword)
}

func (*userService) DeleteUser(ctx context.Context, uuid string) (string, error) {
	return Userstorage.DeleteUser(ctx, uuid)
}

func (*userService) UserInformationByToken(ctx context.Context, uuid string) (models.User, error) {
	return Userstorage.UserInformationByToken(ctx, uuid)
}
