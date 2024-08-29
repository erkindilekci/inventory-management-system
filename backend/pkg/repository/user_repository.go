package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/labstack/gommon/log"
	"ims-intro/pkg/domain"
)

type IUserRepository interface {
	GetUserByUsername(username string) (domain.User, error)
	SignUp(user domain.User) error
}

type UserRepository struct {
	dbPool *pgxpool.Pool
}

func NewUserRepository(dbPool *pgxpool.Pool) IUserRepository {
	return &UserRepository{dbPool}
}

func (repository *UserRepository) GetUserByUsername(username string) (domain.User, error) {
	ctx := context.Background()

	var user domain.User

	selectStatement := "SELECT id, username, password, role FROM users WHERE username = $1"
	userRow := repository.dbPool.QueryRow(ctx, selectStatement, username)

	err := userRow.Scan(&user.Id, &user.Username, &user.Password, &user.Role)
	if err != nil && err.Error() == "no rows in result set" {
		return domain.User{}, errors.New("error while finding user")
	}

	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (repository *UserRepository) SignUp(user domain.User) error {
	ctx := context.Background()

	insertStatement := "INSERT INTO users(username, password, role) VALUES ($1, $2, $3)"

	addNewUser, err := repository.dbPool.Exec(ctx, insertStatement, user.Username, user.Password, user.Role)
	if err != nil {
		log.Errorf("error while adding new user: %v", err)
		return err
	}

	log.Info(fmt.Sprint("User added successfully: %v", addNewUser))
	return nil
}
