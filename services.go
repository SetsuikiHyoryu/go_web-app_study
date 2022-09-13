package main

import (
	"log"
)

func getOne(id int) (user customUser, _error error) {
	user = customUser{}

	_error = db.QueryRow("SELECT id, name FROM custom_user WHERE id = ?;", id).Scan(&user.id, &user.name)

	return
}

func getMany(id int) ([]customUser, error) {
	users := []customUser{}

	rows, _error := db.Query("SELECT id, name FROM custom_user WHERE id >= ?;", id)

	for rows.Next() {
		user := customUser{}
		_error = rows.Scan(&user.id, &user.name)

		if _error != nil {
			log.Fatalln(_error.Error())
		}

		users = append(users, user)
	}

	return users, _error
}

func (user *customUser) Insert() (_error error) {
	sql := "INSERT INTO custom_user(id, name) VALUES(?, ?);"
	statement, _error := db.Prepare(sql)

	if _error != nil {
		log.Fatalln(_error.Error())
	}

	defer statement.Close()

	_, _error = statement.Exec(user.id, user.name)

	return _error
}

func (user *customUser) Update() (_error error) {
	_, _error = db.Exec(
		"UPDATE custom_user SET name = ? where id = ?",
		user.name,
		user.id,
	)

	return
}
