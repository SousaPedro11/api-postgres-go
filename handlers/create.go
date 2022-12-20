package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sousapedro11/api-postgresql-go/models"
)

var Validator = validator.New()

// Create godoc
// @Summary Create a todo
// @Description Create a todo
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body models.Todo true "Todo"
// @Success 200 {object} models.Response
// @Failure 500 {string} string "Internal Server Error"
// @Router / [post]
func Create(writer http.ResponseWriter, request *http.Request) {
	var todo models.Todo
	var errors []*models.EntityValidatorError
	var resp models.Response

	err := json.NewDecoder(request.Body).Decode(&todo)

	if err != nil {
		log.Printf("Error decoding JSON: %v", err)
		http.Error(writer, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = Validator.Struct(todo)

	if err != nil {

		for _, err := range err.(validator.ValidationErrors) {
			var elementError models.EntityValidatorError
			elementError.Field = err.Field()
			elementError.Tag = err.Tag()
			elementError.Value = err.Param()

			errors = append(errors, &elementError)
		}

		log.Printf("Error validating todo: %v", err)
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		responseBody := map[string]any{
			"Error":   "true",
			"Message": &errors,
		}

		if err := json.NewEncoder(writer).Encode(responseBody); err != nil {
			log.Printf("Error decoding JSON: %v", err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}

		return
	}

	id, err := models.Insert(todo)

	if err != nil {
		resp.Error = true
		resp.Message = fmt.Sprintf("Error inserting todo in database: %v", err)

	} else {
		resp.Error = false
		resp.Message = fmt.Sprintf("Todo inserted successfully! Id: %v", id)
	}

	writer.Header().Set("Content-Type", "application/json")

	json.NewEncoder(writer).Encode(resp)
}
