package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ckalagara/pub-a-player/core"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBConnFormat  = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable"
	APIPathPrefix = "/v1/api/"
)

func Main() {
	ctxMain := context.Background()

	// create db client
	db, err := createDBClient()
	if err != nil {
		log.Fatal("Failed to create db client", err)
	}

	// create handler
	h := core.NewHandler(ctxMain, db)

	// start service
	startServer(ctxMain, h)
}

func createDBClient() (*gorm.DB, error) {

	n := os.Getenv("ENV_DB_NAME")
	u := os.Getenv("ENV_DB_USERNAME")
	p := os.Getenv("ENV_DB_PASSWORD")
	port := os.Getenv("ENV_DB_PORT")
	host := os.Getenv("ENV_DB_HOST")

	if u == "" || p == "" {
		return nil, errors.New("DB connection environment variables not set")
	}

	dns := fmt.Sprintf(DBConnFormat, host, u, p, n, port)

	config := postgres.Config{
		DSN:                  dns,
		PreferSimpleProtocol: true,
	}

	dbClient, err := gorm.Open(postgres.New(config), new(gorm.Config))
	if err != nil {
		return nil, err
	}

	return dbClient, nil
}

func startServer(ctx context.Context, h core.Handler) {
	mux := http.NewServeMux()
	mux.Handle(APIPathPrefix, http.StripPrefix(APIPathPrefix, appMux(ctx, h)))

	err := http.ListenAndServe(":8086", mux)
	if err != nil {
		log.Fatal("Failed to serve", err)
	}

}

func appMux(ctx context.Context, h core.Handler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		if h.Health(ctx) != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
		}
	})

	mux.HandleFunc("GET /player", h.GetPlayer)
	mux.HandleFunc("POST /player", h.UpdatePlayer)
	mux.HandleFunc("PUT /player", h.UpdatePlayer)

	mux.HandleFunc("POST /attachment", h.UploadAttachment)

	return mux
}
