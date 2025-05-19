package services

import (
	"context"
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/tropicaltux/weather-subscription-service/internal/models"
	"github.com/tropicaltux/weather-subscription-service/internal/repository"
	"github.com/tropicaltux/weather-subscription-service/internal/repository/db"
)

var (
	ErrInvalidInput         = errors.New("invalid input parameters")
	ErrDuplicateEmail       = errors.New("email already subscribed")
	ErrSubscriptionNotFound = errors.New("subscription not found")
	ErrTokenInvalid         = errors.New("token is invalid")
)

type SubscriptionService struct {
	repo repository.SubscriptionRepository
}

func NewSubscriptionService(repo repository.SubscriptionRepository) *SubscriptionService {
	return &SubscriptionService{
		repo: repo,
	}
}

// generateRandomToken creates a random token of specified length using URL-safe characters
func generateRandomToken(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	charsetLength := big.NewInt(int64(len(charset)))

	result := make([]byte, length)
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, charsetLength)
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}

	return string(result), nil
}

func (s *SubscriptionService) Subscribe(ctx context.Context, email, city string, frequency models.SubscriptionFrequency) (*models.Subscription, error) {
	if email == "" || city == "" {
		return nil, ErrInvalidInput
	}

	// Generate a random token of length 128
	token, err := generateRandomToken(128)
	if err != nil {
		return nil, err
	}

	subscription := &models.Subscription{
		ID:        uuid.New().String(),
		Email:     email,
		City:      city,
		Frequency: frequency,
		Token:     token,
		Confirmed: false,
	}

	err = s.repo.Create(ctx, subscription)
	if err != nil {
		// Check for duplicate email error using proper error type assertion
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // Unique violation code
			return nil, ErrDuplicateEmail
		}
		return nil, err
	}

	return subscription, nil
}

func (s *SubscriptionService) ConfirmSubscription(ctx context.Context, token string) error {
	if token == "" {
		return ErrInvalidInput
	}

	err := s.repo.Confirm(ctx, token)
	if err != nil {
		if errors.Is(err, db.ErrSubscriptionNotFound) {
			return ErrSubscriptionNotFound
		}
		return err
	}

	return nil
}

func (s *SubscriptionService) Unsubscribe(ctx context.Context, token string) error {
	if token == "" {
		return ErrInvalidInput
	}

	err := s.repo.Delete(ctx, token)
	if err != nil {
		if errors.Is(err, db.ErrSubscriptionNotFound) {
			return ErrSubscriptionNotFound
		}
		return err
	}

	return nil
}
