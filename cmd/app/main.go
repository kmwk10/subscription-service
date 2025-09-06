package main

import (
	"log"
	"net/http"

	"github.com/kmwk10/subscription-service/internal/config"
	"github.com/kmwk10/subscription-service/internal/db"
	"github.com/kmwk10/subscription-service/internal/handlers"
	"github.com/kmwk10/subscription-service/internal/repo"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.Load()
	database := db.Connect(cfg)
	defer database.Close()

	subRepo := repo.NewSubscriptionRepo(database)
	handler := &handlers.Handler{Repo: subRepo}

	r := chi.NewRouter()
	r.Post("/subscriptions", handler.CreateSubscription)
	r.Get("/subscriptions", handler.ListSubscriptions)
	r.Get("/subscriptions/{id}", handler.GetSubscription)
	r.Put("/subscriptions/{id}", handler.UpdateSubscription)
	r.Delete("/subscriptions/{id}", handler.DeleteSubscription)
	r.Get("/subscriptions/summary", handler.SumSubscriptions)

	log.Println("Server started on :" + cfg.AppPort)
	log.Fatal(http.ListenAndServe(":"+cfg.AppPort, r))
}
