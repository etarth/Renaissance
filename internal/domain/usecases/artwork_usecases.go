package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"

	"go.uber.org/zap"
)

type artworkUsecase struct {
	cfg               config.Config
	logger            *zap.Logger
	artworkRepository repositories.IArtworkRepository
}

func NewArtworkUsecases(cfg config.Config, logger *zap.Logger, artworkRepository repositories.IArtworkRepository) IArtworkUsecase {
	return &artworkUsecase{
		cfg:               cfg,
		logger:            logger,
		artworkRepository: artworkRepository,
	}
}

func (u *artworkUsecase) GetAllArtworks() ([]dtos.ArtworkDTO, *apperror.AppError) {
	artworks, err := u.artworkRepository.GetAllArtworks()
	if err != nil {
		return nil, apperror.InternalServerError("user not found")
	}

	res := make([]dtos.ArtworkDTO, len(artworks))
	for i := 0; i < len(artworks); i++ {
		res[i] = dtos.ArtworkDTO{
			ArtworkId:   (artworks)[i].ArtworkId,
			ArtistId:    (artworks)[i].ArtistId,
			Title:       (artworks)[i].Title,
			Description: (artworks)[i].Description,
			Category:    (artworks)[i].Category,
			Style:       (artworks)[i].Style,
			Width:       (artworks)[i].Width,
			Height:      (artworks)[i].Height,
			Price:       (artworks)[i].Price,
			ImageURL:    (artworks)[i].ImageURL,
			Stock:       (artworks)[i].Stock,
			CreatedAt:   (artworks)[i].CreatedAt,
			UpdatedAt:   (artworks)[i].UpdatedAt,
		}
	}
	return res, nil
}

func (u *artworkUsecase) GetArtworkById(req *dtos.ArtworkDTO, artworkId string) (*dtos.ArtworkDTO, *apperror.AppError) {
	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
	if err != nil {
		return nil, apperror.InternalServerError("failed to fetch artwork")
	}

	res := &dtos.ArtworkDTO{
		ArtworkId:   artwork.ArtworkId,
		ArtistId:    artwork.ArtistId,
		Title:       artwork.Title,
		Description: artwork.Description,
		Category:    artwork.Category,
		Style:       artwork.Style,
		Width:       artwork.Width,
		Height:      artwork.Height,
		Price:       artwork.Price,
		ImageURL:    artwork.ImageURL,
		Stock:       artwork.Stock,
		CreatedAt:   artwork.CreatedAt,
		UpdatedAt:   artwork.UpdatedAt,
	}
	return res, nil
}

func (u *artworkUsecase) InsertNewArtwork(dto *dtos.InsertNewArtworkDTO) *apperror.AppError {
	newArtwork := entities.Artwork{
		ArtworkId:   dto.ArtworkId,
		ArtistId:    dto.ArtistId,
		Title:       dto.Title,
		Description: dto.Description,
		Category:    dto.Category,
		Style:       dto.Style,
		Width:       dto.Width,
		Height:      dto.Height,
		Price:       dto.Price,
		ImageURL:    dto.ImageURL,
		Stock:       dto.Stock,
		CreatedAt:   dto.CreatedAt,
		UpdatedAt:   dto.UpdatedAt,
	}

	if err := u.artworkRepository.InsertNewArtwork(newArtwork); err != true {
		u.logger.Named("CreateArtwork").Error("Failed to insert user", zap.String("artistID", dto.ArtistId))
		return apperror.InternalServerError("Failed to insert user")
	}

	u.logger.Named("CreateArtwork").Info("Success: ", zap.String("artist_id", newArtwork.ArtistId))
	return nil
}
