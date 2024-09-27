package repositories

import (
	"context"
	"time"
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	userCollection *mongo.Collection
	config infrastructures.Config
}

func NewUserRepository(collection *mongo.Collection, config infrastructures.Config) *UserRepository {
	return &UserRepository{
		userCollection: collection,
		config: config,
	}
}


func (ur *UserRepository) FindUserByEmail(email string) error {
	var user *domain.User
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*20*time.Second)
	defer cancel()
	err := ur.userCollection.FindOne(context, bson.M{"email":email}).Decode(&user)
	return err
}