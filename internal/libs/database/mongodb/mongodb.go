package mongodb

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDBImpl struct {
	URI    string
	DBName string

	client *mongo.Client

	dbs map[string]*mongo.Database

	cols map[string]*mongo.Collection
}

func (c *MongoDBImpl) Init(ctx context.Context, colNames ...string) error {
	url := c.URI + "/" + c.DBName
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		return err
	}
	c.client = client

	log.Println("Init database successful...")

	db := client.Database(c.DBName)

	if db == nil {
		return fmt.Errorf("%s not exist", c.DBName)
	}

	c.dbs = make(map[string]*mongo.Database)
	c.dbs[c.DBName] = db

	c.cols = make(map[string]*mongo.Collection)
	for _, colname := range colNames {
		col := db.Collection(colname)
		c.cols[colname] = col
	}

	return nil
}

func (c *MongoDBImpl) Insert(ctx context.Context, colName string, doc interface{}) error {
	res, err := c.cols[colName].InsertOne(ctx, doc)
	if err != nil {
		log.Printf("Insert: %+v", err)
		return err
	}

	id := res.InsertedID.(primitive.ObjectID)

	if id.IsZero() {
		log.Printf("Unable to insert document")
		return fmt.Errorf("unable to insert document")
	}

	return nil
}

func (c *MongoDBImpl) Query(ctx context.Context, args ...interface{}) {

}
