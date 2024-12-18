package repositories

type Repository interface {
	Artist() IArtistRepository
	Artwork() IArtworkRepository
	// User() UserRepository
	// Attachment() AttachmentRepository
	// Document() DocumentRepository
}
