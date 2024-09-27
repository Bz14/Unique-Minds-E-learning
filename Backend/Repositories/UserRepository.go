package repositories

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) *UserRepository {
	return &UserRepository{
		userCollection: collection,
	}
}