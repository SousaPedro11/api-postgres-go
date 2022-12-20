package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/sousapedro11/api-postgresql-go/models"
)

// List godoc
// @Summary List all todos
// @Description List all todos
// @Tags todos
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Todo
// @Failure 500 {string} string "Internal Server Error"
// @Router / [get]
func List(writer http.ResponseWriter, request *http.Request) {
	todos, err := models.GetAll()

	if err != nil {
		log.Printf("Error getting all todos: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(todos)
}
