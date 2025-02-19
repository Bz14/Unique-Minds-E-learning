package infrastructure

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Db struct{
}

type DB interface{
	Connection(URI string) *mongo.Client 
	ConnectToDatabase()*mongo.Client
	CreateDb(collectionName string)*mongo.Collection
}

func NewDatabase()*Db{
	return &Db{}
}

// Connection function establishes a connection to the database
func (db *Db)Connection(URI string) *mongo.Client {
	// ctx, cancelCtx := context.WithTimeout(context.Background(), time.Duration(dbTimeout)*time.Second)
	// defer cancelCtx()
	connection, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		log.Fatal(err)
	}
	err = connection.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	// defer connection.Disconnect(context.TODO())
	return connection
}

// ConnectToDatabase function connects to the database
func (db *Db)ConnectToDatabase(dbHost string)*mongo.Client{
	connection := db.Connection(dbHost)
	return connection
} 


// CreateDb function creates a database
func (db *Db)CreateDb(dbHost string, dbName string, collectionName string)*mongo.Collection{
	connection := db.ConnectToDatabase(dbHost)
	collection := connection.Database("ELEARNING").Collection(collectionName)
	return collection
} 
