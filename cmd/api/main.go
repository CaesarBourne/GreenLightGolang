package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net/http"
)


	// fmt.Println("Hello, World!")
	const version = "1.0.0"
	// fmt.Println("Version:", version)

	type config struct {
		port int
		env  string
	}

	type application struct {
		config config
		logger *log.Logger
	}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar((&cfg.env, "env", "development", "Environment (development|staging|production)"))
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:  fmt.Sprintf(":%d", cfg.port),
	}

}
