package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {

	err := initMain()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Server Up", "port", port)
	log.Fatal(http.ListenAndServe(port, router))

}

func initMain() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return nil
}
