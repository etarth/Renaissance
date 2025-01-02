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

type reviewRepository struct {
	Context    context.Context
	Collection *mongo.Collection
}

func NewReviewRepository(db *MongoDB) IReviewRepository {
	return &reviewRepository{
		Context:    db.Context,
		Collection: db.Database.Collection("Review"),
	}
}

func (r *reviewRepository) GetAllReviews() ([]entities.Review, error) {
	options := options.Find()
	filter := bson.M{}
	cursor, err := r.Collection.Find(r.Context, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(r.Context)

	pack := make([]entities.Review, 0)
	for cursor.Next(r.Context) {
		var item entities.Review

		err := cursor.Decode(&item)
		if err != nil {
			continue
		}

		pack = append(pack, item)
	}
	return pack, nil
}

func (r *reviewRepository) GetReviewById(reviewId string) (*entities.Review, error) {
	var review entities.Review
	filter := bson.M{"review_id": reviewId}
	err := r.Collection.FindOne(r.Context, filter).Decode(&review)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &review, nil
}

func (r *reviewRepository) GetReviewsByUserId(userId string) ([]entities.Review, error) {
	var reviews []entities.Review
	filter := bson.M{"user_id": userId}

	cursor, err := r.Collection.Find(r.Context, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(r.Context)

	for cursor.Next(r.Context) {
		var review entities.Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *reviewRepository) GetReviewsByArtistId(artistId string) ([]entities.Review, error) {
	var reviews []entities.Review
	filter := bson.M{"artist_id": artistId}

	cursor, err := r.Collection.Find(r.Context, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	defer cursor.Close(r.Context)

	for cursor.Next(r.Context) {
		var review entities.Review
		if err := cursor.Decode(&review); err != nil {
			return nil, err
		}
		reviews = append(reviews, review)
	}

	return reviews, nil
}

func (r *reviewRepository) InsertNewReview(data entities.Review) bool {
	_, err := r.Collection.InsertOne(r.Context, data)
	if err != nil {
		fiberlog.Errorf("Review -> InsertNewReview: %s \n", err)
		// You can log more details if needed:
		fiberlog.Errorf("Failed to insert review: %v", err)
		return false
	}
	return true
}

func (r *reviewRepository) UpdateReviewById(newData entities.Review, reviewId string) error {
	reviewData := bson.M{"review_id": reviewId}

	_, err := r.Collection.UpdateOne(r.Context, reviewData, bson.M{"$set": newData})

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}

func (r *reviewRepository) DeleteReviewById(reviewId string) error {
	reviewData := bson.M{"review_id": reviewId}

	_, err := r.Collection.DeleteOne(r.Context, reviewData)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil
		}
		return err
	}

	return nil
}

func (r *reviewRepository) GetAverageRatingByArtistId(artistId string) (float64, error) {
	// Match reviews with the provided artistId
	matchStage := bson.M{"$match": bson.M{"artist_id": artistId}}

	// Group to calculate the average rating
	groupStage := bson.M{
		"$group": bson.M{
			"_id":           "$artist_id",
			"averageRating": bson.M{"$avg": bson.M{"$toDouble": "$rating"}},
		},
	}

	// Define the aggregation pipeline
	pipeline := []bson.M{matchStage, groupStage}

	// Perform the aggregation
	cursor, err := r.Collection.Aggregate(r.Context, pipeline)
	if err != nil {
		return 0, err
	}
	defer cursor.Close(r.Context)

	// Parse the result
	var result []struct {
		AverageRating float64 `bson:"averageRating"`
	}
	if err := cursor.All(r.Context, &result); err != nil {
		return 0, err
	}

	// Check if there are no matching reviews
	if len(result) == 0 {
		return 0, nil
	}

	return result[0].AverageRating, nil
}
