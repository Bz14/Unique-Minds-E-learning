package repositories

import "go.mongodb.org/mongo-driver/mongo"

type UserRepository struct {
	database *mongo.Collection
}

type UserRepositoryInterface interface{}

func NewUserRepository(db *mongo.Collection) *UserRepository {
	return &UserRepository{
		database: db,
	}
}