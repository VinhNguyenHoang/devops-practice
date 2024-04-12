package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBImpl struct {
	URI string

	DB *mongo.Client
}

func (c *MongoDBImpl) Init(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.URI))
	if err != nil {
		return err
	}
	c.DB = client

	log.Println("Init database successful...")

	return nil
}

func (c *MongoDBImpl) GetDatabase(dbName string) *mongo.Database {
	if c.DB == nil {
		return nil
	}

	return c.DB.Database(dbName)
}

func (c *MongoDBImpl) Query(ctx context.Context, args ...interface{}) {

}
