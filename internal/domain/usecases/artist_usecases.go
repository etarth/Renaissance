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
		u.logger.Named("GetAllArtists").Error("artist not found", zap.Error(err))
		return nil, apperror.NotFoundError("artist not found")
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

func (u *artistUsecase) GetArtistById(artistId string) (*dtos.ArtistDTO, *apperror.AppError) {
	artist, err := u.artistRepository.GetArtistById(artistId)
	if err != nil {
		u.logger.Named("GetArtistbyId").Error("artist not found", zap.Error(err))
		return nil, apperror.NotFoundError("artist not found")
	}

	res := dtos.ArtistDTO{
		ArtistId:        artist.ArtistId,
		Bio:             artist.Bio,
		Website:         artist.Website,
		SocialLinks:     artist.SocialLinks,
		ProfileImageURL: artist.ProfileImageURL,
	}

	return &res, nil
}

func (u *artistUsecase) GetArtistByUserId(userId string) (*dtos.ArtistDTO, *apperror.AppError) {
	artist, err := u.artistRepository.GetArtistByUserId(userId)
	if err != nil {
		u.logger.Named("GetArtistbyUserId").Error("user not found", zap.Error(err))
		return nil, apperror.NotFoundError("user not found")
	}

	res := dtos.ArtistDTO{
		ArtistId:        artist.ArtistId,
		UserId:          artist.UserId,
		Bio:             artist.Bio,
		Website:         artist.Website,
		SocialLinks:     artist.SocialLinks,
		ProfileImageURL: artist.ProfileImageURL,
	}

	return &res, nil
}

func (u *artistUsecase) InsertNewArtist(dto *dtos.InsertNewArtistDTO) *apperror.AppError {
	existingArtist, err := u.artistRepository.GetArtistByUserId(dto.UserId)
	if err != nil {
		u.logger.Named("CreateArtist").Error("user not found", zap.String("userID", dto.UserId), zap.Error(err))
		return apperror.NotFoundError("user not found")
	}

	if existingArtist != nil {
		u.logger.Named("CreateArtist").Error("user already exists", zap.String("userID", dto.UserId))
		return apperror.BadRequestError("user already exists")
	}

	newArtist := entities.Artist{
		ArtistId:        dto.ArtistId,
		UserId:          dto.UserId,
		Bio:             dto.Bio,
		Website:         dto.Website,
		SocialLinks:     dto.SocialLinks,
		ProfileImageURL: dto.ProfileImageURL,
	}

	if err := u.artistRepository.InsertNewArtist(newArtist); err != true {
		u.logger.Named("CreateArtist").Error("Failed to insert artist", zap.String("userID", dto.UserId))
		return apperror.InternalServerError("Failed to insert artist")
	}

	u.logger.Named("CreateArtist").Info("Success: ", zap.String("user_id", newArtist.ArtistId))
	return nil
}

func (u *artistUsecase) UpdateArtistById(newData dtos.ArtistDTO, artistId string) *apperror.AppError {
	artist, err := u.artistRepository.GetArtistById(artistId)
	if err != nil {
		u.logger.Named("UpdateArtistbyId").Error("artist not found", zap.Error(err))
		return apperror.NotFoundError("artist not found")
	}

	artist.Bio = newData.Bio
	artist.Website = newData.Website
	artist.SocialLinks = newData.SocialLinks
	artist.ProfileImageURL = newData.ProfileImageURL

	if err := u.artistRepository.UpdateArtistById(*artist, artistId); err != nil {
		u.logger.Named("UpdateArtistbyId").Error("artist cannot be updated", zap.String("artistID", artistId), zap.Error(err))
		return apperror.NotFoundError("artist cannot be updated")
	}

	u.logger.Named("UpdateArtistbyId").Info("Success: ", zap.String("artist_id", artist.ArtistId))
	return nil
}

func (u *artistUsecase) UpdateArtistByUserId(newData dtos.ArtistDTO, userId string) *apperror.AppError {
	artist, err := u.artistRepository.GetArtistByUserId(userId)
	if err != nil {
		u.logger.Named("UpdateArtistByUserId").Error("artist not found", zap.Error(err))
		return apperror.NotFoundError("artist not found")
	}

	artist.Bio = newData.Bio
	artist.Website = newData.Website
	artist.SocialLinks = newData.SocialLinks
	artist.ProfileImageURL = newData.ProfileImageURL

	if err := u.artistRepository.UpdateArtistByUserId(*artist, userId); err != nil {
		return apperror.InternalServerError("cannot update artist")
	}

	u.logger.Named("UpdateArtistByUserId").Info("Success: ", zap.String("user_id", artist.UserId))
	return nil
}

func (u *artistUsecase) DeleteArtistById(artistId string) *apperror.AppError {
	if err := u.artistRepository.DeleteArtistById(artistId); err != nil {
		u.logger.Named("DeleteArtistById").Error("artist cannot be deleted", zap.String("artistID", artistId), zap.Error(err))
		return apperror.NotFoundError("artist cannot be deleted")
	}

	u.logger.Named("DeleteArtistById").Info("Success: ", zap.String("artist_id", artistId))
	return nil
}

func (u *artistUsecase) DeleteArtistByUserId(userId string) *apperror.AppError {
	if err := u.artistRepository.DeleteArtistByUserId(userId); err != nil {
		u.logger.Named("DeleteArtistByUserId").Error("artist cannot be deleted", zap.String("userId", userId), zap.Error(err))
		return apperror.NotFoundError("artist cannot be deleted")
	}

	u.logger.Named("DeleteArtistByUserId").Info("Success: ", zap.String("user_id", userId))
	return nil
}
