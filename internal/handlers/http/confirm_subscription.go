package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/internal/services"
)

// ConfirmSubscription handles subscription confirmation
func (h *Handler) ConfirmSubscription(ctx context.Context, request api.ConfirmSubscriptionRequestObject) (api.ConfirmSubscriptionResponseObject, error) {
	// Validate token
	if request.Token == "" {
		return api.ConfirmSubscription400JSONResponse{
			Message: "Token is required",
		}, nil
	}

	if request.Token != strings.TrimSpace(request.Token) {
		return api.ConfirmSubscription400JSONResponse{
			Message: "Token is invalid",
		}, nil
	}

	err := h.subscriptionService.ConfirmSubscription(ctx, request.Token)
	if err != nil {
		if err == services.ErrInvalidInput {
			return api.ConfirmSubscription400JSONResponse{
				Message: "Invalid token format",
			}, nil
		}

		return api.ConfirmSubscription404JSONResponse{
			Message: "Subscription not found or already confirmed",
		}, nil
	}

	return api.ConfirmSubscription200Response{}, nil
}
