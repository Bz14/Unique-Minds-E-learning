package utils

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// PaginationByPage is a function to return the pagination options for the mongoDB query
func PaginationByPage(pageNo int64, pageSize int64) *options.FindOptions {
	skip := (int64(pageNo) - 1) * pageSize
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
	findOptions.SetSkip(skip)
	findOptions.SetLimit(pageSize)
	return findOptions
}

