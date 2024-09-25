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

func (d *Database) CreateDB(url string, databaseName string) (*mongo.Database, error){
	connection, err := d.Connect(url)
	if err != nil{
		return nil, err
	}
	database := connection.Database(databaseName)
	return database, nil
}

func (d *Database) CreateCollection (database *mongo.Database, collectionName string) (*mongo.Collection, error){
	collection := database.Collection(collectionName)
	return collection, nil
}