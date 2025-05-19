package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
	"github.com/tropicaltux/weather-subscription-service/internal/services"
)

// Unsubscribe handles unsubscription requests
func (h *Handler) Unsubscribe(ctx context.Context, request api.UnsubscribeRequestObject) (api.UnsubscribeResponseObject, error) {
	// Validate token
	if request.Token == "" {
		return api.Unsubscribe400JSONResponse{
			Message: "Token is required",
		}, nil
	}

	if request.Token != strings.TrimSpace(request.Token) {
		return api.Unsubscribe400JSONResponse{
			Message: "Token is invalid",
		}, nil
	}

	err := h.subscriptionService.Unsubscribe(ctx, request.Token)
	if err != nil {
		if err == services.ErrInvalidInput {
			return api.Unsubscribe400JSONResponse{
				Message: "Invalid token format",
			}, nil
		}

		return api.Unsubscribe404JSONResponse{
			Message: "Subscription not found",
		}, nil
	}

	return api.Unsubscribe200Response{}, nil
}
