package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer() *http.Server {
	r := chi.NewRouter()

	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", createSubscription)
		r.Get("/", listSubscriptions)
		r.Get("/{id}", getSubscription)
		r.Put("/{id}", updateSubscription)
		r.Delete("/{id}", deleteSubscription)
	})

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}

func createSubscription(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("create subscription"))
}

func listSubscriptions(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("list subscriptions"))
}

func getSubscription(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get subscription"))
}

func updateSubscription(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update subscription"))
}

func deleteSubscription(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete subscription"))
}
