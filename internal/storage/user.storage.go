package storage

import (
	"context"
	"database/sql"
	"log"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/DasJalapa/reportes-salud/internal/helper"
	"github.com/DasJalapa/reportes-salud/internal/lib"
	"github.com/DasJalapa/reportes-salud/internal/models"
	"github.com/DasJalapa/reportes-salud/internal/mysql"
)

var db = mysql.Connect()

// NewUserStorage  constructor para userStorage
func NewUserStorage() UserStorage {
	return &repoUser{}
}

type repoUser struct{}

// UserStorage implementa todos lo metodos de usuario
type UserStorage interface {
	Create(ctx context.Context, user *models.User) (string, error)
	Update(ctx context.Context, uuid string, user *models.User) (string, error)
	Login(ctx context.Context, user *models.User) (models.User, error)
	GetOneUser(ctx context.Context, uuid string) (models.User, error)
	GetManyUsers(ctx context.Context) ([]models.User, error)
	Roles(ctx context.Context) ([]models.Rol, error)
	DeleteUser(ctx context.Context, uuid string) (string, error)

	ChangePassword(ctx context.Context, uuidUser, actualPassword, newPassword string) error

	UserInformationByToken(ctx context.Context, uuid string) (models.User, error)
}

func (*repoUser) Create(ctx context.Context, user *models.User) (string, error) {
	user.Username = strings.TrimSpace(user.Username)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	query := "INSERT INTO user (uuid, username, password, rol_id, uuidPerson) values (?, ?, ?, ?, ?);"
	_, err = db.QueryContext(ctx, query, user.ID, user.Username, string(hashedPassword), user.IDRol, user.Person)

	if err != nil {
		log.Println(err)
		return "", lib.ErrDuplicateUser
	}
	return user.ID, nil

}

func (*repoUser) Login(ctx context.Context, user *models.User) (models.User, error) {
	var response models.User
	var passwordClient string

	query := "SELECT u.uuid, u.username, u.password, r.role, u.uuidPerson FROM user u "
	query += "INNER JOIN rol r ON u.rol_id = r.id "
	query += "WHERE binary username = ?;"

	row := db.QueryRowContext(ctx, query, user.Username).Scan(&user.ID, &user.Username, &passwordClient, &user.Rol, &user.Person)

	if row == sql.ErrNoRows {
		return response, lib.ErrUserNotFound
	}

	if row != nil {
		return response, row
	}

	hashedPasswordDatabase := []byte(passwordClient)
	valuePassword := bcrypt.CompareHashAndPassword(hashedPasswordDatabase, []byte(user.Password))
	if valuePassword != nil {
		return response, lib.ErrUserNotFound
	}

	user.Password = ""
	response.Username = user.Username
	response.Rol = user.Rol

	token := helper.GenerateJWT(user)
	response.Token = token
	response.Person = user.Person

	return response, nil
}

func (*repoUser) Update(ctx context.Context, uuid string, user *models.User) (string, error) {
	query := "UPDATE user SET rol_id=?, username=? WHERE uuid = ?;"

	_, err := db.QueryContext(ctx, query, user.IDRol, user.Username, uuid)

	if err != nil {
		println(user.IDRol, user.Username)
		return "", err
	}
	println(user.IDRol, user.Username)
	return user.ID, nil
}

func (*repoUser) GetOneUser(ctx context.Context, uuid string) (models.User, error) {
	user := models.User{}

	query := `
	SELECT 
	u.username, 
	u.password, 
	r.role as rol,
    u.rol_id as id
    FROM user u
        INNER JOIN rol r ON u.rol_id = r.id
    WHERE uuid = ?;`

	err := db.QueryRowContext(ctx, query, uuid).Scan(
		&user.Username,
		&user.Password,
		&user.UserRol.Role,
		&user.IDRol,
	)

	if err == sql.ErrNoRows {
		return user, lib.ErrNotFound
	}

	if err != nil {
		return user, err
	}

	return user, nil
}

func (*repoUser) GetManyUsers(ctx context.Context) ([]models.User, error) {
	user := models.User{}
	users := []models.User{}

	query := `SELECT u.uuid, u.username, r.role FROM user u 
			  INNER JOIN rol r ON u.rol_id = r.id;`

	rows, err := db.QueryContext(ctx, query)
	if err == sql.ErrNoRows {
		return users, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Rol)

		if err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (*repoUser) Roles(ctx context.Context) ([]models.Rol, error) {
	rol := models.Rol{}
	rols := []models.Rol{}

	query := "SELECT id, role FROM rol;"
	rows, err := db.QueryContext(ctx, query)

	if err == sql.ErrNoRows {
		return rols, lib.ErrNotFound
	}

	for rows.Next() {
		err := rows.Scan(&rol.IDRol, &rol.TypeRol)
		if err != nil {
			return rols, err
		}
		rols = append(rols, rol)
	}

	return rols, err
}

func (*repoUser) ChangePassword(ctx context.Context, uuidUser, actualPassword, newPassword string) error {
	trans, err := db.BeginTx(ctx, nil)
	user := models.User{}

	if err != nil {
		return err
	}
	defer trans.Rollback()

	query := "SELECT uuid, password FROM user "
	query += "WHERE uuid = ?;"

	if err = db.QueryRowContext(ctx, query, uuidUser).Scan(&user.ID, &user.Password); err != nil {
		return err
	}

	hashedPasswordDatabase := []byte(user.Password)
	if valuePassword := bcrypt.CompareHashAndPassword(hashedPasswordDatabase, []byte(actualPassword)); valuePassword != nil {
		return lib.ErrUserNotFound
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

	if err != nil {
		return err
	}

	queryUpdate := "UPDATE user SET password = ? WHERE uuid = ?;"
	_, err = db.QueryContext(ctx, queryUpdate, string(hashedPassword), uuidUser)
	if err != nil {
		return err
	}

	if errtrans := trans.Commit(); errtrans != nil {
		return errtrans
	}
	return nil
}

func (*repoUser) UserInformationByToken(ctx context.Context, uuid string) (models.User, error) {
	var response models.User

	query := "SELECT uuid, username FROM user "
	query += "WHERE uuid = ?;"

	row := db.QueryRowContext(ctx, query, uuid).Scan(&response.ID, &response.Username)

	if row == sql.ErrNoRows {
		return response, lib.ErrUserNotFound
	}

	if row != nil {
		return response, row
	}

	return response, nil

}

func (*repoUser) DeleteUser(ctx context.Context, uuid string) (string, error) {
	queryUpdate := "DELETE FROM user WHERE uuid = ?;"

	rows, err := db.ExecContext(ctx, queryUpdate, uuid)
	if err != nil {
		return "", err
	}

	resultDelete, _ := rows.RowsAffected()
	if resultDelete == 0 {
		return "", lib.ErrNotFound
	}

	return uuid, nil
}
