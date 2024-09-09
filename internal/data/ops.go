package data

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func Buckets(vendor string) error {
	filter := bson.M{
		"photographerID": vendor,
	}

	var result []Bucket
	return FindAll(COLLECTION_BUCKET, filter, &result)
}

func UpdateStatus(key, image string, status int) error {
	filter := bson.M{
		"key":         key,
		"images.path": image,
	}
	update := bson.M{
		"$set": bson.M{
			"images.$.status": status,
			"updated_at":      time.Now().Format(time.RFC3339),
		},
	}

	return UpdateOne(COLLECTION_BUCKET, filter, update)
}
