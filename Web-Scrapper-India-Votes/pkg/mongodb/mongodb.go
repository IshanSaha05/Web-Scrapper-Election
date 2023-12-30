package mongodb

import (
	"context"
	"time"

	"github.com/IshanSaha05/IndiaVotes/pkg/config"
	"github.com/IshanSaha05/IndiaVotes/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	context   context.Context
	client    *mongo.Client
	database  *mongo.Database
	collecton *mongo.Collection
}

func (object *MongoDB) GetMongoClient(ctx context.Context, timeout int) error {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.MongoSite))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(timeout))
	defer cancel()

	object.context = ctx
	object.client = client
	object.database = nil
	object.collecton = nil

	return nil
}

func (object *MongoDB) InsertIntoDB(datas []models.Data) error {
	for _, data := range datas {
		_, err := object.collecton.InsertOne(object.context, data)

		if err != nil {
			return err
		}
	}

	return nil
}
