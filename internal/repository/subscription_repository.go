package repository

import (
	"context"

	"github.com/tropicaltux/weather-subscription-service/internal/models"
)

// SubscriptionRepository defines the interface for subscription data operations
type SubscriptionRepository interface {
	// Create creates a new subscription in the database
	Create(ctx context.Context, subscription *models.Subscription) error

	// Confirm sets a subscription as confirmed
	Confirm(ctx context.Context, token string) error

	// Delete removes a subscription
	Delete(ctx context.Context, token string) error

	// GetAllActiveSubscriptionsSortedByCity returns all confirmed subscriptions sorted by city
	GetAllActiveSubscriptionsSortedByCity(ctx context.Context) ([]models.Subscription, error)
}
