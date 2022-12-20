package models

import "github.com/sousapedro11/api-postgresql-go/db"

func Delete(id int64) (int64, error) {
	connection, err := db.OpenConnection()

	if err != nil {
		return 0, err
	}

	defer connection.Close()

	sql := "DELETE FROM todos WHERE id = $1"

	res, err := connection.Exec(sql, id)

	if err != nil {
		return 0, err
	}

	return res.RowsAffected()
}
