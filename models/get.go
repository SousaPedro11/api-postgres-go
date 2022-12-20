package models

import "github.com/sousapedro11/api-postgresql-go/db"

func Get(id int64) (todo Todo, err error) {

	connection, err := db.OpenConnection()

	if err != nil {
		return
	}

	defer connection.Close()

	sql := "SELECT * FROM todos WHERE id = $1"

	err = connection.QueryRow(sql, id).Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Done, &todo.CreatedAt)

	if err != nil {
		return
	}

	return
}
