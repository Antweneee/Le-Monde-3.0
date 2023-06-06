package main

import (
	r "backend/router"
	db "backend/database"
)

func main() {
	var database = db.DatabaseInteraction()

	r := r.Router(database)
	r.Run()
}
