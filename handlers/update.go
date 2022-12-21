package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sousapedro11/api-postgresql-go/models"
)

// Update godoc
// @Summary Update a todo
// @Description Update a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body models.Todo true "Todo"
// @Success 200 {object} models.Response
// @Failure 500 {string} string "Internal Server Error"
// @Failure 400 {string} string "Bad Request"
// @Router /{id} [put]
func Update(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Error converting id to integer: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var todo models.Todo

	err = json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var rowsAffected int64

	rowsAffected, err = models.Patch(int64(id), todo)

	if err != nil {
		log.Printf("Error updating todo: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		log.Printf("Todo not found: %v", err)
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if rowsAffected > 1 {
		log.Printf("More than registries todo was updated: %v updated", rowsAffected)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Todo updated successfully",
	}

	writer.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(writer).Encode(resp); err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
