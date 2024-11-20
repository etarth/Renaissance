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
}

func NewUsecase(repo repositories.Repository, cfg config.Config, logger *zap.Logger) Usecase {
	return &usecase{
		// MiddlewareUsecase: NewMiddlewareUsecase(cfg, logger.Named("MiddlewareSvc"), repo.User()),
		// AuthUsecase:       NewAuthUsecase(cfg, logger.Named("AuthSvc"), repo.User()),
		// UserUsecase:       NewUserUsecase(cfg, logger.Named("UserSvc"), repo.User()),
		// AttachmentUsecase: NewAttachmentUsecase(cfg, logger.Named("AttachmentSvc"), repo.Attachment()),
		// DocumentUsecase:   NewDocumentUsecase(cfg, logger.Named("DocumentSvc"), repo.Document(), repo.User()),
	}
}

// func (u *usecase) Middleware() MiddlewareUsecase {
// 	return u.MiddlewareUsecase
// }
