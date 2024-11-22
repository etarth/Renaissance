package repositories

type Repository interface {
	Artist() IArtistRepository
	// User() UserRepository
	// Attachment() AttachmentRepository
	// Document() DocumentRepository
}
