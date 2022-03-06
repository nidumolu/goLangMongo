package main

import (
	"context"
	"fmt"
	"os"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	fmt.Println("clientOptions type:", reflect.TypeOf(clientOptions), "\n")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo.Connect() ERROR:", err)
		os.Exit(1)
	}
	// Access a MongoDB collection through a database
	usersCollection := client.Database("mymongo").Collection("Icecreams")
	fmt.Println("Collection type:", reflect.TypeOf(usersCollection), "\n")

	// insert a single document into a collection
	// create a bson.D object
	iceCream := bson.D{{"flavour", "Chocolate"}, {"price", 30}}
	// insert the bson object using InsertOne()
	result, err := usersCollection.InsertOne(context.TODO(), iceCream)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result.InsertedID)

	// insert multiple documents into a collection
	// create a slice of bson.D objects
	iceCreams := []interface{}{
		bson.D{{"flavour", "Mint Chocolate chip"}, {"price", 40}},
		bson.D{{"flavour", "Vanilla"}, {"price", 20}},
		bson.D{{"flavour", "Strawberry"}, {"price", 35}},
	}
	// insert the bson object slice using InsertMany()
	results, err := usersCollection.InsertMany(context.TODO(), iceCreams)
	// check for errors in the insertion
	if err != nil {
		panic(err)
	}
	// display the ids of the newly inserted objects
	fmt.Println(results.InsertedIDs)
}
