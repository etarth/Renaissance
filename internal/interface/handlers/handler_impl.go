package handlers

import (
	"Backend/internal/domain/usecases"
	"Backend/pkg/validator"
)

type handler struct {
	// MiddlewareHandler *MiddlewareHandler
	// AuthHandler       *AuthHandler
	// UserHandler       *UserHandler
	// AttachmentHandler *AttachmentHandler
	// DocumentHandler   *DocumentHandler
	ArtistHandler *ArtistHandler
}

func NewHandler(usecases usecases.Usecase, validator validator.DTOValidator) Handler {
	return &handler{
		// MiddlewareHandler: NewMiddlewareHandler(usecases.Middleware()),
		// AuthHandler:       NewAuthHandler(usecases.Auth()),
		// UserHandler:       NewUserHandler(usecases.User(), validator),
		// AttachmentHandler: NewAttachmentHandler(usecases.Attachment()),
		// DocumentHandler:   NewDocumentHandler(usecases.Document(), validator),
		ArtistHandler: NewArtistHandler(usecases.Artist(), validator),
	}
}

// func (h *handler) Middleware() *MiddlewareHandler {
// 	return h.MiddlewareHandler
// }

func (h *handler) Artist() *ArtistHandler {
	return h.ArtistHandler
}
