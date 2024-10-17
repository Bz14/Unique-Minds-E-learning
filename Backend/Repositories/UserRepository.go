package repositories

import (
	"context"
	"time"
	domain "unique-minds/Domain"
	infrastructures "unique-minds/Infrastructures"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	session mongo.Session
	userCollection *mongo.Collection
	unverifiedCollection *mongo.Collection
	config infrastructures.Config
}

func NewUserRepository(session mongo.Session, collection *mongo.Collection, unverifiedCollection *mongo.Collection, config infrastructures.Config) *UserRepository {
	return &UserRepository{
		session: session,
		userCollection: collection,
		unverifiedCollection: unverifiedCollection,
		config: config,
	}
}


func (ur *UserRepository) FindUserByEmail(email string) error {
	var user *domain.User
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	err := ur.userCollection.FindOne(context, bson.M{"email":email}).Decode(&user)
	return err
}

func (ur *UserRepository) FindUnverifiedUserByEmail(email string) (domain.User, error) {
	var user domain.User
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	err := ur.unverifiedCollection.FindOne(context, bson.M{"email":email}).Decode(&user)
	return user, err
}

func (ur *UserRepository) SaveUnverifiedUser(user *domain.User) error{
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	user.ID = primitive.NewObjectID()
	_, err := ur.unverifiedCollection.InsertOne(context, user)
	return err
}
func (ur *UserRepository) CreateUser(user *domain.User) error{
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()
	user.ID = primitive.NewObjectID()
	_, err := ur.userCollection.InsertOne(context, user)
	return err
}

func (ur *UserRepository) UpdateUnverifiedUser(email string, currentTime time.Time, token string, expires time.Time)error{
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	filter := bson.M{
		"email" : email,
	}

	update := bson.M{
		"$set" : bson.M{
			"updated_at" : currentTime,
			"verification_token" : token,
			"token_expire" : expires,
		},
	}
	_, err := ur.unverifiedCollection.UpdateOne(context, filter, update)
	return err
}

func (ur *UserRepository) FindUserByToken(token string) (domain.User, error){
	var user domain.User
	timeOut := ur.config.TimeOut
	context, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	filter := bson.M{
		"verification_token" : token,
	}
	err := ur.unverifiedCollection.FindOne(context, filter).Decode(&user)
	return user, err
}

func (ur *UserRepository) SignUpUser(user domain.User) error {
	timeOut := ur.config.TimeOut
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeOut)*time.Second)
	defer cancel()

	defer ur.session.EndSession(ctx)

	callback := func(sesCtx mongo.SessionContext) (interface{}, error) {
		_, err := ur.userCollection.InsertOne(sesCtx, user)
		if err != nil {
			return nil, err
		}

		_, err = ur.unverifiedCollection.DeleteOne(sesCtx, bson.M{"email": user.Email})
		if err != nil {
			return nil, err
		}

		return nil, nil
	}
	_, err := ur.session.WithTransaction(ctx, callback)
	return err
}
