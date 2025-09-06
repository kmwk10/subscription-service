package repo

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/kmwk10/subscription-service/internal/models"
)

type SubscriptionRepo struct {
	DB *sql.DB
}

func NewSubscriptionRepo(db *sql.DB) *SubscriptionRepo {
	return &SubscriptionRepo{DB: db}
}

func (r *SubscriptionRepo) Create(ctx context.Context, s *models.Subscription) error {
	_, err := r.DB.ExecContext(ctx,
		"INSERT INTO subscriptions (service_name, price, user_id, start_date, end_date) VALUES ($1, $2, $3, $4, $5)",
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate,
	)
	if err != nil {
		log.Printf("Error creating subscription: %v", err)
		return err
	}
	log.Printf("Created subscription for user %s, service %s", s.UserID, s.ServiceName)
	return nil
}

func (r *SubscriptionRepo) GetByID(ctx context.Context, id int) (*models.Subscription, error) {
	s := &models.Subscription{}
	err := r.DB.QueryRowContext(ctx,
		"SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions WHERE id=$1", id,
	).Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("Subscription with id %d not found", id)
		} else {
			log.Printf("Error fetching subscription by id %d: %v", id, err)
		}
		return nil, err
	}
	return s, nil
}

func (r *SubscriptionRepo) Update(ctx context.Context, s *models.Subscription) error {
	_, err := r.DB.ExecContext(ctx,
		"UPDATE subscriptions SET service_name=$1, price=$2, user_id=$3, start_date=$4, end_date=$5 WHERE id=$6",
		s.ServiceName, s.Price, s.UserID, s.StartDate, s.EndDate, s.ID,
	)
	if err != nil {
		log.Printf("Error updating subscription id %d: %v", s.ID, err)
		return err
	}
	log.Printf("Updated subscription id %d", s.ID)
	return nil
}

func (r *SubscriptionRepo) Delete(ctx context.Context, id int) error {
	_, err := r.DB.ExecContext(ctx, "DELETE FROM subscriptions WHERE id=$1", id)
	if err != nil {
		log.Printf("Error deleting subscription id %d: %v", id, err)
		return err
	}
	log.Printf("Deleted subscription id %d", id)
	return nil
}

func (r *SubscriptionRepo) List(ctx context.Context) ([]*models.Subscription, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, service_name, price, user_id, start_date, end_date FROM subscriptions")
	if err != nil {
		log.Printf("Error listing subscriptions: %v", err)
		return nil, err
	}
	defer rows.Close()

	subs := []*models.Subscription{}
	for rows.Next() {
		s := &models.Subscription{}
		if err := rows.Scan(&s.ID, &s.ServiceName, &s.Price, &s.UserID, &s.StartDate, &s.EndDate); err != nil {
			log.Printf("Error scanning subscription: %v", err)
			return nil, err
		}
		subs = append(subs, s)
	}
	log.Printf("Listed %d subscriptions", len(subs))
	return subs, nil
}

func (r *SubscriptionRepo) SumPrice(ctx context.Context, userID string, serviceName string, start, end time.Time) (int, error) {
	query := `SELECT COALESCE(SUM(price), 0) FROM subscriptions WHERE start_date >= $1 AND start_date <= $2`
	args := []interface{}{start, end}
	i := 3
	if userID != "" {
		query += fmt.Sprintf(" AND user_id = $%d", i)
		args = append(args, userID)
		i++
	}
	if serviceName != "" {
		query += fmt.Sprintf(" AND service_name = $%d", i)
		args = append(args, serviceName)
		i++
	}

	var sum int
	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&sum)
	if err != nil {
		log.Printf("Error calculating sum price: %v", err)
		return 0, err
	}
	log.Printf("Calculated sum price: %d for user %s service %s", sum, userID, serviceName)
	return sum, nil
}
