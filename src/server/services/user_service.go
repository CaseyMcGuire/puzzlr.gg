package services

import (
	"context"
	ent "gameboard/src/server/db/ent/codegen"
	"gameboard/src/server/db/ent/codegen/user"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	dbClient ent.Client
}

func (u UserService) CreateUser(ctx context.Context, email string, password string) (*ent.User, error) {
	hashedPassword, err := HashPassword(password)
	if err != nil {
		return nil, err
	}

	newUser, err := u.dbClient.User.
		Create().
		SetEmail(email).
		SetHashedPassword(hashedPassword).
		Save(ctx)

	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (u UserService) UserExists(ctx context.Context, email string) (bool, error) {
	exists, err := u.dbClient.User.
		Query().
		Where(user.Email(email)).
		Exist(ctx)

	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u UserService) UserWithPasswordExists(ctx context.Context, email string, password string) (bool, error) {
	storedUser, err := u.dbClient.User.
		Query().
		Where(user.Email(email)).
		Only(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}

	return checkPassword(storedUser.HashedPassword, password), nil
}

func HashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

func checkPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func NewUserService(
	dbClient ent.Client,
) *UserService {

	return &UserService{
		dbClient: dbClient,
	}
}
