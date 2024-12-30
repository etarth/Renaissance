package usecases

import (
	"Backend/internal/domain/entities"
	"Backend/internal/interface/dtos"
	"Backend/internal/interface/repositories"
	"Backend/pkg/apperror"
	"Backend/pkg/config"
	"time"

	"go.uber.org/zap"
)

type wishlistUsecase struct {
	cfg                config.Config
	logger             *zap.Logger
	wishlistRepository repositories.IWishlistRepository
}

func NewWishlistUsecases(cfg config.Config, logger *zap.Logger, wishlistRepository repositories.IWishlistRepository) IWishlistUsecase {
	return &wishlistUsecase{
		cfg:                cfg,
		logger:             logger,
		wishlistRepository: wishlistRepository,
	}
}

func (u *wishlistUsecase) GetAllWishlistsByUserId(req *dtos.WishlistDTO, userId string) ([]dtos.WishlistDTO, *apperror.AppError) {
	wishlists, err := u.wishlistRepository.GetAllWishlistsByUserId(userId)
	if err != nil {
		return nil, apperror.InternalServerError("could not retrieve wishlists for the user")
	}

	res := make([]dtos.WishlistDTO, len(wishlists))
	for i, wishlist := range wishlists {
		res[i] = dtos.WishlistDTO{
			FavoriteId: wishlist.FavoriteId,
			UserId:     wishlist.UserId,
			ArtworkId:  wishlist.ArtworkId,
			CreatedAt:  wishlist.CreatedAt,
		}
	}

	return res, nil
}

func (u *wishlistUsecase) InsertNewWishlist(dto *dtos.InsertNewWishlistDTO) *apperror.AppError {
	newWishlist := entities.Wishlist{
		FavoriteId: dto.FavoriteId,
		ArtworkId:  dto.ArtworkId,
		UserId:     dto.UserId,
		CreatedAt:  time.Now(),
	}

	if err := u.wishlistRepository.InsertNewWishlist(newWishlist); err != true {
		u.logger.Named("CreateWishlist").Error("Failed to insert wishlist", zap.String("wishlistID", dto.FavoriteId))
		return apperror.InternalServerError("Failed to insert wishlist")
	}

	u.logger.Named("CreateWishlist").Info("Success: ", zap.String("artist_id", newWishlist.FavoriteId))
	return nil
}

// func (u *artworkUsecase) GetArtworkById(req *dtos.ArtworkDTO, artworkId string) (*dtos.ArtworkDTO, *apperror.AppError) {
// 	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
// 	if err != nil {
// 		return nil, apperror.InternalServerError("failed to fetch artwork")
// 	}

// 	res := &dtos.ArtworkDTO{
// 		ArtworkId:   artwork.ArtworkId,
// 		ArtistId:    artwork.ArtistId,
// 		Title:       artwork.Title,
// 		Description: artwork.Description,
// 		Category:    artwork.Category,
// 		Style:       artwork.Style,
// 		Width:       artwork.Width,
// 		Height:      artwork.Height,
// 		Price:       artwork.Price,
// 		ImageURL:    artwork.ImageURL,
// 		Stock:       artwork.Stock,
// 		CreatedAt:   artwork.CreatedAt,
// 		UpdatedAt:   artwork.UpdatedAt,
// 	}
// 	return res, nil
// }

// func (u *artworkUsecase) UpdateArtworkById(newData dtos.UpdateArtworkByIdDTO, artworkId string) *apperror.AppError {
// 	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
// 	if err != nil {
// 		return apperror.InternalServerError("failed to fetch artwork")
// 	}
// 	if artwork == nil {
// 		return apperror.NotFoundError("artwork not found")
// 	}

// 	// Reflect over newData and artwork to apply updates
// 	newDataValue := reflect.ValueOf(newData)
// 	artworkValue := reflect.ValueOf(artwork).Elem() // Dereference pointer

// 	for i := 0; i < newDataValue.NumField(); i++ {
// 		field := newDataValue.Type().Field(i) // Get the field definition
// 		newValue := newDataValue.Field(i)     // Get the value of the field in newData

// 		// Only update if the field is not the zero value (empty or nil)
// 		if !newValue.IsZero() {
// 			artworkField := artworkValue.FieldByName(field.Name)
// 			if artworkField.IsValid() && artworkField.CanSet() {
// 				// Handle both pointer and non-pointer types
// 				if artworkField.Kind() == reflect.Ptr && newValue.Kind() != reflect.Ptr {
// 					// If artwork field is a pointer but new value is not a pointer, we create a new pointer for the field
// 					artworkField.Set(reflect.New(artworkField.Type().Elem()).Elem())
// 				}
// 				// Update the field with the new value
// 				artworkField.Set(newValue)
// 			}
// 		}
// 	}
// 	artwork.UpdatedAt = time.Now() // Set the updatedAt timestamp

// 	if err := u.artworkRepository.UpdateArtworkById(*artwork, artworkId); err != nil {
// 		u.logger.Named("UpdateArtworkById").Error("Failed to update artwork", zap.String("artwork_id", artworkId))
// 		return apperror.InternalServerError("failed to update artwork")
// 	}

// 	u.logger.Named("UpdateArtworkById").Info("Success", zap.String("artwork_id", artworkId))
// 	return nil
// }

// func (u *artworkUsecase) DeleteArtworkById(artworkId string) *apperror.AppError {
// 	artwork, err := u.artworkRepository.GetArtworkById(artworkId)
// 	if err != nil {
// 		return apperror.InternalServerError("failed to fetch artwork")
// 	}
// 	if artwork == nil {
// 		return apperror.NotFoundError("artwork not found")
// 	}

// 	err = u.artworkRepository.DeleteArtworkById(artworkId)
// 	if err != nil {
// 		u.logger.Named("DeleteArtworkById").Error("Failed to delete artwork", zap.String("artwork_id", artworkId))
// 		return apperror.InternalServerError("failed to delete artwork")
// 	}

// 	u.logger.Named("DeleteArtworkById").Info("Artwork deleted successfully", zap.String("artwork_id", artworkId))
// 	return nil
// }
