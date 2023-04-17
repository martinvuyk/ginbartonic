package router

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func getFromEnv(name string, defaultVal string) string {
	value := os.Getenv(name)
	if value == "" {
		value = defaultVal
	}
	return value
}

func getServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}

func router02() http.Handler {
	e := gin.New()
	e.Use(gin.Recovery())
	e.GET("/", func(c *gin.Context) {
		c.JSON(
			http.StatusOK,
			gin.H{
				"code":    http.StatusOK,
				"message": "Welcome server 02",
			},
		)
	})

	return e
}

// Initialize servers
func Serve() {
	endpoints := getRouters()

	port := getFromEnv("SERVER_PORT", "4599")
	serverV1 := getServer(port, endpoints.api[0])

	monitoring := getFromEnv("MONITORING_PORT", "4788")
	serverMonitoring := getServer(monitoring, endpoints.metrics[0])

	g.Go(func() error {
		return serverV1.ListenAndServe()
	})
	g.Go(func() error {
		return serverMonitoring.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
