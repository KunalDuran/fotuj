package data

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"time"

// 	"go.mongodb.org/mongo-driver/bson"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// )

// var (
// 	COLLECTION_BUCKET = "bucket"
// 	COLLECTION_VENDOR = "vendor"
// )

// type MongoDB struct {
// 	Client *mongo.Client
// }

// func InitDB(mongoURI string) *MongoDB {

// 	var err error
// 	if mongoURI == "" {
// 		mongoURI = "mongodb://localhost:27017"
// 		log.Default().Println("MONGO_URI not found, using default value")
// 	}

// 	clientOptions := options.Client().ApplyURI(mongoURI)
// 	c, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return &MongoDB{
// 		Client: c,
// 	}
// }

// func (m MongoDB) UpdateStatus(key, image string, status int) error {
// 	filter := bson.M{
// 		"key":         key,
// 		"images.path": image,
// 	}
// 	update := bson.M{
// 		"$set": bson.M{
// 			"images.$.status": status,
// 			"updated_at":      time.Now().Format(time.RFC3339),
// 		},
// 	}

// 	return UpdateOne(COLLECTION_BUCKET, filter, update)
// }

// func ListProjects(vendor string) error {
// 	filter := bson.M{
// 		"photographerID": vendor,
// 	}

// 	var result []Project
// 	return FindAll(COLLECTION_BUCKET, filter, &result)
// }

// func GetCollection(name string) *mongo.Collection {
// 	return db.Database("fotuj").Collection(name)
// }

// func InsertOne(collection string, document interface{}) error {
// 	c := GetCollection(collection)
// 	_, err := c.InsertOne(context.Background(), document)
// 	return err
// }

// func FindOne(collection string, filter map[string]interface{}, result interface{}) error {
// 	var f = make(bson.M)
// 	for k, v := range filter {
// 		f[k] = v
// 	}

// 	c := GetCollection(collection)
// 	err := c.FindOne(context.Background(), f).Decode(result)
// 	return err
// }

// func FindAll(collection string, filter bson.M, results interface{}, opts ...*options.FindOptions) error {
// 	c := GetCollection(collection)
// 	cursor, err := c.Find(context.Background(), filter)
// 	if err != nil {
// 		return err
// 	}
// 	defer cursor.Close(context.Background())

// 	return cursor.All(context.Background(), results)
// }

// func UpdateOne(collection string, filter bson.M, update bson.M, opts ...*options.UpdateOptions) error {
// 	c := GetCollection(collection)
// 	r, err := c.UpdateOne(context.Background(), filter, update, opts...)
// 	fmt.Println(r)
// 	return err
// }

// func DeleteOne(collection string, filter bson.M) error {
// 	c := GetCollection(collection)
// 	_, err := c.DeleteOne(context.Background(), filter)
// 	return err
// }
