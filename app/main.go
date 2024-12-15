package main

import (
	"context"
	"database/sql"
	"ipr/infra/router"
	"ipr/infra/session"
	"ipr/infra/template"
	"ipr/modules/daily_activity"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
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

	sessionManager := session.NewSessionManager("./sessions", os.Getenv("SESSION_SECRET_KEY"), os.Getenv("SESSION_KEY"))

	router.InitializeRouter()
	template.InitializeRenderer("./templates", isDev)
	ctx := context.Background()

	// > module user
	user.RegisterRoutes(user.NewDependencies(db, ctx), sessionManager)
	// < module user

	// > module daily_activity
	daily_activity.RegisterRoutes(daily_activity.NewDependencies(db, ctx))
	// < module daily_activity

	// Start the HTTP server
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router.GetChiRouter(sessionManager)))
}

func initDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", os.Getenv("DB_CONNECTION"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
