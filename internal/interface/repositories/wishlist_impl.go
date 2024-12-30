package repositories

import (
	"Backend/internal/domain/entities"
	. "Backend/pkg/database"
	"context"

	fiberlog "github.com/gofiber/fiber/v2/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type wishlistRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

func NewWishlistRepository(db *MongoDB) IWishlistRepository {
	return &wishlistRepository{
		Context:    db.Context,
		Collection: db.Database.Collection("Wishlist"),
	}
}

func (r *wishlistRepository) GetAllWishlistsByUserId(userId string) ([]entities.Wishlist, error) {
	filter := bson.M{"user_id": userId}
	options := options.Find()

	cursor, err := r.Collection.Find(r.Context, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.Context)

	wishlists := make([]entities.Wishlist, 0)
	for cursor.Next(r.Context) {
		var item entities.Wishlist
		if err := cursor.Decode(&item); err != nil {
			continue
		}
		wishlists = append(wishlists, item)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	print("kuy")
	return wishlists, nil
}

func (r *wishlistRepository) InsertNewWishlist(data entities.Wishlist) bool {
	if _, err := r.Collection.InsertOne(r.Context, data); err != nil {
		fiberlog.Errorf("Artist -> InsertNewArtist: %s \n", err)
		return false
	}
	return true
}

// func (r *artistRepository) GetArtistById(artistId string) (*entities.Artist, error) {
// 	var artist entities.Artist
// 	filter := bson.M{"artist_id": artistId}

// 	err := r.Collection.FindOne(r.Context, filter).Decode(&artist)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}
// 	return &artist, nil
// }

// func (r *artistRepository) GetArtistByUserId(userId string) (*entities.Artist, error) {
// 	var artist entities.Artist
// 	filter := bson.M{"user_id": userId}

// 	err := r.Collection.FindOne(r.Context, filter).Decode(&artist)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil, nil
// 		}
// 		return nil, err
// 	}

// 	return &artist, nil
// }

// func (r *artistRepository) UpdateArtistById(newData entities.Artist, artistId string) error {
// 	artistData := bson.M{"artist_id": artistId}

// 	_, err := r.Collection.UpdateOne(r.Context, artistData, bson.M{"$set": newData})

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }

// func (r *artistRepository) UpdateArtistByUserId(newData entities.Artist, userId string) error {
// 	artistData := bson.M{"user_id": userId}

// 	_, err := r.Collection.UpdateOne(r.Context, artistData, bson.M{"$set": newData})

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }

// func (r *artistRepository) DeleteArtistById(artistId string) error {
// 	artistData := bson.M{"artist_id": artistId}

// 	_, err := r.Collection.DeleteOne(r.Context, artistData)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }

// func (r *artistRepository) DeleteArtistByUserId(userId string) error {
// 	artistData := bson.M{"user_id": userId}

// 	_, err := r.Collection.DeleteOne(r.Context, artistData)

// 	if err != nil {
// 		if err == mongo.ErrNoDocuments {
// 			return nil
// 		}
// 		return err
// 	}

// 	return nil
// }
