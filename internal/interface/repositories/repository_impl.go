package repositories

import (
	"Backend/pkg/config"
	. "Backend/pkg/database"
	"Backend/pkg/s3client"
)

type repository struct {
	ArtistRepository     IArtistRepository
	ArtworkRepository    IArtworkRepository
	WishlistRepository   IWishlistRepository
	OrderItemsRepository IOrderItemsRepository
	// UserRepository       UserRepository
	// AttachmentRepository AttachmentRepository
	// DocumentRepository   DocumentRepository
}

func NewRepository(cfg config.Config, db *MongoDB, s3 s3client.S3Client) Repository {
	return &repository{
		ArtistRepository:     NewArtistRepository(db),
		ArtworkRepository:    NewArtworkRepository(db),
		WishlistRepository:   NewWishlistRepository(db),
		OrderItemsRepository: NewOrderItemsRepository(db),
		// UserRepository:       NewUserRepository(db),
		// AttachmentRepository: NewAttachmentRepository(db, s3),
		// DocumentRepository:   NewDocumentRepository(db),
	}
}

func (r *repository) Artist() IArtistRepository {
	return r.ArtistRepository
}

func (r *repository) Artwork() IArtworkRepository {
	return r.ArtworkRepository
}

func (r *repository) Wishlist() IWishlistRepository {
	return r.WishlistRepository
}

func (r *repository) OrderItems() IOrderItemsRepository {
	return r.OrderItemsRepository
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
