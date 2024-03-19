package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

// config struct to hold all the configuration settings
type config struct {
	port int
	env string
}

// to hold the dependencies for HTTP handlers, helpers, and middleware
type application struct {
	config config
	logger *log.Logger
}
func main() {
	fmt.Println("Hello world!")

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: mux,
		IdleTimeout: 10 * time.Second,
		ReadTimeout: 30 * time.Second,
	}

	// Start the HTTP server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)



}