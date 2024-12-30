package handlers

type Handler interface {
	// Middleware() *MiddlewareHandler
	// Auth() *AuthHandler
	// User() *UserHandler
	// Attachment() *AttachmentHandler
	// Document() *DocumentHandler
	Artist() *ArtistHandler
	Artwork() *ArtworkHandler
	Wishlist() *WishlistHandler
}
