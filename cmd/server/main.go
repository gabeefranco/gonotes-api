package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gabeefranco/gonotes-api/internal/config"
	"github.com/gabeefranco/gonotes-api/internal/db"
	"github.com/gabeefranco/gonotes-api/internal/http/handlers"
	"github.com/gabeefranco/gonotes-api/internal/http/routes"
	"github.com/gabeefranco/gonotes-api/internal/repository"
	"github.com/gabeefranco/gonotes-api/internal/service"
	"github.com/go-chi/chi/v5"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := chi.NewRouter()

	config, err := config.NewConfig()
	if err != nil {
		log.Fatalln("error loading environment variables")
	}

	dbConn, err := db.NewDB(config)
	if err != nil {
		log.Fatalln(err)
	}

	usersRepository := repository.NewSqlUsersRepository(dbConn)

	usersService := service.NewUsersService(usersRepository)

	usersHandler := handlers.NewUsersHandler(*usersService)

	usersRoutes := routes.NewUsersRoutes(*usersHandler)
	usersRoutes.Setup(r)

	http.ListenAndServe(fmt.Sprintf(":%s", config.Port), r)
}
