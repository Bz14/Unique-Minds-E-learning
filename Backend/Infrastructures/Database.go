package infrastructures

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct{}

type DatabaseInterface interface {
}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) Connect(URL string) (*mongo.Client, error) {
	connection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))
	if err != nil{
		return nil, err
	}
	err = connection.Ping(context.Background(), nil)
	if err != nil{
		return nil, err
	}
	return connection, nil
}

func (d *Database) Disconnect(connection *mongo.Client) error {
	err := connection.Disconnect(context.Background())
	if err != nil{
		return err
	}
	return nil
}

func (d *Database) CreateCollection (url string, databaseName string, collectionName string) (*mongo.Collection, error){
	connection, err := d.Connect(url)
	if err != nil{
		return nil, err
	}
	collection := connection.Database(databaseName).Collection(collectionName)
	return collection, nil
}