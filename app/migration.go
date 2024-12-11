package app

import (
	"database/sql"
	"errors"
	"flag"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	direction := flag.String("direction", "up", "Direction of migrations: up or down")
	step := flag.Int("step", 0, "Number of steps to migrate. 0 means all.")
	flag.Parse()

	dsn, err := dsn()
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to create driver instance: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to create migration instance: %v", err)
	}

	// Виконання міграцій
	switch *direction {
	case "up":
		if *step > 0 {
			for i := 0; i < *step; i++ {
				if err := m.Steps(1); err != nil && err != migrate.ErrNoChange {
					log.Fatalf("Failed to apply migration step: %v", err)
				}
			}
		} else {
			if err := m.Up(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Failed to apply migrations: %v", err)
			}
		}
	case "down":
		if *step > 0 {
			for i := 0; i < *step; i++ {
				if err := m.Steps(-1); err != nil && err != migrate.ErrNoChange {
					log.Fatalf("Failed to apply migration step: %v", err)
				}
			}
		} else {
			if err := m.Down(); err != nil && err != migrate.ErrNoChange {
				log.Fatalf("Failed to rollback migrations: %v", err)
			}
		}
	default:
		log.Fatalf("Invalid direction: %s. Use 'up' or 'down'.", *direction)
	}

	log.Println("Migrations completed successfully!")
}

func dsn() (string, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return "", errors.New("some obligatory env db parameters missing")
	}

	return "postgres://" + dbUser + ":" + dbPassword + "@" + dbHost + ":" + dbPort + "/" + dbName + "?sslmode=disable", nil
}
