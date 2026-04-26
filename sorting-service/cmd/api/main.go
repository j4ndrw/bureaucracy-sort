package main

import (
	"j4ndrw/bureaucracysort/config/pkg/config"
	"j4ndrw/bureaucracysort/config/pkg/mq"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not load config. %s - Aborting...", err.Error())
	}
	conn, err := mq.Connect(&cfg)
	mq.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	http.ListenAndServe(":3000", r)
}
