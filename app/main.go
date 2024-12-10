package app

import (
	"context"
	"database/sql"
	"errors"
	"ipr/middleware"
	middlaware "ipr/middleware/http"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"ipr/common"
	"ipr/modules/user"
)

func main() {
	if err := godotenv.Load(".env.local"); err != nil {
		log.Fatalf("Failed to load .env.local: %v", err)
	}

	db, err := initDB()
	if err != nil {
		log.Fatalf("Failed to init db: %v", err)
	}

	defer db.Close()

	isDev := os.Getenv("APP_ENV") == "dev"

	sessionManager := common.NewSessionManager("./sessions", "super-secret-key", "ipr-session")
	httpMiddlewares := []middleware.Middleware{
		middlaware.ErrorMiddleware,
	}
	router := common.NewRouter(httpMiddlewares...)
	common.InitializeRenderer("./templates", isDev)
	ctx := context.Background()

	// > module user
	userDeps := user.NewDependencies(db, ctx)
	user.RegisterRoutes(userDeps, router, sessionManager)
	// < module user

	// Start the HTTP server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func initDB() (*sql.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		return nil, errors.New("some obligatory env db parameters missing")
	}

	connStr := "host=" + dbHost +
		" port=" + dbPort +
		" user=" + dbUser +
		" password=" + dbPassword +
		" dbname=" + dbName +
		" sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
