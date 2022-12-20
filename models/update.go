package models

import (
	"reflect"
	"strconv"
	"strings"

	"github.com/sousapedro11/api-postgresql-go/db"
)

func Update(id int64, todo Todo) (int64, error) {

	connection, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	sql := "UPDATE todos SET title = $1, description = $2, done = $3 WHERE id = $4"

	res, err := connection.Exec(sql, todo.Title, todo.Description, todo.Done, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}

func Patch(id int64, todo Todo) (int64, error) {

	connection, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	var parameters []string
	var values []interface{}

	reflectType := reflect.TypeOf(todo)
	reflectValue := reflect.ValueOf(todo)

	for i := 0; i < reflectType.NumField(); i++ {
		if reflectValue.Field(i).Interface() != nil && reflectType.Field(i).Name != "ID" && reflectValue.Field(i).Interface() != "" {
			parameters = append(parameters, reflectType.Field(i).Name+" = $"+strconv.Itoa(len(parameters)+1))
			values = append(values, reflectValue.Field(i).Interface())
		}
	}

	if len(parameters) == 0 {
		return 0, nil
	}
	sql := "UPDATE todos SET " + strings.Join(parameters, ", ") + " WHERE id = $" + strconv.Itoa(len(parameters)+1)

	values = append(values, id)

	res, err := connection.Exec(sql, values...)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
