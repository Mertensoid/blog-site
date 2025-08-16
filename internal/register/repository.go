package register

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

type UsersRepository struct {
	dbpool *pgxpool.Pool
	logger *zerolog.Logger
}

func NewUsersRepository(dbpool *pgxpool.Pool, logger *zerolog.Logger) *UsersRepository {
	r := &UsersRepository{
		dbpool: dbpool,
		logger: logger,
	}
	return r
}

func (r *UsersRepository) addUser(form RegisterData) error {
	query := `INSERT INTO Users (email, pass, name, regtime) 
				VALUES (@email, @pass, @name, @regtime)
				`
	args := pgx.NamedArgs{
		"email":   form.Email,
		"pass":    form.Password,
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
