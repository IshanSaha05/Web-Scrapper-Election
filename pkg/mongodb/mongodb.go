package mongodb

import (
	"context"
	"fmt"

	"github.com/IshanSaha05/IndiaVotes/pkg/config"
	"github.com/IshanSaha05/IndiaVotes/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	context    context.Context
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func (object *MongoDB) GetMongoClient() error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoSite))
	if err != nil {
		return err
	}

	object.context = context.Background()
	object.client = client
	object.database = nil
	object.collection = nil

	return nil
}

func (object *MongoDB) SetMongoDatabase(databaseName string) error {
	// Getting all the database names present in the client.
	allDBNames, err := object.client.ListDatabaseNames(object.context, bson.D{})

	if err != nil {
		fmt.Println("Error while fetching all database names present, to compare if the passed name already exists or not.")
		return err
	}

	// Parsing through the list of all the database names.
	for _, name := range allDBNames {

		// If the database with the passed name is already present, error is thrown.
		if name == databaseName {
			return fmt.Errorf(fmt.Sprintf("Database named \"%s\" already exists.", databaseName))
		}
	}

	// Database with passed name is not present, thus new database is created.
	object.database = object.client.Database(databaseName)

	return nil
}

func (object *MongoDB) SetMongoCollection(collectionName string) error {
	// Getting all the collection names present in the database.
	allCollectionNames, err := object.database.ListCollectionNames(object.context, bson.D{})

	if err != nil {
		fmt.Println("Error while fetching all collection names present, to compare if the passed name already exists or not.")
		return err
	}

	// Parsing through the list of all the collection names.
	for _, name := range allCollectionNames {

		// If the collection with the passed name is already present, just the collection is assigned.
		if name == collectionName {
			object.collection = object.database.Collection(collectionName)
			return nil
		}
	}

	// Collection with passed name is not present, thus new collection is created.
	err = object.database.CreateCollection(object.context, collectionName)

	if err != nil {
		return err
	}

	object.collection = object.database.Collection(collectionName)

	return nil
}

func (object *MongoDB) InsertIntoDB(datas []models.ACData) error {
	for _, data := range datas {
		_, err := object.collection.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}
