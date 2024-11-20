package repositories

import (
	"Backend/pkg/config"
	ds "Backend/pkg/database"
	"Backend/pkg/s3client"
)

type repository struct {
	// UserRepository       UserRepository
	// AttachmentRepository AttachmentRepository
	// DocumentRepository   DocumentRepository
}

func NewRepository(cfg config.Config, db *ds.MongoDB, s3 s3client.S3Client) Repository {
	return &repository{
		// UserRepository:       NewUserRepository(db),
		// AttachmentRepository: NewAttachmentRepository(db, s3),
		// DocumentRepository:   NewDocumentRepository(db),
	}
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
