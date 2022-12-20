package models

import "github.com/sousapedro11/api-postgresql-go/db"

func GetAll() (todos []Todo, err error) {

	connection, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	sql := "SELECT * FROM todos"

	rows, err := connection.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo

		err = rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}

	return
}
