package repo

import (
	"database/sql"

	"github.com/kmwk10/subscription-service/internal/models"
)

type SubscriptionRepo struct {
	DB *sql.DB
}

func NewSubscriptionRepo(db *sql.DB) *SubscriptionRepo {
	return &SubscriptionRepo{DB: db}
}

func (r *SubscriptionRepo) Create(s *models.Subscription) error {
	_, err := r.DB.Exec(
		"INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5)",
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate,
	)
	return err
}

func (r *SubscriptionRepo) GetByID(id int) (*models.Subscription, error) {
	s := &models.Subscription{}
	err := r.DB.QueryRow(
		"SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id=$1", id,
	).Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
	return s, err
}

func (r *SubscriptionRepo) Update(s *models.Subscription) error {
	_, err := r.DB.Exec(
		"UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5 WHERE id=$6",
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, s.ID,
	)
	return err
}

func (r *SubscriptionRepo) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM subscriptions WHERE id=$1", id)
	return err
}

func (r *SubscriptionRepo) List() ([]*models.Subscription, error) {
	rows, err := r.DB.Query("SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	subs := []*models.Subscription{}
	for rows.Next() {
		s := &models.Subscription{}
		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
			return nil, err
		}
		subs = append(subs, s)
	}
	return subs, nil
}
