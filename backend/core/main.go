package main

import (
	r "main/router"
)

func main() {

	r := r.Router()
	
	r.Run()
}
