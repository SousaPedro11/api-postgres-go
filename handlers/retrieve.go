package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sousapedro11/api-postgresql-go/models"
)

// Retrieve godoc
// @Summary Retrieve a todo
// @Description Retrieve a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Todo
// @Failure 500 {string} string "Internal Server Error"
// @Failure 404 {string} string "Not Found"
// @Router /{id} [get]
func Retrieve(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Error converting id to integer: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	todo, err := models.Get(int64(id))

	if err != nil {
		log.Printf("Error getting todo: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if todo == (models.Todo{}) {
		log.Printf("Todo not found: %v", err)
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	writer.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(writer).Encode(todo); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
