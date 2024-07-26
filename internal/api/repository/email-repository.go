package repository

import (
	"FirstAPI/internal/api/customError"
	"FirstAPI/internal/api/model"
	"FirstAPI/internal/infra/db"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmailRepository struct {
	db         *db.MongoDBService
	collection *mongo.Collection
}

func NewEmailRepository(db *db.MongoDBService) *EmailRepository {
	collection := db.Client.Database("emails").Collection("emails")

	return &EmailRepository{
		db:         db,
		collection: collection,
	}
}

func (repo *EmailRepository) Find() ([]model.Email, error) {
	cursor, err := repo.collection.Find(context.TODO(), bson.D{})

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return []model.Email{}, customError.ErrNotFound
		}
		return []model.Email{}, errors.New(err.Error())
	}
	defer func(cursor *mongo.Cursor, ctx context.Context) {
		err := cursor.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cursor, context.TODO())

	var emails []model.Email
	for cursor.Next(context.TODO()) {
		var email model.Email
		err := cursor.Decode(&email)
		if err != nil {
			panic(err)
		}
		emails = append(emails, email)
	}

	return emails, nil
}

func (repo *EmailRepository) FindByEmail(email string) (model.Email, error) {
	var emailDoc model.Email
	err := repo.collection.FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&emailDoc)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Email{}, customError.ErrNotFound
		}
		return model.Email{}, errors.New(err.Error())
	}

	return emailDoc, nil
}

func (repo *EmailRepository) FindByID(id string) (model.Email, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Email{}, err
	}

	var email model.Email
	err = repo.collection.FindOne(context.TODO(), bson.D{{"_id", oid}}).Decode(&email)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Email{}, customError.ErrNotFound
		}
		return model.Email{}, errors.New(err.Error())
	}

	return email, nil
}

func (repo *EmailRepository) Insert(email, status string) (model.Email, error) {
	result, err := repo.collection.InsertOne(context.TODO(), bson.D{
		{"email", email},
		{"status", status},
	})
	if err != nil {
		return model.Email{}, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return model.Email{}, errors.New("invalid Object ID")
	}

	emailDoc := model.Email{
		ID:     oid.Hex(),
		Email:  email,
		Status: status,
	}

	return emailDoc, nil
}

func (repo *EmailRepository) Delete(id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := repo.collection.DeleteOne(context.TODO(), bson.D{{"_id", oid}})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return customError.ErrNotFound
	}

	return nil
}

func (repo *EmailRepository) Update(id string, email string, status string) (model.Email, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return model.Email{}, err
	}

	filter := bson.D{{"_id", oid}}
	update := bson.D{}

	if email != "" {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{"email", email}}})
	}

	if status != "" {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{"status", status}}})
	}

	result, err := repo.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return model.Email{}, err
	}

	if result.MatchedCount == 0 {
		return model.Email{}, customError.ErrNotFound
	}

	emailDoc, err := repo.FindByID(id)
	if err != nil {
		return model.Email{}, err
	}

	return emailDoc, nil
}
