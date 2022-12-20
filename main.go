package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/sousapedro11/api-postgresql-go/configs"
	"github.com/sousapedro11/api-postgresql-go/handlers"

	_ "github.com/sousapedro11/api-postgresql-go/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title API PostgreSQL Go
// @version 1.0
// @description API PostgreSQL Go
// @termsOfService http://swagger.io/terms/

// @contact.name Pedro Sousa
// @contact.email ppls2106@gmail.com

// license.name MIT
// license.url http://opensource.org/licenses/MIT
func main() {
	err := configs.Load()

	if err != nil {
		panic(err)
	}

	fmt.Printf("Server running on port %s\n", configs.GetServerPort())

	router := chi.NewRouter()
	router.Post("/", handlers.Create)
	router.Get("/{id}", handlers.Retrieve)
	router.Patch("/{id}", handlers.Update)
	router.Put("/{id}", handlers.Update)
	router.Delete("/{id}", handlers.Delete)
	router.Get("/", handlers.List)

	router.Mount("/swagger", httpSwagger.Handler())

	http.ListenAndServe(fmt.Sprintf(":%s", configs.GetServerPort()), router)
}
