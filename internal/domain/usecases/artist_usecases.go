package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"

	"go.uber.org/zap"
)

type artistUsecase struct {
	cfg              config.Config
	logger           *zap.Logger
	artistRepository repositories.IArtistRepository
}

func NewArtistUsecases(cfg config.Config, logger *zap.Logger, artistRepository repositories.IArtistRepository) IArtistUsecase {
	return &artistUsecase{
		cfg:              cfg,
		logger:           logger,
		artistRepository: artistRepository,
	}
}

func (u *artistUsecase) GetAllArtists() ([]dtos.ArtistDTO, *apperror.AppError) {
	artists, err := u.artistRepository.GetAllArtists()
	if err != nil {
		return nil, apperror.InternalServerError("user not found")
	}

	res := make([]dtos.ArtistDTO, len(artists))
	for i := 0; i < len(artists); i++ {
		res[i] = dtos.ArtistDTO{
			ArtistId:        (artists)[i].ArtistId,
			Bio:             (artists)[i].Bio,
			Website:         (artists)[i].Website,
			SocialLinks:     (artists)[i].SocialLinks,
			ProfileImageURL: (artists)[i].ProfileImageURL,
		}
	}
	return res, nil
}

func (u *artistUsecase) InsertNewArtist(dto *dtos.InsertNewArtistDTO) *apperror.AppError {
	// existingArtist, err := u.artistRepository.GetArtistByUserId(dto.UserId)
	// if err != nil {
	// 	u.logger.Named("CreateArtist").Error("user not found", zap.String("userID", dto.UserId), zap.Error(err))
	// 	return apperror.NotFoundError("user not found")
	// }

	// if existingArtist != nil {
	// 	u.logger.Named("CreateArtist").Error("user already exists", zap.String("userID", dto.UserId))
	// 	return apperror.BadRequestError("user already exists")
	// }

	newArtist := entities.Artist{
		ArtistId:        dto.ArtistId,
		UserId:          dto.UserId,
		Bio:             dto.Bio,
		Website:         dto.Website,
		SocialLinks:     dto.SocialLinks,
		ProfileImageURL: dto.ProfileImageURL,
	}

	if err := u.artistRepository.InsertNewArtist(newArtist); err != true {
		u.logger.Named("CreateArtist").Error("Failed to insert user", zap.String("userID", dto.UserId))
		return apperror.InternalServerError("Failed to insert user")
	}

	u.logger.Named("CreateArtist").Info("Success: ", zap.String("user_id", newArtist.UserId))
	return nil
}
