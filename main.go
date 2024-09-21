package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"
	"personal/health-app/daos"
	datebase "personal/health-app/database"
	"personal/health-app/handlers"

	"github.com/pkg/errors"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

//go:embed public
var FS embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := datebase.GetDB()
	if err != nil {
		println(errors.Wrapf(err, "failed to connect database").Error())
		panic(err)
	}
	log.Print("database connected")

	daos := daos.NewDAOs(db)
	countHendler := handlers.NewCounterHandler(*daos)

	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Handle("/", handlers.Make(handlers.HandleHomeIndex))
	router.Handle("/counter", handlers.Make(countHendler.HandleCountersIndex))
	router.Handle("/counter/{id}/increment", handlers.Make(countHendler.HandleCounterIncrementUpdate))
	router.Handle("/counter/{id}/decrement", handlers.Make(countHendler.HandleCounterDecrementUpdate))

	port := os.Getenv("HTTP_LISTEN_PORT")
	slog.Info("app running on", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
