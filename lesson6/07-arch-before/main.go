package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx"
)

func CreateDBConn() (*pgx.Conn, error) {
	return pgx.Connect(pgx.ConnConfig{
		Host:     os.Getenv("PG_HOST"),
		Database: "users",
		User:     os.Getenv("PG_USER"),
		Password: os.Getenv("PG_PASSWD"),
	})
}

type UserResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func main() {
	dbConn, err := CreateDBConn()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	app.Get("/users/:user_id", func(c *fiber.Ctx) error {
		userID := c.Params("user_id")
		q := "select u.id, u.name from users u where u.id = $1"

		u := &UserResponse{}
		if err := dbConn.QueryRow(q, userID).Scan(&u.ID, &u.Name); err != nil {
			return err
		}

		return c.JSON(u)
	})

	app.Post("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		q := "insert into users(name) values($1) returning id"

		u := &UserResponse{Name: name}

		if err := dbConn.QueryRow(q, name).Scan(&u.ID); err != nil {
			return err
		}

		return c.JSON(u)
	})

	log.Fatal(app.Listen(":3000"))
}
