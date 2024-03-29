package networking

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

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

// Initialize servers
func Serve() {
	endpoints := getRoutersV1()

	port := getFromEnv("SERVER_PORT", "4599")
	serverV1 := getServer(port, endpoints.api)

	monitoringV1 := getFromEnv("MONITORING_PORT", "4788")
	serverMonitoring := getServer(monitoringV1, endpoints.metrics)

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
