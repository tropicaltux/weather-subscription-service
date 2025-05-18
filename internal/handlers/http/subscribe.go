package http

import (
	"context"
	"strings"

	api "github.com/tropicaltux/weather-subscription-service/internal/api/http"
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

	if request.Body.City == "" {
		return api.Subscribe400JSONResponse{
			Message: "city parameter is required",
		}, nil
	}

	// TODO: Add business logic
	return api.Subscribe200Response{}, nil
}
