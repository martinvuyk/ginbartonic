package main

import (
	"src/api"
	"src/api/router"
)

func main() {

	api.Init()

	router.Serve()
}
