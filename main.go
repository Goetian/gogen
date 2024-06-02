package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/goetian/gogen/handler"
	supabase "github.com/goetian/gogen/pkg/sb"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	err := initMain()
	if err != nil {
		log.Fatal(err)
	}
	router := chi.NewMux()
	router.Use(handler.WithAuthentication)

	staticFileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Handle("/static/*", staticFileServer)

	router.Get("/", handler.MakeHandler(handler.HandleHomeIndex))
	router.Get("/login", handler.MakeHandler(handler.HandleLogInUser))
	router.Post("/handleLogin", handler.MakeHandler(handler.HandleLogin))
	port := os.Getenv("HTTP_LISTEN_ADDR")
	slog.Info("Server Up", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

func initMain() error {
	if err := godotenv.Load(); err != nil {
		return err
	}
	return supabase.Init()
}
