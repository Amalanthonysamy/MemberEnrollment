package store

import (
	"context"
	"fmt"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//create mongodb connection establish function and return collection object which will be used by other functions
func MongoDBConnectionHelper() (*mongo.Collection, error) {
	//timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB connection URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}
	defer client.Disconnect(ctx)
	collection := client.Database("MemberEnrollment").Collection("Members")
	return collection, nil
}

//create mongodb connection helper function and return collection object
func InsertMemberToMongoDB(collection, member *Member) (bool, error) {
	_, err := collection.InsertOne(context.Background(), member)
	if err != nil {
		return false, err
	}
	return true, nil
}
