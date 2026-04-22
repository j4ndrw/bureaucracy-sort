package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/j4ndrw/bureaucracysort/gateway/internal/reverseproxy"
	"github.com/j4ndrw/bureaucracysort/gateway/pkg/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Could not log application config - %v. Aborting...", err.Error())
	}

	gateway, err := config.NewGateway(&cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	for prefix, cluster := range gateway {
		handler, err := reverseproxy.New(cluster)
		if err != nil {
			log.Fatal(err.Error())
		}
		r.Mount(prefix, handler)
	}

	http.ListenAndServe(
		fmt.Sprintf("%s:%d", cfg.Gateway.Host, cfg.Gateway.Port),
		r,
	)
}
