package main

import (
	r "backend/router"
)

func main() {
	r := r.Router()
	r.Run()
}
