package main

import (
	"database/sql"
	"flag"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func main() {
	// Прапори для міграцій
	direction := flag.String("direction", "up", "Direction of migrations: up or down")
	step := flag.Int("step", 0, "Number of steps to migrate. 0 means all.")
	flag.Parse()

	// Завантаження змінних середовища
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("Failed to load .env.local: %v", err)
	}

	// Підключення до бази даних
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION"))
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Налаштування папки міграцій
	migrationsDir := "migrations"

	// Отримання поточного стану міграцій
	currentVersion, err := goose.GetDBVersion(db)
	if err != nil {
		log.Fatalf("Failed to get current migration version: %v", err)
	}

	switch *direction {
	case "up":
		if *step > 0 {
			for i := 0; i < *step; i++ {
				if err := goose.UpByOne(db, migrationsDir); err != nil {
					log.Fatalf("Failed to apply one migration: %v", err)
				}
			}
		} else {
			if err := goose.Up(db, migrationsDir); err != nil {
				log.Fatalf("Failed to apply all migrations: %v", err)
			}
		}
	case "down":
		if *step > 0 {
			targetVersion := currentVersion
			for i := 0; i < *step; i++ {
				targetVersion--
				if err := goose.DownTo(db, migrationsDir, targetVersion); err != nil {
					log.Fatalf("Failed to rollback one migration: %v", err)
				}
			}
		} else {
			if err := goose.Down(db, migrationsDir); err != nil {
				log.Fatalf("Failed to rollback all migrations: %v", err)
			}
		}
	default:
		log.Fatalf("Invalid direction: %s. Use 'up' or 'down'.", *direction)
	}

	log.Println("Migrations completed successfully!")
}
