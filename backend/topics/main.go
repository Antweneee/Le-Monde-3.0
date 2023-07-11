package main

import (
	r "main/router"
	db "main/database"
)

func main() {
	var database = db.DatabaseInit()

	r := r.Router(database)
	
	r.Run(":8083")
}
