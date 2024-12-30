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
	ArtistHandler   *ArtistHandler
	ArtworkHandler  *ArtworkHandler
	WishlistHandler *WishlistHandler
}

func NewHandler(usecases usecases.Usecase, validator validator.DTOValidator) Handler {
	return &handler{
		// MiddlewareHandler: NewMiddlewareHandler(usecases.Middleware()),
		// AuthHandler:       NewAuthHandler(usecases.Auth()),
		// UserHandler:       NewUserHandler(usecases.User(), validator),
		// AttachmentHandler: NewAttachmentHandler(usecases.Attachment()),
		// DocumentHandler:   NewDocumentHandler(usecases.Document(), validator),
		ArtworkHandler:  NewArtworkHandler(usecases.Artwork(), validator),
		ArtistHandler:   NewArtistHandler(usecases.Artist(), validator),
		WishlistHandler: NewWishlistHandler(usecases.Wishlist(), validator),
	}
}

// func (h *handler) Middleware() *MiddlewareHandler {
// 	return h.MiddlewareHandler
// }

func (h *handler) Artist() *ArtistHandler {
	return h.ArtistHandler
}

func (h *handler) Artwork() *ArtworkHandler {
	return h.ArtworkHandler
}

func (h *handler) Wishlist() *WishlistHandler {
	return h.WishlistHandler
}
