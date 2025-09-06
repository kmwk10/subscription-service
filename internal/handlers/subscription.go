package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/kmwk10/subscription-service/internal/models"
	"github.com/kmwk10/subscription-service/internal/repo"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	Repo *repo.SubscriptionRepo
}

// POST /subscriptions
func (h *Handler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var s models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.Repo.Create(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

// GET /subscriptions/{id}
func (h *Handler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	s, err := h.Repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// PUT /subscriptions/{id}
func (h *Handler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	var s models.Subscription
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.ID = id
	if err := h.Repo.Update(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(s)
}

// DELETE /subscriptions/{id}
func (h *Handler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	if err := h.Repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// GET /subscriptions
func (h *Handler) ListSubscriptions(w http.ResponseWriter, r *http.Request) {
	subs, err := h.Repo.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(subs)
}

// GET /subscriptions/summary
func (h *Handler) SumSubscriptions(w http.ResponseWriter, r *http.Request) {
	userID := r.URL.Query().Get("user_id")
	serviceName := r.URL.Query().Get("service_name")
	startStr := r.URL.Query().Get("start")
	endStr := r.URL.Query().Get("end")

	start, err := time.Parse("2006-01", startStr)
	if err != nil {
		http.Error(w, "invalid start date", http.StatusBadRequest)
		return
	}
	end, err := time.Parse("2006-01", endStr)
	if err != nil {
		http.Error(w, "invalid end date", http.StatusBadRequest)
		return
	}

	sum, err := h.Repo.SumPrice(r.Context(), userID, serviceName, start, end)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"total": sum})
}
