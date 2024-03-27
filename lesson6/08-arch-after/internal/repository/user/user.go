package user

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/klakovsky/users/internal/domain"
)

type Repo struct {
	dbConn *pgx.Conn
}

func NewUserRepo(dbConn *pgx.Conn) *Repo {
	return &Repo{dbConn: dbConn}
}

func (r *Repo) GetUser(ctx context.Context, ID int) (domain.User, error) {
	var u domain.User
	query := "select id, name from users where id = $1"
	err := r.dbConn.QueryRow(ctx, query, ID).Scan(&u.ID, &u.Name)
	return u, err
}

func (r *Repo) CreateUser(ctx context.Context, name string) (domain.User, error) {
	u := domain.User{Name: name}
	query := "insert into users (name) values ($1) returning id"
	err := r.dbConn.QueryRow(ctx, query, name).Scan(&u.ID)
	if err != nil {
		return u, err
	}
	return u, err
}
