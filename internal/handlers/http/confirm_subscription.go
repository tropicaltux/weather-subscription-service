package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
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

	// TODO: Add business logic
	return api.ConfirmSubscription200Response{}, nil
}
