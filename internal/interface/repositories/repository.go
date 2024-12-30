package repositories

type Repository interface {
	Artist() IArtistRepository
	Artwork() IArtworkRepository
	Wishlist() IWishlistRepository
	// User() UserRepository
	// Attachment() AttachmentRepository
	// Document() DocumentRepository
}
