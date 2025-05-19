package db

import (
	"context"

	"github.com/tropicaltux/weather-subscription-service/internal/models"
	"github.com/tropicaltux/weather-subscription-service/internal/repository"
	"gorm.io/gorm"
)

type PostgresSubscriptionRepository struct {
	db *gorm.DB
}

func NewPostgresSubscriptionRepository(db *gorm.DB) repository.SubscriptionRepository {
	return &PostgresSubscriptionRepository{db: db}
}

func (r *PostgresSubscriptionRepository) Create(ctx context.Context, subscription *models.Subscription) error {
	return r.db.WithContext(ctx).Create(subscription).Error
}

func (r *PostgresSubscriptionRepository) Confirm(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Model(&models.Subscription{}).
		Where("token = ?", token).
		Updates(map[string]interface{}{
			"confirmed": true,
		}).Error
}

func (r *PostgresSubscriptionRepository) Delete(ctx context.Context, token string) error {
	return r.db.WithContext(ctx).Where("token = ?", token).Delete(&models.Subscription{}).Error
}

func (r *PostgresSubscriptionRepository) GetAllActiveSubscriptionsSortedByCity(ctx context.Context) ([]models.Subscription, error) {
	var subscriptions []models.Subscription

	result := r.db.WithContext(ctx).
		Where("confirmed = ?", true).
		Order("city ASC").
		Find(&subscriptions)

	if result.Error != nil {
		return nil, result.Error
	}

	return subscriptions, nil
}
