package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
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

	// TODO: Add business logic
	return api.Unsubscribe200Response{}, nil
}
