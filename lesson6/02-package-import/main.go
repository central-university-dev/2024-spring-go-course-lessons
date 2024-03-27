package main

import (
	"fmt"

	"github.com/exaring/otelpgx"
	oldPgx "github.com/jackc/pgx"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var cfg pgconn.Config
	var errNoRows = pgx.ErrNoRows
	var tx oldPgx.Tx
	var opts []otelpgx.Option

	fmt.Println("hello world", cfg, errNoRows, tx, opts)
}
