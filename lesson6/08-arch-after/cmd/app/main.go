package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"

	"github.com/klakovsky/users/internal/gateway/httpapi"
	userrepo "github.com/klakovsky/users/internal/repository/user"
	"github.com/klakovsky/users/internal/usecase"
)

func main() {
	ctx := context.Background()
	dbConn, err := pgx.Connect(ctx, os.Getenv("PG_URL"))
	if err != nil {
		log.Fatal(err.Error())
	}

	repo := userrepo.NewUserRepo(dbConn)
	uc := usecase.New(repo, time.Second, time.Second)

	s := httpapi.NewServer(uc)
	s.Run()
}
