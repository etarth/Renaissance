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

type artistRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

func NewArtistRepository(db *MongoDB) IArtistRepository {
	return &artistRepository{
		Context:    db.Context,
		Collection: db.Database.Collection("Artist"),
	}
}

func (r *artistRepository) GetAllArtists() ([]entities.Artist, error) {
	options := options.Find()
	filter := bson.M{}
	cursor, err := r.Collection.Find(r.Context, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.Context)

	pack := make([]entities.Artist, 0)
	for cursor.Next(r.Context) {
		var item entities.Artist

		err := cursor.Decode(&item)
		if err != nil {
			continue
		}

		pack = append(pack, item)
	}
	return pack, nil
}

func (r *artistRepository) GetArtistById(artistId string) (*entities.Artist, error) {
	var artist entities.Artist
	filter := bson.M{"artist_id": artistId}

	err := r.Collection.FindOne(r.Context, filter).Decode(&artist)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	return &artist, nil
}

func (r *artistRepository) GetArtistByUserId(userId string) (*entities.Artist, error) {
	var artist entities.Artist
	filter := bson.M{"user_id": userId}

	err := r.Collection.FindOne(r.Context, filter).Decode(&artist)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &artist, nil
}

func (r *artistRepository) InsertNewArtist(data entities.Artist) bool {
	if _, err := r.Collection.InsertOne(r.Context, data); err != nil {
		fiberlog.Errorf("Artist -> InsertNewArtist: %s \n", err)
		return false
	}
	return true
}

func (r *artistRepository) UpdateArtistById(newData entities.Artist, artistId string) error {
	artistData := bson.M{"artist_id": artistId}

	_, err := r.Collection.UpdateOne(r.Context, artistData, bson.M{"$set": newData})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}

func (r *artistRepository) UpdateArtistByUserId(newData entities.Artist, userId string) error {
	artistData := bson.M{"user_id": userId}

	_, err := r.Collection.UpdateOne(r.Context, artistData, bson.M{"$set": newData})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}

func (r *artistRepository) DeleteArtistById(artistId string) error {
	artistData := bson.M{"artist_id": artistId}

	_, err := r.Collection.DeleteOne(r.Context, artistData)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}

func (r *artistRepository) DeleteArtistByUserId(userId string) error {
	artistData := bson.M{"user_id": userId}

	_, err := r.Collection.DeleteOne(r.Context, artistData)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}
