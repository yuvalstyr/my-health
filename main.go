package main

import (
	"embed"
	"log"
	"log/slog"
	"net/http"
	"os"
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

	_, err = datebase.GetDB()
	if err != nil {
		println(errors.Wrapf(err, "failed to connect database").Error())
		panic(err)
	}

	log.Print("database connected")

	router := chi.NewMux()

	router.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(FS))))
	router.Handle("/", handlers.Make(handlers.HandleHomeIndex))

	port := os.Getenv("HTTP_LISTEN_PORT")
	slog.Info("app running on", "port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
//
// 	dbURL := os.Getenv("DBhandlers.HandleHomeIndex))
// 		return views.Render(ctx, dashboard.Show(dishes, activities))
// }

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
//
// 	dbURL := os.Getenv("DB_URL")
//
// 	dbInstance, err := db.New(dbURL)
// 	if err != nil {
// 		println(errors.Wrapf(err, "failed to connect database").Error())
// 		panic(err)
// 	}
// 	daoFactory := dao.NewDAOs(dbInstance.DB)
//
// 	app := echo.New()
//
// 	handlersFactory := handlers.NewHandlersFactory(*daoFactory)
// 	handlers.InitRoutes(app, &handlersFactory)
//
// 	// TODO: temp route
// 	app.GET("/", func(ctx echo.Context) error {
// 		var dishes []model.MealDish
// 		res := dbInstance.DB.Find(&dishes)
// 		if res.Error != nil {
// 			return res.Error
// 		}
//
// 		activities, err := daoFactory.ActivityDAO.GetActivityDetails("", "2024-02-02")
// 		if err != nil {
// 			return errors.Wrap(err, "could not get activity details")
// 		}
// 		if len(activities) == 0 {
// 			return errors.Wrap(err, "not activities found")
// 		}
// 		var activityTypes []model.ActivityType
// 		res = dbInstance.DB.Find(&activityTypes)
// 		if res.Error != nil {
// 			return res.Error
// 		}
//
// 		return views.Render(ctx, dashboard.Show(dishes, activities))
// 	})
//
// 	app.Start("localhost:4040")
// }
