package data

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db                *mongo.Client
	COLLECTION_BUCKET = "bucket"
	COLLECTION_VENDOR = "vendor"
)

func InitDB(mongoURI string) {

	var err error
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
		log.Default().Println("MONGO_URI not found, using default value")
	}

	clientOptions := options.Client().ApplyURI(mongoURI)
	db, err = mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
}

func GetCollection(name string) *mongo.Collection {
	return db.Database("fotuj").Collection(name)
}

func GetCollections() []string {
	collections, err := db.Database("fotuj").ListCollectionNames(context.Background(), bson.M{})
	if err != nil {
		panic(err)
	}
	return collections
}

func InsertOne(collection string, document interface{}) error {
	c := GetCollection(collection)
	_, err := c.InsertOne(context.Background(), document)
	return err
}

func FindOne(collection string, filter map[string]interface{}, result interface{}) error {
	var f = make(bson.M)
	for k, v := range filter {
		f[k] = v
	}

	c := GetCollection(collection)
	err := c.FindOne(context.Background(), f).Decode(result)
	return err
}

func FindAll(collection string, filter bson.M, results interface{}, opts ...*options.FindOptions) error {
	c := GetCollection(collection)
	cursor, err := c.Find(context.Background(), filter)
	if err != nil {
		return err
	}
	defer cursor.Close(context.Background())

	return cursor.All(context.Background(), results)
}

func UpdateOne(collection string, filter bson.M, update bson.M, opts ...*options.UpdateOptions) error {
	c := GetCollection(collection)

	_, err := c.UpdateOne(context.Background(), filter, update, opts...)
	return err
}

func DeleteOne(collection string, filter bson.M) error {
	c := GetCollection(collection)
	_, err := c.DeleteOne(context.Background(), filter)
	return err
}
