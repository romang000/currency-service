package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserAlreadyExist = errors.New("user already exist")
	ErrUserNotFound     = errors.New("user not found")
)

type User struct {
	Login    string
	Password string
}

type UserRepository struct {
	db *sql.DB
}

func NewUser(db *sql.DB) (UserRepository, error) {
	return UserRepository{db: db}, nil
}

func (repo *UserRepository) AddUser(ctx context.Context, user User) error {

	if _, err := repo.GetUser(ctx, user.Login); err == nil {
		return ErrUserAlreadyExist
	}

	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = repo.db.ExecContext(
		ctx,
		`INSERT INTO users (username, password_hash) VALUES ($1, $2)`,
		user.Login, hashedPassword,
	)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

func (repo *UserRepository) GetUser(ctx context.Context, login string) (User, error) {
	var user User
	err := repo.db.QueryRowContext(
		ctx,
		`SELECT username, password_hash FROM users
				WHERE username = $1
			`,
		login,
	).Scan(&user.Login, &user.Password)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return User{}, ErrUserNotFound
		}
		return User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed hash password: %w", err)
	}
	return string(res), nil
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}
	return true
}
