package repositories

type Repository interface {
	Artist() IArtistRepository
	Artwork() IArtworkRepository
	Review() IReviewRepository
	// User() UserRepository
	// Attachment() AttachmentRepository
	// Document() DocumentRepository
}
