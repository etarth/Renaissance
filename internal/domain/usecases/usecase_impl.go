package usecases

import (
	"Backend/internal/interface/repositories"
	"Backend/pkg/config"

	"go.uber.org/zap"
)

type usecase struct {
	// MiddlewareUsecase MiddlewareUsecase
	// AuthUsecase       AuthUsecase
	// UserUsecase       UserUsecase
	// AttachmentUsecase AttachmentUsecase
	// DocumentUsecase   DocumentUsecase
	ArtistUsecase   IArtistUsecase
	ArtworkUsecase  IArtworkUsecase
	WishlistUsecase IWishlistUsecase
}

func NewUsecase(repo repositories.Repository, cfg config.Config, logger *zap.Logger) Usecase {
	return &usecase{
		// MiddlewareUsecase: NewMiddlewareUsecase(cfg, logger.Named("MiddlewareSvc"), repo.User()),
		// AuthUsecase:       NewAuthUsecase(cfg, logger.Named("AuthSvc"), repo.User()),
		// UserUsecase:       NewUserUsecase(cfg, logger.Named("UserSvc"), repo.User()),
		// AttachmentUsecase: NewAttachmentUsecase(cfg, logger.Named("AttachmentSvc"), repo.Attachment()),
		// DocumentUsecase:   NewDocumentUsecase(cfg, logger.Named("DocumentSvc"), repo.Document(), repo.User()),
		ArtistUsecase:   NewArtistUsecases(cfg, logger.Named("ArtistSvc"), repo.Artist()),
		ArtworkUsecase:  NewArtworkUsecases(cfg, logger.Named("ArtworkSvc"), repo.Artwork()),
		WishlistUsecase: NewWishlistUsecases(cfg, logger.Named("WishlistSvc"), repo.Wishlist()),
	}
}

// func (u *usecase) Middleware() MiddlewareUsecase {
// 	return u.MiddlewareUsecase
// }

func (u *usecase) Artist() IArtistUsecase {
	return u.ArtistUsecase
}

func (u *usecase) Artwork() IArtworkUsecase {
	return u.ArtworkUsecase
}

func (u *usecase) Wishlist() IWishlistUsecase {
	return u.WishlistUsecase
}
