package router

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

// Initialize servers
func Serve() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "4599"
	}
	routers := getRouters()
	server01 := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      routers[0],
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	g.Go(func() error {
		return server01.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		log.Fatal(err)
	}
}
