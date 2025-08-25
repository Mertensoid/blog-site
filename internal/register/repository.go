package register

import (
	"blog-site/package/bcrypt"
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UsersRepository struct {
	dbpool     *pgxpool.Pool
	logger     *zerolog.Logger
	cryptograf *bcrypt.Crypto
}

func NewUsersRepository(dbpool *pgxpool.Pool, logger *zerolog.Logger, cryptograf *bcrypt.Crypto) *UsersRepository {
	r := &UsersRepository{
		dbpool:     dbpool,
		logger:     logger,
		cryptograf: cryptograf,
	}
	return r
}

func (r *UsersRepository) addUser(form RegisterData) error {
	hash := r.cryptograf.HashPassword(form.Password)
	fmt.Println(hash)
	query := `INSERT INTO Users (email, pass, name, regtime) 
				VALUES (@email, @pass, @name, @regtime)
				`
	args := pgx.NamedArgs{
		"email":   form.Email,
		"pass":    hash,
		"name":    form.Name,
		"regtime": time.Now(),
	}
	_, err := r.dbpool.Exec(context.Background(), query, args)
	if err != nil {
		return fmt.Errorf("Невозможно создать пользователя: %w", err)
	}
	return nil
}

func (r *UsersRepository) GetUser(email string) User {
	query := `SELECT * FROM users
				WHERE email = $1`
	fmt.Println(query)
	user := User{}
	err := r.dbpool.QueryRow(context.Background(), query,
		email).Scan(&user.Id, &user.Email, &user.Password, &user.Name, &user.RegTime)
	if err != nil {
		r.logger.Error().Msg(err.Error())
	}
	return user
}

func (r *UsersRepository) checkUser(form LoginForm) (User, error) {

	query := `SELECT * FROM users WHERE email = $1`
	user := User{}
	err := r.dbpool.QueryRow(context.Background(), query, form.Email).Scan(&user.Id,
		&user.Email, &user.Password, &user.Name, &user.RegTime)
	if err != nil {
		r.logger.Error().Msg(err.Error())
		return User{}, err
	}
	if !r.cryptograf.CheckPasswordHash(form.Password, user.Password) {
		return User{}, errors.New("Incorrect password")
	}
	return user, nil
}
