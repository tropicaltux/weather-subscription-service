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
			Message: "Invalid token format",
		}, nil
	}

	err := h.subscriptionService.ConfirmSubscription(ctx, request.Token)
	if err != nil {
		if err == services.ErrTokenInvalid {
			return api.ConfirmSubscription400JSONResponse{
				Message: "Invalid token",
			}, nil
		}

		if err == services.ErrSubscriptionNotFound {
			return api.ConfirmSubscription404JSONResponse{
				Message: "Subscription not found",
			}, nil
		}

		if err == services.ErrInvalidInput {
			return api.ConfirmSubscription400JSONResponse{
				Message: "Invalid input parameters",
			}, nil
		}

		return api.ConfirmSubscription400JSONResponse{
			Message: "Failed to confirm subscription: " + err.Error(),
		}, nil
	}

	return api.ConfirmSubscription200Response{}, nil
}
