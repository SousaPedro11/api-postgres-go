package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/sousapedro11/api-postgresql-go/models"
)

// Delete godoc
// @Summary Delete a todo
// @Description Delete a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} models.Response
// @Failure 500 {string} string "Internal Server Error"
// @Failure 404 {string} string "Not Found"
// @Router /{id} [delete]
func Delete(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(request, "id"))

	if err != nil {
		log.Printf("Error converting id to integer: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rowsAffected, err := models.Delete(int64(id))

	if err != nil {
		log.Printf("Error deleting todo: %v", err)
		http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		log.Printf("Todo not found: %v", err)
		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	} else if rowsAffected > 1 {
		log.Printf("More than registries todo was deleted: %v deleted", rowsAffected)
	}

	resp := map[string]any{
		"Error":   false,
		"Message": "Todo deleted successfully",
	}

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(resp)
}
