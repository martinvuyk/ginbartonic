package main

import (
	"src/api"
	"src/api/router"

	_ "src/docs"
)

// @title		Example GinBarTonic API
// @version		1.0
func main() {

	api.Init()

	router.Serve()
}
