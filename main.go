package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ddaraujo/neoway_etl_test/db"
	"github.com/ddaraujo/neoway_etl_test/handler"
	"github.com/joho/godotenv"
)

func main() {
	// Read .ENV file
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Some error occured. Err: %s", envErr)
	}

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")

	// Initialize database (5 retries in case of database connection failed)
	count := 0
	for {
		database, err := db.Initialize(dbUser, dbPassword, dbName)
		if err != nil {
			if count < 10 {
				log.Println("Error connectiong to database. Retrying...")
				time.Sleep(5 * time.Second)
				count++
				continue
			} else {
				log.Fatalf("Could not set up database (error): %v", err)
			}
		}

		defer database.Conn.Close()

		// Http handler
		log.Println("Starting HTTP handler")
		httpHandler := handler.NewHandler(database)
		server := &http.Server{
			Handler: httpHandler,
		}
		defer Stop(server)

		ch := make(chan os.Signal, 1)

		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

		log.Println(fmt.Sprint(<-ch))
		log.Println("Stopping API server.")

		break
	}

}

func Stop(server *http.Server) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Could not shut down server correctly: %v\n", err)
		os.Exit(1)
	}
}
