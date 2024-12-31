package repositories

type Repository interface {
	Artist() IArtistRepository
	Artwork() IArtworkRepository
	Wishlist() IWishlistRepository
	OrderItems() IOrderItemsRepository
	// User() UserRepository
	// Attachment() AttachmentRepository
	// Document() DocumentRepository
}
