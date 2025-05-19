package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/internal/models"
	"github.com/tropicaltux/weather-subscription-service/internal/services"
)

// Subscribe handles subscription requests
func (h *Handler) Subscribe(ctx context.Context, request api.SubscribeRequestObject) (api.SubscribeResponseObject, error) {
	// Validate request body
	if request.Body == nil {
		return api.Subscribe400JSONResponse{
			Message: "Request body is required",
		}, nil
	}

	request.Body.City = strings.TrimSpace(request.Body.City)
	email := string(request.Body.Email)
	email = strings.TrimSpace(email)

	if request.Body.City == "" {
		return api.Subscribe400JSONResponse{
			Message: "city parameter is required",
		}, nil
	}

	// TODO: Implement city validation to check if city exists and is correctly formatted

	if email == "" {
		return api.Subscribe400JSONResponse{
			Message: "email parameter is required",
		}, nil
	}

	// TODO: Implement more robust email validation beyond the basic OpenAPI format check
	// Consider checking MX records and implementing disposable email detection

	var frequency models.SubscriptionFrequency
	switch request.Body.Frequency {
	case "hourly":
		frequency = models.FrequencyHourly
	case "daily":
		frequency = models.FrequencyDaily
	default:
		return api.Subscribe400JSONResponse{
			Message: "frequency must be either 'hourly' or 'daily'",
		}, nil
	}

	_, err := h.subscriptionService.Subscribe(ctx, email, request.Body.City, frequency)
	if err != nil {
		switch err {
		case services.ErrInvalidInput:
			return api.Subscribe400JSONResponse{
				Message: "Invalid input parameters",
			}, nil
		case services.ErrDuplicateEmail:
			return api.Subscribe409JSONResponse{
				Message: "Email is already subscribed",
			}, nil
		default:
			return api.Subscribe400JSONResponse{
				Message: "Failed to create subscription: " + err.Error(),
			}, nil
		}
	}

	return api.Subscribe200Response{}, nil
}
