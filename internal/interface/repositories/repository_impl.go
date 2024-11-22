package repositories

import (
	"Backend/pkg/config"
	. "Backend/pkg/database"
	"Backend/pkg/s3client"
)

type repository struct {
	ArtistRepository IArtistRepository
	// UserRepository       UserRepository
	// AttachmentRepository AttachmentRepository
	// DocumentRepository   DocumentRepository
}

func NewRepository(cfg config.Config, db *MongoDB, s3 s3client.S3Client) Repository {
	return &repository{
		ArtistRepository: NewArtistRepository(db),
		// UserRepository:       NewUserRepository(db),
		// AttachmentRepository: NewAttachmentRepository(db, s3),
		// DocumentRepository:   NewDocumentRepository(db),
	}
}

func (r *repository) Artist() IArtistRepository {
	return r.ArtistRepository
}

// func (r *repository) User() UserRepository {
// 	return r.UserRepository
// }

// func (r *repository) Attachment() AttachmentRepository {
// 	return r.AttachmentRepository
// }

// func (r *repository) Document() DocumentRepository {
// 	return r.DocumentRepository
// }
