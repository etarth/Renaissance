package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"

	"go.uber.org/zap"

	"reflect"
	"time"
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

func (u *artworkUsecase) UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError {
	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
	if err != nil {
		return apperror.InternalServerError("failed to fetch artwork")
	}
	if artwork == nil {
		return apperror.NotFoundError("artwork not found")
	}

	// Reflect over newData and artwork to apply updates
	newDataValue := reflect.ValueOf(newData)
	artworkValue := reflect.ValueOf(artwork).Elem() // Dereference pointer

	for i := 0; i < newDataValue.NumField(); i++ {
		field := newDataValue.Type().Field(i) // Get the field definition
		newValue := newDataValue.Field(i)     // Get the value of the field in newData

		// Only update if the field is not the zero value (empty or nil)
		if !newValue.IsZero() {
			artworkField := artworkValue.FieldByName(field.Name)
			if artworkField.IsValid() && artworkField.CanSet() {
				// Handle both pointer and non-pointer types
				if artworkField.Kind() == reflect.Ptr && newValue.Kind() != reflect.Ptr {
					// If artwork field is a pointer but new value is not a pointer, we create a new pointer for the field
					artworkField.Set(reflect.New(artworkField.Type().Elem()).Elem())
				}
				// Update the field with the new value
				artworkField.Set(newValue)
			}
		}
	}
	artwork.UpdatedAt = time.Now() // Set the updatedAt timestamp

	if err := u.artworkRepository.UpdateArtworkById(*artwork, artworkId); err != nil {
		u.logger.Named("UpdateArtworkById").Error("Failed to update artwork", zap.String("artwork_id", artworkId))
		return apperror.InternalServerError("failed to update artwork")
	}

	u.logger.Named("UpdateArtworkById").Info("Success", zap.String("artwork_id", artworkId))
	return nil
}

func (u *artworkUsecase) DeleteArtworkById(artworkId string) *apperror.AppError {
	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
	if err != nil {
		return apperror.InternalServerError("failed to fetch artwork")
	}
	if artwork == nil {
		return apperror.NotFoundError("artwork not found")
	}

	err = u.artworkRepository.DeleteArtworkById(artworkId)
	if err != nil {
		u.logger.Named("DeleteArtworkById").Error("Failed to delete artwork", zap.String("artwork_id", artworkId))
		return apperror.InternalServerError("failed to delete artwork")
	}

	u.logger.Named("DeleteArtworkById").Info("Artwork deleted successfully", zap.String("artwork_id", artworkId))
	return nil
}
