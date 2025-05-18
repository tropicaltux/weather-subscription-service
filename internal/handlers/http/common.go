package http

// Handler implements api.StrictServerInterface
type Handler struct{}

// NewHandler creates a new HTTP handler
func NewHandler() *Handler {
	return &Handler{}
}
