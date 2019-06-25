package expenseDB

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func ExpCollFunc() *mongo.Collection{
	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//defaultdata := &types.Expense{2, "hitapi", "one", 150, "2019-06-24T17:39:53.51804651+05:30", "2019-06-24T17:39:53.518046732+05:30"}
	collection := client.Database("ExpDB").Collection("ExpColl")
	//_, _ = collection.InsertOne(ctx, defaultdata)
	return collection
}



