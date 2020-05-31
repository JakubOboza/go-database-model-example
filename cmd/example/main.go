package main

import (
	"fmt"
	"log"

	"database.example/models/users"
	"database.example/utils"
	"github.com/golang-migrate/migrate/v4"
	// We need this for the postgres client to work!
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	// Migrations related imports
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	dbHost := utils.GetEnv("DBHOST", "localhost")
	dbPort := utils.GetEnv("DBPORT", "6677")
	dbUser := utils.GetEnv("DBUSER", "postgres")
	dbPassword := utils.GetEnv("DBPASSWORD", "")
	dbName := utils.GetEnv("DATABASE", "example_development")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Migrate database
	m, err := migrate.New(
		"file://db/migrations",
		connStr)
	if err != nil {
		log.Fatal(err)
	}
	if err := m.Up(); err != nil {
		fmt.Println("Database Migration", err)
	}

	// Setup connection poo
	db, err := sqlx.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	// Setup package
	users.Setup(db)

	newUser, err := users.Store.CreateUser("test@test.com", "this67is99!hard@pass")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("newUser ", newUser)

	user, err := users.Store.FindByEmail("test@test.com")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("User fetched by email", user)

}
