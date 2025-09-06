package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewServer(h *Handler, port string) *http.Server {
	r := chi.NewRouter()

	r.Route("/subscriptions", func(r chi.Router) {
		r.Post("/", h.CreateSubscription)
		r.Get("/", h.ListSubscriptions)
		r.Get("/{id}", h.GetSubscription)
		r.Put("/{id}", h.UpdateSubscription)
		r.Delete("/{id}", h.DeleteSubscription)
		r.Get("/summary", h.SumSubscriptions)
	})

	return &http.Server{
		Addr:    ":" + port,
		Handler: r,
	}
}
