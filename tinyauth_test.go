package tinyauth

import (
	"log"
	"testing"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	dbc, err := sqlx.Connect("postgres",
		"host=localhost port=5432 user=listmonk password=listmonk dbname=listmonk sslmode=off")
	if err != nil {
		log.Fatalf("error connecting to DB: %v", err)
	}
	db = dbc
}

func TestCreateUser(t *testing.T) {
	ta := New(Options{
		UsersTable: "users",
	}, pgstore.New())
}
