package main

import (
	"src/api"
	"src/api/router"

	_ "src/docs"
)

// @title		Example GinBarTonic API
// @version		1.0
// @description	This is an example implementation for the ginBarTonic Framework

// @contact.url    http://example.com/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:4599
// @BasePath  /api/v1

func main() {

	api.Init()

	router.Serve()
}
