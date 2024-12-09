package domain

// ResponseWrapper struct

type ResponseWrapper struct{
	Message string `bson:"message", json:"message"`
	Data interface{} `bson:"data", json:"data"`
	Status int `bson:status, json:"status"`
}