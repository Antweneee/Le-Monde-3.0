package main

import (
	db "main/database"
	r "main/router"
)

func main() {
	var database = db.DatabaseInit()

	r := r.Router(database)

	r.Run(":8084")
}
